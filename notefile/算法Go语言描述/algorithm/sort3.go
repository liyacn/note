package algorithm

/*
稳定：冒泡、插入、归并
不稳定：选择排序、希尔排序、堆排序、快速排序

小规模数据使用插入排序性能最优，大规模数据一般选用分治策略的排序算法（快速排序平均性能最优，归并排序可满足稳定性要求）。
为避免快速排序划分极度不均导致性能退化，常用堆排版作为后备算法。
对于链表排序，所有需要用到非邻位交换的排序算法都不适用，可用插入、归并。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	l := head
	for r := l.Next; r != nil; r = l.Next {
		if r.Val < l.Val {
			pre := dummy
			for pre.Next.Val <= r.Val {
				pre = pre.Next
			}
			l.Next = r.Next
			r.Next = pre.Next
			pre.Next = r
		} else {
			l = l.Next
		}
	}
	return dummy.Next
}

func mergeList(l, r *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for l != nil && r != nil {
		if r.Val < l.Val {
			tail.Next = r
			r = r.Next
		} else {
			tail.Next = l
			l = l.Next
		}
		tail = tail.Next
	}
	if l == nil {
		tail.Next = r
	} else {
		tail.Next = l
	}
	return dummy.Next
}

func MergeSortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	l := MergeSortList1(head)
	r := MergeSortList1(mid)
	return mergeList(l, r)
}

// 指定位置剪断链表并返回后半部分
func cutList(node *ListNode, step int) *ListNode {
	for step > 1 && node != nil {
		step--
		node = node.Next
	}
	if node == nil {
		return nil
	}
	right := node.Next
	node.Next = nil //剪断
	return right
}

func MergeSortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := 2 //链表长度
	for cur := head.Next.Next; cur != nil; cur = cur.Next {
		n++
	}
	dummy := &ListNode{Next: head}
	for step := 1; step < n; step <<= 1 {
		tail := dummy
		cur := dummy.Next
		for cur != nil {
			left := cur
			right := cutList(left, step)
			cur = cutList(right, step)
			for left != nil && right != nil {
				if right.Val < left.Val {
					tail.Next = right
					right = right.Next
				} else {
					tail.Next = left
					left = left.Next
				}
				tail = tail.Next
			}
			if left == nil {
				tail.Next = right
			} else {
				tail.Next = left
			}
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}
	return dummy.Next
}
