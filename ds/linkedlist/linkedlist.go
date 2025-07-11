package linkedlist

import "fmt"

type node struct {
	value interface{}
	next  *node
}

type LinkedList struct {
	head *node
}

func newNode(value interface{}) *node {
	return &node{
		value: value,
		next:  nil,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
	}
}

// Append 往链表尾部追加
// 1、如果链表为空，直接插入
// 2、如果链表不为空，找到最后一个节点，插入
func (ll *LinkedList) Append(value interface{}) {
	n := newNode(value)
	if ll.head == nil {
		ll.head = n
		return
	}

	nextNode := ll.head
	for nextNode.next != nil {
		nextNode = nextNode.next
	}

	nextNode.next = n
}

func (ll *LinkedList) Prepend(value interface{}) {
	n := newNode(value)
	n.next = ll.head
	ll.head = n
}

func (ll *LinkedList) Delete(value interface{}) error {
	if ll.head == nil {
		return fmt.Errorf("linkedlist is empty")
	}

	n := ll.head
	var preNode *node
	for n != nil {
		if n.value == value {
			break
		}

		preNode = n
		n = n.next
	}

	if n == nil {
		return fmt.Errorf("value %v not exists", value)
	}

	// 要移除的是头节点
	if preNode == nil {
		ll.head = n.next
		return nil
	}

	preNode.next = n.next
	return nil
}

func (ll *LinkedList) Print() {
	n := ll.head
	for n != nil {
		fmt.Printf("%v -> ", n.value)
		n = n.next
	}
	fmt.Println("nil")
}
