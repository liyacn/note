package datastruct

// 栈和队列都可以用数组或链表来实现

type ArrayStack []any

func (s ArrayStack) Size() int {
	return len(s)
}

func (s ArrayStack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *ArrayStack) Clear() {
	*s = nil
}

func (s *ArrayStack) Push(v any) {
	*s = append(*s, v)
}

func (s *ArrayStack) Pop() any {
	size := s.Size()
	if size == 0 {
		return nil
	}
	end := size - 1
	top := (*s)[end]
	(*s)[end] = nil // don't stop the GC from reclaiming the item eventually
	*s = (*s)[:end]
	return top
}

func (s ArrayStack) Top() any {
	size := s.Size()
	if size == 0 {
		return nil
	}
	return s[size-1]
}

func (s ArrayStack) Bottom() any {
	if s.Size() == 0 {
		return nil
	}
	return s[0]
}

type ListNode struct {
	Val  any
	Next *ListNode
}

type LinkedQueue struct {
	head, tail *ListNode // 链表头尾
	size       int       // 队列长度
}

func (q *LinkedQueue) Size() int {
	return q.size
}

func (q *LinkedQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedQueue) Clear() {
	q.head = nil
	q.tail = nil
	q.size = 0
}

func (q *LinkedQueue) Push(v any) {
	node := &ListNode{Val: v}
	if q.size == 0 {
		q.head = node
		q.tail = node
	} else {
		q.tail.Next = node
		q.tail = q.tail.Next
	}
	q.size++
}

func (q *LinkedQueue) Pop() any {
	if q.size == 0 {
		return nil
	}
	val := q.head.Val
	q.head = q.head.Next
	if q.size == 1 {
		q.tail = nil
	}
	q.size--
	return val
}

func (q *LinkedQueue) Front() any {
	if q.size == 0 {
		return nil
	}
	return q.head.Val
}

func (q *LinkedQueue) Back() any {
	if q.size == 0 {
		return nil
	}
	return q.tail.Val
}

// Go中还能使用带缓冲的channel实现单向有锁队列。
