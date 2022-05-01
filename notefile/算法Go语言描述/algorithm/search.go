package algorithm

/*
折半计算 mid=(low+higt)/2
如需防止溢出可变形为 mid=low+(high-low)/2
*/

//二分查找/折半查找，前提是线性表中的记录必须是有序（通常从小到大有序）。时间复杂度O(logn)。
func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) >> 1
		if target < arr[mid] {
			high = mid - 1
		} else if target > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

//查找最先出现的位置
func BinarySearchBegin(arr []int, target int) int {
	end := len(arr) - 1
	low, high := 0, end
	for low <= high {
		mid := (low + high) >> 1
		if target > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1 //小于或等于都收缩右边界
		}
	}
	if low > end || arr[low] != target {
		return -1
	}
	return low
}

//查找最后出现的位置
func BinarySearchEnd(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) >> 1
		if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1 //大于或等于都收缩左边界
		}
	}
	if high < 0 || arr[high] != target {
		return -1
	}
	return high
}

//查找目标首尾出现的位置
func BinarySearchRange(arr []int, target int) (int, int) {
	end := len(arr) - 1
	//先找左边界low
	low, high := 0, end
	for low <= high {
		mid := (low + high) >> 1
		if target > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low > end || arr[low] != target {
		return -1, -1
	}
	//再找右边界r，有左边界必然右边界可求，不用判断越界
	l, r := low, end
	for l <= r {
		m := (l + r) >> 1
		if target < arr[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return low, r
}

//插值查找基于二分查找，将查找点的选择改进为自适应选择，提高查找效率，时间复杂度同为O(logn)。
//对于表长较大，而关键字分布又比较均匀的查找表来说，插值查找的平均性能比折半查找要好。
func InsertSearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (target-arr[low])/(arr[high]-arr[low])*(high-low)
		if target > arr[mid] {
			low = mid + 1
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
