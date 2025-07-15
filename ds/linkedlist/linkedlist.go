package linkedlist

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

type LinkedList struct {
	head *Node
}

func newNode(value interface{}) *Node {
	return &Node{
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

	if ll.head.value == value {
		ll.head = ll.head.next
		return nil
	}

	current := ll.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
		return nil
	} else {
		return fmt.Errorf("value %v not exists", value)
	}
}

func (ll *LinkedList) Print() {
	n := ll.head
	for n != nil {
		fmt.Printf("%v -> ", n.value)
		n = n.next
	}
	fmt.Println("nil")
}
