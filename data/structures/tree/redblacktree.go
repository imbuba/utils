package tree

import (
	"fmt"
	"sync"

	"github.com/imbuba/utils/data"
)

const (
	RED   = true
	BLACK = false
)

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{
		pool: sync.Pool{
			New: func() interface{} {
				return &node{}
			},
		},
	}
}

type node struct {
	key   data.Comparable
	value interface{}
	left  *node
	right *node
	color bool
}

func (n *node) Reuse() {
	n.key = n.key.Default()
	n.value = nil
	n.left = nil
	n.right = nil
	n.color = BLACK
}

type RedBlackTree struct {
	root *node
	pool sync.Pool
}

func (t *RedBlackTree) isEmpty() bool {
	return t.root == nil
}

func (t *RedBlackTree) Put(key data.Comparable, value interface{}) {
	t.root = t.put(t.root, key, value)
	t.root.color = BLACK
}

func (t *RedBlackTree) put(x *node, key data.Comparable, value interface{}) *node {
	if x == nil {
		h := t.pool.Get().(*node)
		h.key = key
		h.value = value
		h.color = RED
		return h
	}
	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		x.left = t.put(x.left, key, value)
	} else if cmp > 0 {
		x.right = t.put(x.right, key, value)
	} else {
		x.value = value
	}
	if t.isRed(x.right) && !t.isRed(x.left) {
		x = t.rotateLeft(x)
	}
	if t.isRed(x.left) && t.isRed(x.left.left) {
		x = t.rotateRight(x)
	}
	if t.isRed(x.left) && t.isRed(x.right) {
		t.flipColors(x)
	}
	return x
}

func (t *RedBlackTree) Get(key data.Comparable) interface{} {
	x := t.root
	for x != nil {
		cmp := key.CompareTo(x.key)
		if cmp < 0 {
			x = x.left
		} else if cmp > 0 {
			x = x.right
		} else {
			return x.value
		}
	}
	return nil
}

func (t *RedBlackTree) Contains(key data.Comparable) bool {
	return t.Get(key) != nil
}

func (t *RedBlackTree) Delete(key data.Comparable) {
	if t.isEmpty() || !t.Contains(key) {
		return
	}
	if !t.isRed(t.root.left) && !t.isRed(t.root.right) {
		t.root.color = RED
	}
	t.root = t.deleteKey(t.root, key)
	if !t.isEmpty() {
		t.root.color = BLACK
	}
}

func (t *RedBlackTree) deleteKey(x *node, key data.Comparable) *node {
	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		if !t.isRed(x.left) && !t.isRed(x.left.left) {
			x = t.moveRedLeft(x)
		}
		x.left = t.deleteKey(x.left, key)
	} else {
		if t.isRed(x.left) {
			x = t.rotateRight(x)
		}
		if key.Equals(x.key) && x.right == nil {
			x.Reuse()
			t.pool.Put(x)
			return nil
		}
		if !t.isRed(x.right) && !t.isRed(x.right.left) {
			x = t.moveRedRight(x)
		}
		if key.Equals(x.key) {
			h := t.min(x.right)
			x.key = h.key
			x.value = h.value
			x.right = t.deleteMin(x.right)
		} else {
			x.right = t.deleteKey(x.right, key)
		}
	}
	return t.balance(x)
}

func (t *RedBlackTree) DeleteMin() {
	if t.isEmpty() {
		return
	}
	if !t.isRed(t.root.left) && !t.isRed(t.root.right) {
		t.root.color = RED
	}
	t.root = t.deleteMin(t.root)
	if !t.isEmpty() {
		t.root.color = BLACK
	}
}

func (t *RedBlackTree) deleteMin(x *node) *node {
	if x.left == nil {
		x.Reuse()
		t.pool.Put(x)
		return nil
	}
	if !t.isRed(x.left) && !t.isRed(x.left.left) {
		x = t.moveRedLeft(x)
	}
	x.left = t.deleteMin(x.left)
	return t.balance(x)
}

