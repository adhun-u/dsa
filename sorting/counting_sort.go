package sorting

func CountingSort(nums []int) []int {

	max := nums[0]

	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}

	count := make([]int, max+1)

	for i := range nums {
		count[nums[i]]++
	}

	index := 0
	for value, fr := range count {
		for fr > 0 {
			nums[index] = value
			index++
			fr--
		}
	}

	return nums
}
