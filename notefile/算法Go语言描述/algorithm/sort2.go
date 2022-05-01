package algorithm

/*堆排序是指利用堆这种数据结构所设计的一种排序算法。
大顶堆用于实现升序排列，小顶堆用于实现降序排列。
堆排序的平均时间复杂度为 Ο(nlogn)，空间复杂度为O(1)。
*/

func HeapSort(arr []int) []int {
	count := len(arr)
	for i := count/2 - 1; i >= 0; i-- {
		down(arr, i, count)
	}
	for i := count - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		down(arr, 0, i)
	}
	return arr
}

func down(arr []int, root, count int) {
	child := root*2 + 1
	for child < count {
		if r := child + 1; r < count && arr[child] < arr[r] {
			child = r
		}
		if arr[root] >= arr[child] {
			return
		}
		arr[root], arr[child] = arr[child], arr[root]
		root = child // 继续往下沉
		child = root*2 + 1
	}
}

/*快速排序使用分治法（Divide and conquer）策略来把一个串行（list）分为两个子串行（sub-lists）。
从数列中挑出一个元素，称为 “基准”（pivot）;
重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
递归的最底部情形，是数列的大小是零或一，也就是永远都已经被排序好了。
时间复杂度O(nlogn)，空间复杂度O(logn)。
*/

//单路快排
func QuickSort1(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	pivot := arr[0]
	p := 0
	for i := 1; i < length; i++ {
		if arr[i] < pivot {
			p++
			arr[i], arr[p] = arr[p], arr[i]
		}
	}
	arr[0], arr[p] = arr[p], arr[0]
	QuickSort1(arr[:p])
	QuickSort1(arr[p+1:])
}
func QuickSortOne(arr []int, left, right int) { //前闭后开区间[left,right)
	if right-left < 2 {
		return
	}
	pivot := arr[left]
	p := left
	for i := left + 1; i < right; i++ {
		if arr[i] < pivot {
			p++
			arr[i], arr[p] = arr[p], arr[i]
		}
	}
	arr[left], arr[p] = arr[p], arr[left]
	QuickSortOne(arr, 0, p)
	QuickSortOne(arr, p+1, right)
}

//双路快排
func QuickSort2(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	pivot := arr[0]
	i, j := 1, length-1
	for {
		for i < length && arr[i] < pivot {
			i++
		}
		for j > 0 && arr[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[0], arr[j] = arr[j], arr[0]
	QuickSort2(arr[:j])
	QuickSort2(arr[j+1:])
}
func QuickSortTwo(arr []int, left, right int) { //前闭后开区间[left,right)
	if right-left < 2 {
		return
	}
	pivot := arr[left]
	i, j := left+1, right-1
	for {
		for i < right && arr[i] < pivot {
			i++
		}
		for j > left && arr[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[left], arr[j] = arr[j], arr[left]
	QuickSortTwo(arr, left, j)
	QuickSortTwo(arr, j+1, right)
}

//三路快排
func QuickSort3(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	pivot := arr[0]
	l, i, r := 0, 1, length-1
	for i <= r {
		if arr[i] > pivot {
			arr[i], arr[r] = arr[r], arr[i]
			r--
		} else if arr[i] < pivot {
			arr[i], arr[l] = arr[l], arr[i]
			i++
			l++
		} else {
			i++
		}
	}
	//循环结束后，[l,r]区间内的值均为pivot
	QuickSort3(arr[:l])
	QuickSort3(arr[r+1:])
}
func QuickSortThree(arr []int, left, right int) { //前闭后开区间[left,right)
	if right-left < 2 {
		return
	}
	pivot := arr[left]
	l, i, r := left, left+1, right-1
	for i <= r {
		if arr[i] > pivot {
			arr[i], arr[r] = arr[r], arr[i]
			r--
		} else if arr[i] < pivot {
			arr[i], arr[l] = arr[l], arr[i]
			i++
			l++
		} else {
			i++
		}
	}
	//循环结束后，[l,r]区间内的值均为pivot
	QuickSortThree(arr, left, l)
	QuickSortThree(arr, r+1, right)
}

/*
平均时间上，堆排序的时间常数比快排要大一些，因此通常会慢一些。

三种快排效率比较：
单路快排在重复较少的情况下表现最佳，
随着重复率增加双路快排优于单路，
大量重复的情况下三路快排优势最大。

pdqsort（Pattern-defeating quicksort）是一种融合插入排序、堆排序和优化后的快排的新型排序算法，在rust和go1.19中采用。
对于短序列（go中长度12以内），使用插入排序。
其他情况，使用快速排序保证整体性能。
当快速排序表象不佳，使用堆排序保证最坏情况小时间复杂度仍然为O(nlogn)。

算法复杂度比较：
               Best      Avg       Worst
InsertionSort  O(n)      O(n^2)    O(n^2)
QuickSort      O(nlogn)  O(nlogn)  O(n^2)
HeapSort       O(nlogn)  O(nlogn)  O(nlogn)
pdqsort        O(n)      O(nlogn)  O(nlogn)
*/