func (t *RedBlackTree) DeleteMax() {
	if t.isEmpty() {
		return
	}
	if !t.isRed(t.root.left) && !t.isRed(t.root.right) {
		t.root.color = RED
	}
	t.root = t.deleteMax(t.root)
	if !t.isEmpty() {
		t.root.color = BLACK
	}
}

func (t *RedBlackTree) deleteMax(x *node) *node {
	if t.isRed(x.left) {
		x = t.rotateRight(x)
	}
	if x.right == nil {
		x.Reuse()
		t.pool.Put(x)
		return nil
	}
	if !t.isRed(x.right) && !t.isRed(x.right.left) {
		x = t.moveRedRight(x)
	}
	x.right = t.deleteMax(x.right)
	return t.balance(x)
}

func (t *RedBlackTree) Floor(key data.Comparable) interface{} {
	node := t.floor(t.root, key)
	if node != nil {
		return node.value
	}
	return nil
}

func (t *RedBlackTree) floor(x *node, key data.Comparable) *node {
	if x == nil {
		return nil
	}
	cmp := key.CompareTo(x.key)
	if cmp == 0 {
		return x
	}
	if cmp < 0 {
		return t.floor(x.left, key)
	}
	temp := t.floor(x.right, key)
	if temp == nil {
		return x
	}
	return temp
}

func (t *RedBlackTree) Ceil(key data.Comparable) interface{} {
	node := t.ceil(t.root, key)
	if node != nil {
		return node.value
	}
	return nil
}

func (t *RedBlackTree) ceil(x *node, key data.Comparable) *node {
	if x == nil {
		return nil
	}
	cmp := key.CompareTo(x.key)
	if cmp == 0 {
		return x
	}
	if cmp > 0 {
		return t.ceil(x.right, key)
	}
	temp := t.ceil(x.left, key)
	if temp == nil {
		return x
	}
	return temp
}

func (t *RedBlackTree) isRed(x *node) bool {
	if x == nil {
		return false
	}
	return x.color == RED
}

func (t *RedBlackTree) rotateLeft(x *node) *node {
	h := x.right
	x.right = h.left
	h.left = x
	h.color = x.color
	x.color = RED
	return h
}

func (t *RedBlackTree) rotateRight(x *node) *node {
	h := x.left
	x.left = h.right
	h.right = x
	h.color = x.color
	x.color = RED
	return h
}

func (t *RedBlackTree) flipColors(x *node) {
	x.color = !x.color
	x.left.color = !x.left.color
	x.right.color = !x.right.color
}

func (t *RedBlackTree) moveRedLeft(x *node) *node {
	t.flipColors(x)
	if t.isRed(x.right.left) {
		x.right = t.rotateRight(x.right)
		x = t.rotateLeft(x)
		t.flipColors(x)
	}
	return x
}

func (t *RedBlackTree) moveRedRight(x *node) *node {
	t.flipColors(x)
	if t.isRed(x.left.left) {
		x = t.rotateRight(x)
		t.flipColors(x)
	}
	return x
}

func (t *RedBlackTree) balance(x *node) *node {
	if t.isRed(x.right) {
		x = t.rotateLeft(x)
	}
	if t.isRed(x.left) && t.isRed(x.left.left) {
		x = t.rotateRight(x)
	}
	if t.isRed(x.left) && t.isRed(x.right) {
		t.flipColors(x)
	}
	return x
}

func (t *RedBlackTree) min(x *node) *node {
	if x == nil {
		return nil
	}
	for {
		if x.left == nil {
			return x
		}
		x = x.left
	}
}

func (t *RedBlackTree) String() string {
	str := "RedBlackTree\n"
	if !t.isEmpty() {
		output(t.root, "", true, &str)
	}
	return str
}

func (node *node) String() string {
	return fmt.Sprintf("%v", node.key)
}

func output(node *node, prefix string, isTail bool, str *string) {
	if node.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.left, newPrefix, true, str)
	}
}
