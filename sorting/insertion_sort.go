package sorting

func InsertSort(nums []int) []int {

	for i := 1; i < len(nums); i++ {
		tmp := nums[i]
		backIndex := i - 1

		if nums[backIndex] > tmp {
			for backIndex >= 0 && nums[backIndex] > tmp {
				back := nums[backIndex]
				nums[backIndex+1] = back
				nums[backIndex] = tmp
				backIndex--
			}
		}
	}
	return nums
}
