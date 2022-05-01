package algorithm

/* 归并排序是建立在归并操作上的排序算法，是采用分治法的一个非常典型的应用。
将已有序的子序列合并，得到完全有序的序列；两个有序表合并成一个有序表，称为二路归并。
申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
设定两个指针，最初位置分别为两个已经排序序列的起始位置；
比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
重复上一步直到某一指针达到序列尾；将另一序列剩下的所有元素直接复制到合并序列尾。
时间复杂度O(nlogn)，辅助空间复杂度O(n)。
方法1自顶向下递归，递归栈的空间复杂度O(nlogn)。
方法2自底向上非递归实现，无需额外递归栈空间。
*/

// 二路归并nums相邻两段有序区间[l,m)和[m,r)，使用tmp辅助空间
func merge(nums, tmp []int, l, m, r int) {
	copy(tmp[l:r], nums[l:r]) //先复制到辅助空间，再将归并结果直接放原数组
	i, j, k := l, m, l
	for i < m && j < r {
		if tmp[j] < tmp[i] {
			nums[k] = tmp[j]
			j++
		} else { //相等值先将左段放到结果，以免破坏稳定性
			nums[k] = tmp[i]
			i++
		}
		k++
	}
	if i < m { //左段有剩余需复制
		copy(nums[k:r], tmp[i:m])
	}
	// j<r时k=j，nums[k:r]与tmp[j:r]相同，右段有剩余无需复制
}
func mergeSort(nums, tmp []int, l, r int) { //前闭后开区间[l,r)
	if r-l < 2 {
		return
	}
	m := (l + r) >> 1
	mergeSort(nums, tmp, l, m)
	mergeSort(nums, tmp, m, r)
	merge(nums, tmp, l, m, r)
}
func MergeSort1(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	tmp := make([]int, n)
	mergeSort(nums, tmp, 0, n)
}
func MergeSort2(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	tmp := make([]int, n)
	for step := 1; step < n; step <<= 1 {
		l, m, r := 0, step, step<<1
		for r <= n {
			merge(nums, tmp, l, m, r)
			l = r
			m = l + step
			r = m + step
		}
		if m < n { //是否有右半段，此时r一定越界
			merge(nums, tmp, l, m, n)
		}
	}
}

/* 堆排序是指利用堆这种数据结构所设计的一种排序算法。
大顶堆用于实现升序排列，小顶堆用于实现降序排列。
时间复杂度为 Ο(nlogn)，空间复杂度为O(1)。
*/

func HeapSort(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		down(nums, i, n)
	}
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		down(nums, 0, i)
	}
}
func down(nums []int, i, n int) {
	j := i*2 + 1
	for j < n {
		if r := j + 1; r < n && nums[j] < nums[r] {
			j = r
		}
		if nums[i] >= nums[j] {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i = j
		j = i*2 + 1
	}
}

/* 快速排序使用分治法策略来把一个串行分为两个子串行。
从数列中挑出一个元素，称为“基准”（pivot）;
重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
递归的最底部情形，是数列的大小是零或一，也就是永远都已经被排序好了。
平均时间复杂度O(nlogn)，最坏的情况下会退化为O(n^2)；原地排序，递归栈空间复杂度O(logn)。
单路快排实现最简单，双路快排一般情况下性能较好，三路快排在处理大量重复元素时性能更优。
*/

func QuickSort1(nums []int) { //单路快排
	n := len(nums)
	if n < 2 {
		return
	}
	pivot := nums[0]
	j := 0
	for i := 1; i < n; i++ {
		if nums[i] < pivot {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[0], nums[j] = nums[j], nums[0]
	QuickSort1(nums[:j])
	QuickSort1(nums[j+1:])
}
func QuickSort2(nums []int) { //双路快排
	n := len(nums)
	if n < 2 {
		return
	}
	pivot := nums[0]
	i, j := 1, n-1
	for {
		for i < n && nums[i] < pivot {
			i++
		}
		for j > 0 && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[0], nums[j] = nums[j], nums[0]
	QuickSort2(nums[:j])
	QuickSort2(nums[j+1:])
}
func QuickSort3(nums []int) { //三路快排
	n := len(nums)
	if n < 2 {
		return
	}
	pivot := nums[0]
	l, i, r := 0, 1, n-1
	for i <= r {
		if nums[i] > pivot {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		} else if nums[i] < pivot {
			nums[i], nums[l] = nums[l], nums[i]
			i++
			l++
		} else {
			i++
		}
	} //循环结束后，[l,r]区间内的值均为pivot
	QuickSort3(nums[:l])
	QuickSort3(nums[r+1:])
}
