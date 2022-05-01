package datastruct

import (
	"container/list"
	"container/ring"
	"fmt"
)

// 最简单的单向链表
type ListNode struct {
	Val  any
	Next *ListNode
}

func NewLinkedList(val ...any) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range val {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func (list *ListNode) Print() {
	for cur := list; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, "->")
	}
	fmt.Println("nil")
}

// "container/list"包实现了一个双向链表
func ExampleList() {
	l := list.New()                                       // 创建一个链表
	a := l.PushBack("A")                                  // 链表尾插入新元素
	b := l.PushFront("B")                                 // 链表头插入新元素
	l.InsertBefore("C", a)                                // 节点前插入元素
	l.InsertAfter("D", b)                                 // 节点后插入元素
	fmt.Println(l.Len(), l.Front().Value, l.Back().Value) // 4 B A

	for e := l.Front(); e != nil; e = e.Next() { // 顺序遍历
		fmt.Print(e.Value, " ") // B D C A
	}
	fmt.Println()

	val := l.Remove(a)        // 删除一个元素
	fmt.Println(l.Len(), val) // 3 A

	for e := l.Back(); e != nil; e = e.Prev() { //逆序遍历
		fmt.Print(e.Value, " ") // C D B
	}
	fmt.Println()
}

// "container/ring"包实现了一个双向环链表
func ExampleRing() {
	r := ring.New(6)
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next() // 顺向移动
	}
	for j := 0; j < r.Len(); j++ {
		r = r.Prev()            // 逆向移动
		fmt.Print(r.Value, " ") // 5 4 3 2 1 0
	}
	fmt.Println()

	r = r.Move(2) // 大于0顺向移动，小于0逆向移动
	r.Do(func(i any) {
		fmt.Print(i, " ") // 2 3 4 5 0 1
	})
	fmt.Println()
}

// Golang中还能使用带缓冲的channel实现单向有锁队列。
