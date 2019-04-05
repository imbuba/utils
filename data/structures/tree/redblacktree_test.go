package tree

import (
	"testing"

	"github.com/finnan444/utils/data"
)

func fillTree(tree *RedBlackTree) {
	tree.Put(data.Int(5), "e")
	tree.Put(data.Int(6), "f")
	tree.Put(data.Int(7), "g")
	tree.Put(data.Int(3), "c")
	tree.Put(data.Int(4), "d")
	tree.Put(data.Int(1), "a")
	tree.Put(data.Int(2), "b")
	tree.Put(data.Int(1), "x")
}

func TestRedBlackTreePut(t *testing.T) {
	tree := NewRedBlackTree()
	fillTree(tree)
	tests1 := map[data.Int]interface{}{
		data.Int(1): "x",
		data.Int(2): "b",
		data.Int(3): "c",
		data.Int(4): "d",
		data.Int(5): "e",
		data.Int(6): "f",
		data.Int(7): "g",
	}

	for k, v := range tests1 {
		actualValue := tree.Get(k)
		if actualValue == nil || actualValue != v {
			t.Errorf("Got %v expected %v", actualValue, v)
		}
	}
	actualValue := tree.Get(data.Int(8))
	if actualValue != nil {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
}

func TestRedBlackTreeRemove(t *testing.T) {
	tree := NewRedBlackTree()
	fillTree(tree)

	tree.Delete(data.Int(5))
	tree.Delete(data.Int(6))
	tree.Delete(data.Int(7))
	tree.Delete(data.Int(8))
	tree.Delete(data.Int(5))

	tests2 := map[data.Int]interface{}{
		data.Int(1): "x",
		data.Int(2): "b",
		data.Int(3): "c",
		data.Int(4): "d",
	}

	for k, v := range tests2 {
		actualValue := tree.Get(k)
		if actualValue == nil || actualValue != v {
			t.Errorf("Got %v expected %v", actualValue, v)
		}
	}

	for i := 5; i < 10; i++ {
		actualValue := tree.Get(data.Int(i))
		if actualValue != nil {
			t.Errorf("Got %v expected %v", actualValue, nil)
		}
	}

	tree.Delete(data.Int(1))
	tree.Delete(data.Int(4))
	tree.Delete(data.Int(2))
	tree.Delete(data.Int(3))
	tree.Delete(data.Int(2))
	tree.Delete(data.Int(2))

	for i := 0; i < 10; i++ {
		actualValue := tree.Get(data.Int(i))
		if actualValue != nil {
			t.Errorf("Got %v expected %v", actualValue, nil)
		}
	}

	if !tree.isEmpty() {
		t.Error("Tree is not empty")
	}
}

func TestRedBlackTreeMin(t *testing.T) {
	tree := NewRedBlackTree()

	if actualValue := tree.min(tree.root); actualValue != nil {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}

	fillTree(tree)

	if actualValue, expectedValue := tree.min(tree.root).key, data.Int(1); !actualValue.Equals(expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := tree.min(tree.root).value, "x"; actualValue.(string) != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	tree.Delete(data.Int(1))

	if actualValue, expectedValue := tree.min(tree.root).key, data.Int(2); !actualValue.Equals(expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := tree.min(tree.root).value, "b"; actualValue.(string) != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeCeilingAndFloor(t *testing.T) {
	tree := NewRedBlackTree()

	if node := tree.Floor(data.Int(0)); node != nil {
		t.Errorf("Got %v expected %v", node, nil)
	}
	if node := tree.Ceil(data.Int(0)); node != nil {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}

	fillTree(tree)

	if node := tree.Floor(data.Int(4)); node.(string) != "d" {
		t.Errorf("Got %v expected %v", node, "d")
	}
	if node := tree.Floor(data.Int(0)); node != nil {
		t.Errorf("Got %v expected %v", node, nil)
	}

	if node := tree.Ceil(data.Int(4)); node.(string) != "d" {
		t.Errorf("Got %v expected %v", node, "d")
	}
	if node := tree.Ceil(data.Int(8)); node != nil {
		t.Errorf("Got %v expected %v", node, nil)
	}
}

func benchmarkGet(b *testing.B, tree *RedBlackTree, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Get(data.Int(n))
		}
	}
}

func benchmarkPut(b *testing.B, tree *RedBlackTree, size int) {
	bd := new(string)
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Put(data.Int(n), bd)
		}
	}
}

func benchmarkRemove(b *testing.B, tree *RedBlackTree, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Delete(data.Int(n))
		}
	}
}

func BenchmarkRedBlackTreeGet100(b *testing.B) {
	b.StopTimer()
	bd := new(string)
	size := 100
	tree := NewRedBlackTree()
	for n := 0; n < size; n++ {
		tree.Put(data.Int(n), bd)
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet1000(b *testing.B) {
	b.StopTimer()
	bd := new(string)
	size := 1000
	tree := NewRedBlackTree()
	for n := 0; n < size; n++ {
		tree.Put(data.Int(n), bd)
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet10000(b *testing.B) {
	b.StopTimer()
	bd := new(string)
	size := 10000
	tree := NewRedBlackTree()
	for n := 0; n < size; n++ {
		tree.Put(data.Int(n), bd)
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet100000(b *testing.B) {
	b.StopTimer()
	bd := new(string)
	size := 100000
	tree := NewRedBlackTree()
	for n := 0; n < size; n++ {
		tree.Put(data.Int(n), bd)
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreePut100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewRedBlackTree()
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut1000(b *testing.B) {
	b.StopTimer()
	bd := new(string)
	size := 1000
	tree := NewRedBlackTree()
	for n := 0; n < size; n++ {
		tree.Put(data.Int(n), bd)
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut10000(b *testing.B) {
	b.StopTimer()
	bd := new(string)
	size := 10000
	tree := NewRedBlackTree()
	for n := 0; n < size; n++ {
		tree.Put(data.Int(n), bd)
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut100000(b *testing.B) {
	b.StopTimer()
	bd := new(string)
	size := 100000
	tree := NewRedBlackTree()
	for n := 0; n < size; n++ {
		tree.Put(data.Int(n), bd)
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreeRemove100(b *testing.B) {
	b.StopTimer()
	bd := new(string)
	size := 100
	tree := NewRedBlackTree()
	for n := 0; n < size; n++ {
		tree.Put(data.Int(n), bd)
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove1000(b *testing.B) {
	b.StopTimer()
	bd := new(string)
	size := 1000
	tree := NewRedBlackTree()
	for n := 0; n < size; n++ {
		tree.Put(data.Int(n), bd)
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove10000(b *testing.B) {
	b.StopTimer()
	bd := new(string)
	size := 10000
	tree := NewRedBlackTree()
	for n := 0; n < size; n++ {
		tree.Put(data.Int(n), bd)
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove100000(b *testing.B) {
	b.StopTimer()
	bd := new(string)
	size := 100000
	tree := NewRedBlackTree()
	for n := 0; n < size; n++ {
		tree.Put(data.Int(n), bd)
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}
