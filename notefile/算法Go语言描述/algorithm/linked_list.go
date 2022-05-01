package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表
func ReverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		head, head.Next, pre = head.Next, pre, head
	}
	return pre
}

//递归反转
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}

//插入排序
func InsertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	l := head
	for r := l.Next; r != nil; r = l.Next {
		if r.Val < l.Val {
			pre := r.Next
			cur := dummy
			for cur.Next.Val < r.Val {
				cur = cur.Next
			}
			r.Next = cur.Next
			cur.Next = r
			l.Next = pre
		} else {
			l = l.Next
		}
	}
	return dummy.Next
}

//自下而上归并排序
func MergeSortList(head *ListNode) *ListNode {
	length := 0 //链表长度
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}
	dummy := &ListNode{Next: head}
	for step := 1; step < length; step <<= 1 {
		tail := dummy
		cur := tail.Next
		for cur != nil {
			left := cur
			right := cutList(left, step) //分割链表
			cur = cutList(right, step)   //剩余部分
			//合并两个有序链表
			for left != nil && right != nil {
				if left.Val < right.Val {
					tail.Next = left
					left = left.Next
				} else {
					tail.Next = right
					right = right.Next
				}
				tail = tail.Next
			}
			if left == nil {
				tail.Next = right
			} else {
				tail.Next = left
			}
			//tail更新到末尾
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}
	return dummy.Next
}

//指定位置剪断链表并返回后半部分头结点
func cutList(node *ListNode, step int) *ListNode {
	for node != nil && step > 1 {
		node = node.Next
		step--
	}
	if node == nil {
		return nil
	}
	right := node.Next
	node.Next = nil //剪断
	return right
}
