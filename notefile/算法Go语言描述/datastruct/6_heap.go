package datastruct

import (
	"container/heap"
	"fmt"
)

/*
最大堆：
     9
    /\
   5  7
  /\
 1  3
数组表示：[9,5,7,1,3]

下标特点：
根的下标为0，最后一个元素的下标为size-1。
任一结点下标为i，其父结点p下标为(i-1)/2，左子结点l下标为2*i+1，右子结点下r标为2*i+2(即l+1)。

最大堆实现细节(两个操作)：
	push：向堆中插入数据时，首先在堆的末尾插入数据，如果该数据比父结点还大，那么交换，然后不断向上提升，直到没有大小颠倒为止。
	pop：从堆中删除最大值时，首先把最后一个值复制到根结点上，并且删除最后一个数值，
		然后和子结点比较，如果值小于子结点则交换，然后不断向下交换，直到没有大小颠倒为止。
		在向下交换过程中，如果有两个子结点都大于自己，就选择较大的。
最大堆有两个核心操作，一个是上浮，一个是下沉，分别对应push和pop。

最大堆从构建到移除，总的平均时间复杂度是：O(nlogn)。

堆数据结构可以用来实现最短路径算法中的优先队列，从而提高算法的效率。
*/

type MaxHeap struct {
	items []int
	size  int
}

func (h *MaxHeap) Size() int {
	return h.size
}

func (h *MaxHeap) IsEmpty() bool {
	return h.size == 0
}

// 最大堆插入元素
func (h *MaxHeap) Push(v int) {
	h.items = append(h.items, v)
	i := h.size // 要插入结点的下标
	for i > 0 {
		p := (i - 1) / 2
		if v <= h.items[p] {
			break // 如果插入的值小于等于父结点，可以直接退出循环
		}
		h.items[i] = h.items[p]
		i = p // 否则将父结点与该结点互换，然后向上翻转，将最大的元素一直往上推
	}
	h.items[i] = v //插入值放在不会再翻转的位置
	h.size++
}

// 最大堆移除根元素，也就是最大的元素
func (h *MaxHeap) Pop() (int, bool) {
	if h.size == 0 {
		return 0, false
	}
	ret := h.items[0] //取出根元素
	h.size--
	v := h.items[h.size] //最后一个值先取出来
	i := 0
	for {
		a := 2*i + 1 // 左子结点下标
		if a >= h.size {
			break // 左子结点越界，表示没有左子结点，也就没有右子结点
		}
		if b := a + 1; b < h.size && h.items[b] > h.items[a] {
			a = b // 此时a为两个子结点中较大结点的下标
		}
		if v >= h.items[a] {
			break // 父结点的值都不小于两个子结点了，不需要再向下翻转
		}
		h.items[i] = h.items[a] // 较大的子结点与父结点交换
		i = a                   // 继续向下翻转
	}
	h.items[i] = v             // 将最后一个元素的值放在不会再翻转的位置
	h.items = h.items[:h.size] // 删除最后一个空位
	return ret, true
}

func TestMaxHeap() {
	list := []int{5, 7, 6, 0, 8, 9, 3, 1, 4, 2}
	h := new(MaxHeap)
	for _, v := range list {
		h.Push(v) // 无序添加进堆
	}
	for h.Size() > 0 {
		fmt.Println(h.Pop()) // 从大到小弹出
	}
}

// "container/heap"包提供了对任意类型（实现了heap.Interface接口）的堆操作。

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func IntHeapSort(arr []int) {
	h := IntHeap(arr)
	heap.Init(&h)
	for h.Len() > 0 {
		heap.Pop(&h)
	}
}
