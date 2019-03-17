package tree

// NewIntRedBlackTree creates new RedBlackTree
func NewIntRedBlackTree() *IntRedBlackTree {
	return &IntRedBlackTree{}
}

type intNode struct {
	key   int
	value int
	left  *intNode
	right *intNode
	color bool
}

// IntRedBlackTree struct
type IntRedBlackTree struct {
	root *intNode
}

func (t *IntRedBlackTree) isEmpty() bool {
	return t.root == nil
}

// Put add BattleData at a given key
func (t *IntRedBlackTree) Put(key int, value int) {
	t.root = t.put(t.root, key, value)
	t.root.color = BLACK
}

func (t *IntRedBlackTree) put(x *intNode, key int, value int) *intNode {
	if x == nil {
		node := &intNode{}
		node.key = key
		node.value = value
		node.color = RED
		return node
	}
	if key < x.key {
		x.left = t.put(x.left, key, value)
	} else if key > x.key {
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

// Get returns BattleData at a given key
func (t *IntRedBlackTree) Get(key int) int {
	x := t.root
	for x != nil {
		if key < x.key {
			x = x.left
		} else if key > x.key {
			x = x.right
		} else {
			return x.value
		}
	}
	return 0
}

// Contains check that given key in RBT
func (t *IntRedBlackTree) Contains(key int) bool {
	return t.Get(key) != 0
}

// Floor returns element that nearest below or equal for the key
func (t *IntRedBlackTree) Floor(key int) int {
	node := t.floor(t.root, key)
	if node != nil {
		return node.value
	}
	return 0
}

func (t *IntRedBlackTree) floor(x *intNode, key int) *intNode {
	if x == nil {
		return nil
	}
	if x.key == key {
		return x
	}
	if key < x.key {
		return t.floor(x.left, key)
	}
	temp := t.floor(x.right, key)
	if temp == nil {
		return x
	}
	return temp
}

// Ceil returns element that nearest above or equal for the key
func (t *IntRedBlackTree) Ceil(key int) int {
	node := t.ceil(t.root, key)
	if node != nil {
		return node.value
	}
	return 0
}

func (t *IntRedBlackTree) ceil(x *intNode, key int) *intNode {
	if x == nil {
		return nil
	}
	if x.key == key {
		return x
	}
	if key > x.key {
		return t.ceil(x.right, key)
	}
	temp := t.ceil(x.left, key)
	if temp == nil {
		return x
	}
	return temp
}

func (t *IntRedBlackTree) isRed(x *intNode) bool {
	if x == nil {
		return false
	}
	return x.color == RED
}

func (t *IntRedBlackTree) rotateLeft(x *intNode) *intNode {
	h := x.right
	x.right = h.left
	h.left = x
	h.color = x.color
	x.color = RED
	return h
}

func (t *IntRedBlackTree) rotateRight(x *intNode) *intNode {
	h := x.left
	x.left = h.right
	h.right = x
	h.color = x.color
	x.color = RED
	return h
}

func (t *IntRedBlackTree) flipColors(x *intNode) {
	x.color = !x.color
	x.left.color = !x.left.color
	x.right.color = !x.right.color
}
