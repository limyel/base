package linkedlist

import (
	"reflect"
	"testing"
)

// toSlice 是一个辅助函数，用于将链表转换为切片，方便测试断言
func (ll *LinkedList) toSlice() []interface{} {
	var result []interface{}
	n := ll.head
	for n != nil {
		result = append(result, n.value)
		n = n.next
	}
	return result
}

func TestLinkedList(t *testing.T) {
	t.Run("Append", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)

		expected := []interface{}{1, 2, 3}
		if !reflect.DeepEqual(ll.toSlice(), expected) {
			t.Errorf("Append() failed. Expected %v, got %v", expected, ll.toSlice())
		}
	})

	t.Run("Prepend", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Append(2)
		ll.Prepend(1)
		ll.Append(3)
		ll.Prepend(0)

		expected := []interface{}{0, 1, 2, 3}
		if !reflect.DeepEqual(ll.toSlice(), expected) {
			t.Errorf("Prepend() failed. Expected %v, got %v", expected, ll.toSlice())
		}
	})

	t.Run("Delete", func(t *testing.T) {
		// 子测试：删除中间节点
		t.Run("delete middle node", func(t *testing.T) {
			ll := NewLinkedList()
			ll.Append(1)
			ll.Append(2)
			ll.Append(3)
			err := ll.Delete(2)
			if err != nil {
				t.Fatalf("Delete() returned an unexpected error: %v", err)
			}
			expected := []interface{}{1, 3}
			if !reflect.DeepEqual(ll.toSlice(), expected) {
				t.Errorf("Delete() failed. Expected %v, got %v", expected, ll.toSlice())
			}
		})

		// 子测试：删除头节点
		t.Run("delete head node", func(t *testing.T) {
			ll := NewLinkedList()
			ll.Append(1)
			ll.Append(2)
			ll.Append(3)
			err := ll.Delete(1)
			if err != nil {
				t.Fatalf("Delete() returned an unexpected error: %v", err)
			}
			expected := []interface{}{2, 3}
			if !reflect.DeepEqual(ll.toSlice(), expected) {
				t.Errorf("Delete() failed. Expected %v, got %v", expected, ll.toSlice())
			}
		})

		// 子测试：删除尾节点
		t.Run("delete tail node", func(t *testing.T) {
			ll := NewLinkedList()
			ll.Append(1)
			ll.Append(2)
			ll.Append(3)
			err := ll.Delete(3)
			if err != nil {
				t.Fatalf("Delete() returned an unexpected error: %v", err)
			}
			expected := []interface{}{1, 2}
			if !reflect.DeepEqual(ll.toSlice(), expected) {
				t.Errorf("Delete() failed. Expected %v, got %v", expected, ll.toSlice())
			}
		})

		// 子测试：删除不存在的节点
		t.Run("delete non-existent node", func(t *testing.T) {
			ll := NewLinkedList()
			ll.Append(1)
			ll.Append(2)
			err := ll.Delete(4)
			if err == nil {
				t.Errorf("Delete() should have returned an error, but it didn't")
			}
			expected := []interface{}{1, 2}
			if !reflect.DeepEqual(ll.toSlice(), expected) {
				t.Errorf("Delete() a non-existent value should not change the list. Expected %v, got %v", expected, ll.toSlice())
			}
		})

		// 子测试：从空链表中删除
		t.Run("delete from empty list", func(t *testing.T) {
			ll := NewLinkedList()
			err := ll.Delete(1)
			if err == nil {
				t.Errorf("Delete() from an empty list should have returned an error, but it didn't")
			}
		})

		// 子测试：删除后链表为空
		t.Run("delete until list is empty", func(t *testing.T) {
			ll := NewLinkedList()
			ll.Append(1)
			err := ll.Delete(1)
			if err != nil {
				t.Fatalf("Delete() returned an unexpected error: %v", err)
			}
			if ll.head != nil {
				t.Errorf("List should be empty, but head is not nil. Got %v", ll.toSlice())
			}
		})
	})
}
