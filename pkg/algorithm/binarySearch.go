package algorithm

// 二分查找数组
//1、若target在数组中则返回下标（若数组中存在多个target，则随机返回某个target下标）
//2、若target不存在数组中，返回-1

func BinarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for {
		if high < low {
			break
		}
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
