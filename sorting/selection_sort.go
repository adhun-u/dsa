package sorting

func SelectionSort(nums []int) []int {

	for i := range nums {
		minIndex := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		temp := nums[i]
		nums[i] = nums[minIndex]
		nums[minIndex] = temp
	}
	return nums
}
