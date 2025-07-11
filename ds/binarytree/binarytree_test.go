package binarytree

import (
	"reflect"
	"testing"
)

// newTestTree 创建一个用于测试的树
//
//	    A
//	   / \
//	  B   C
//	 /   / \
//	D   E   F
func newTestTree() *BinaryTree {
	nodeA := &node{value: "A"}
	nodeB := &node{value: "B"}
	nodeC := &node{value: "C"}
	nodeD := &node{value: "D"}
	nodeE := &node{value: "E"}
	nodeF := &node{value: "F"}

	nodeA.left = nodeB
	nodeA.right = nodeC
	nodeB.left = nodeD
	nodeC.left = nodeE
	nodeC.right = nodeF

	return &BinaryTree{Root: nodeA}
}

func TestBinaryTree_PreOrder(t *testing.T) {
	tree := newTestTree()
	expected := []interface{}{"A", "B", "D", "C", "E", "F"}
	result := tree.PreOrder()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PreOrder failed. Expected %v, got %v", expected, result)
	}
}

func TestBinaryTree_InOrder(t *testing.T) {
	tree := newTestTree()
	expected := []interface{}{"D", "B", "A", "E", "C", "F"}
	result := tree.InOrder()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InOrder failed. Expected %v, got %v", expected, result)
	}
}

func TestBinaryTree_PostOrder(t *testing.T) {
	tree := newTestTree()
	expected := []interface{}{"D", "B", "E", "F", "C", "A"}
	result := tree.PostOrder()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PostOrder failed. Expected %v, got %v", expected, result)
	}
}
