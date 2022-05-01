package datastruct

// 栈和队列都可以用数组或链表来实现

type ArrayStack struct {
	items []any // 底层切片
	size  int   // 栈的元素个数
}

func (s *ArrayStack) Size() int {
	return s.size
}

func (s *ArrayStack) IsEmpty() bool {
	return s.size == 0
}

func (s *ArrayStack) Clear() {
	s.items = s.items[:0]
	s.size = 0
}

func (s *ArrayStack) Push(v any) {
	s.items = append(s.items, v)
	s.size++
}

func (s *ArrayStack) Pop() any {
	if s.size == 0 {
		return nil
	}
	s.size--
	top := s.items[s.size]
	s.items = s.items[:s.size]
	return top
}

func (s *ArrayStack) Top() any {
	if s.size == 0 {
		return nil
	}
	return s.items[s.size-1]
}

func (s *ArrayStack) Bottom() any {
	if s.size == 0 {
		return nil
	}
	return s.items[0]
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
