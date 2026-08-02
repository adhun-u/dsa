package sorting

func BubbleSort(nums []int) []int {
	for i := range nums {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}

	return nums
}
