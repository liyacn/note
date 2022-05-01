package algorithm

/* 冒泡排序
比较相邻的元素。如果第一个比第二个大，就交换他们两个。
对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。
针对所有的元素重复以上的步骤，除了最后一个。
持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
时间复杂度O(n^2)，空间复杂度O(1)。
方法1的比较次数n(n-1)/2，交换次数0~n(n-1)/2；
方法2和3将最少比较次数降到n-1。
*/

func BubbleSort1(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
func BubbleSort2(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		flag := false
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
func BubbleSort3(nums []int) {
	for i := len(nums) - 1; i > 0; {
		lastSwap := 0
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				lastSwap = j
			}
		}
		i = lastSwap
	}
}

/* 选择排序
首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
重复上一步，直到所有元素均排序完毕。
时间复杂度O(n^2)，空间复杂度O(1)。
同样的时间复杂度，选择排序的性能上还是要略优于冒泡排序。
方法1比较次数n(n-1)/2，交换次数0~n-1；
方法2比较次数约为偶数(n^2)/4、奇数(n^2-1)/4，交换次数0~n；
双向选择排序通过减少比较次数，在一定程度上提高了排序效率，尤其是对于大规模数据。
*/

func SelectionSort1(nums []int) {
	r := len(nums) - 1
	for i := 0; i < r; i++ {
		m := i // 最小索引
		for j := i + 1; j <= r; j++ {
			if nums[j] < nums[m] {
				m = j
			}
		}
		if m != i {
			nums[m], nums[i] = nums[i], nums[m]
		}
	}
}
func SelectionSort2(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		minIndex := left
		maxIndex := left
		for i := left + 1; i <= right; i++ {
			if nums[i] < nums[minIndex] {
				minIndex = i
			}
			if nums[i] > nums[maxIndex] {
				maxIndex = i
			}
		}
		if minIndex != left {
			nums[minIndex], nums[left] = nums[left], nums[minIndex]
		}
		if maxIndex == left {
			maxIndex = minIndex
		}
		if maxIndex != right {
			nums[maxIndex], nums[right] = nums[right], nums[maxIndex]
		}
		left++
		right--
	}
}

/* 插入排序
从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。
将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
时间复杂度O(n^2)，空间复杂度O(1)。
同样的时间复杂度，直接插入排序法比冒泡和选择排序的性能要好一些。
比较次数n-1~n(n-1)/2，移动次数0~n(n-1)/2
*/

func InsertionSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		k := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > k {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = k
	}
}

/* 希尔排序，也称递减增量排序算法，是插入排序的一种更高效的改进版本。
选择一个增量序列t1，t2，……，tk，其中ti>tj, tk = 1；按增量序列个数k，对序列进行k趟排序；
每趟排序，根据对应的增量step，将待排序列分割成若干长度为m的子序列，分别对各子表进行直接插入排序。
时间复杂度取决于step增量，空间复杂度O(1)。
时间复杂度最好情况接近O(nlogn)，最坏情况接近O(n^2)。
方法1使用折半序列，平均时间复杂度约为O(n^1.3)；
方法2使用Knuth序列，平均时间复杂度约为O(n^1.25)。
*/

func ShellSort1(nums []int) {
	n := len(nums)
	for step := n >> 1; step > 0; step >>= 1 {
		for i := step; i < n; i++ {
			for j := i - step; j >= 0 && nums[j+step] < nums[j]; j -= step {
				nums[j], nums[j+step] = nums[j+step], nums[j]
			}
		}
	}
}
func ShellSort2(nums []int) {
	n := len(nums)
	step := 1
	for step < n/3 {
		step = step*3 + 1
	}
	for step > 0 {
		for i := step; i < n; i++ {
			for j := i - step; j >= 0 && nums[j+step] < nums[j]; j -= step {
				nums[j], nums[j+step] = nums[j+step], nums[j]
			}
		}
		step /= 3
	}
}
