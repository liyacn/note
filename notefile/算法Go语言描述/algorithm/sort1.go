package algorithm

/* 冒泡排序
比较相邻的元素。如果第一个比第二个大，就交换他们两个。
对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。
针对所有的元素重复以上的步骤，除了最后一个。
持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
比较次数： C(n,2)=n(n-1)/2
时间复杂度O(n^2)，空间复杂度O(1)。
*/

func BubbleSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 优化，相同的循环次数，更少的交换次数
func BubbleSort2(arr []int) {
	end := len(arr) - 1
	for i := 0; i < end; i++ {
		for j := end; j > i; j-- {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

/* 选择排序
首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
重复上一步，直到所有元素均排序完毕。
时间复杂度O(n^2)，空间复杂度O(1)。
同样的时间复杂度，选择排序的性能上还是要略优于冒泡排序。
*/

func SelectionSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		m := i // 最小索引
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[m] {
				m = j
			}
		}
		if m != i {
			arr[m], arr[i] = arr[i], arr[m]
		}
	}
}

/* 插入排序
从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。
将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
时间复杂度O(n^2)，空间复杂度O(1)。
同样的时间复杂度，直接插入排序法比冒泡和选择排序的性能要好一些。
*/

func InsertionSort(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

/* 希尔排序，也称递减增量排序算法，是插入排序的一种更高效的改进版本。
选择一个增量序列t1，t2，……，tk，其中ti>tj, tk = 1；按增量序列个数k，对序列进行k趟排序；
每趟排序，根据对应的增量gap，将待排序列分割成若干长度为m的子序列，分别对各子表进行直接插入排序。
一般的初次取序列的一半为增量，以后每次减半，直到增量为1。
时间复杂度取决于gap增量，约为O(n^1.3)，空间复杂度O(1)。
*/

func ShellSort(arr []int) {
	length := len(arr)
	for gap := length >> 1; gap > 0; gap >>= 1 {
		for i := gap; i < length; i++ {
			for j := i - gap; j >= 0 && arr[j+gap] < arr[j]; j -= gap {
				arr[j], arr[j+gap] = arr[j+gap], arr[j]
			}
		}
	}
}

func ShellSort2(arr []int) {
	length := len(arr)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1 //动态计算间隔
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			for j := i - gap; j >= 0 && arr[j+gap] < arr[j]; j -= gap {
				arr[j], arr[j+gap] = arr[j+gap], arr[j]
			}
		}
		gap /= 3
	}
}

/* 归并排序是建立在归并操作上的一种有效的排序算法。是采用分治法（Divide and Conquer）的一个非常典型的应用。
申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
设定两个指针，最初位置分别为两个已经排序序列的起始位置；
比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
重复上一步直到某一指针达到序列尾；将另一序列剩下的所有元素直接复制到合并序列尾。
时间复杂度O(nlogn)，空间复杂度O(n)。
*/

// 相邻两段有序区间[left,mid)和[mid,right)归并排序
func merge(arr []int, left, mid, right int) {
	result := make([]int, 0, right-left)
	i, j := left, mid
	for i < mid && j < right {
		if arr[i] < arr[j] {
			result = append(result, arr[i])
			i++
		} else {
			result = append(result, arr[j])
			j++
		}
	}
	for i < mid {
		result = append(result, arr[i])
		i++
	}
	for j < right {
		result = append(result, arr[j])
		j++
	}
	//将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	copy(arr[left:right], result)
}

func MergeSort(arr []int, left, right int) { //前闭后开区间[left,right)
	if right-left < 2 {
		return
	}
	mid := (left + right) >> 1
	MergeSort(arr, left, mid)
	MergeSort(arr, mid, right)
	merge(arr, left, mid, right)
}

func MergeSort1(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	mid := length >> 1
	MergeSort1(arr[:mid])
	MergeSort1(arr[mid:])
	merge(arr, 0, mid, length)
}

func MergeSort2(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	for step := 1; step < length; step <<= 1 {
		left, mid, right := 0, step, step+step
		for right <= length {
			merge(arr, left, mid, right)
			left = right
			mid = left + step
			right = mid + step
		}
		if mid < length { //是否有右半段，此时right一定越界
			merge(arr, left, mid, length)
		}
	}
}

/* 小结
O(n^2)的排序效率：BubbleSort < BubbleSort2 < SelectionSort < InsertionSort
ShellSort与ShellSort2效率大致相当，MergeSort1与MergeSort2效率大致相当。
中等规模排序效率ShellSort优于MergeSort。大规模数据排序效率MergeSort优于ShellSort。
*/
