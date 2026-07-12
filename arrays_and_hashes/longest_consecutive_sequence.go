package arraysandhashes

func LongestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]struct{})

	for i := range nums {
		num := nums[i]
		set[num] = struct{}{}
	}

	long := 0

	for i := range nums {

		num := nums[i]

		if _, exists := set[num-1]; exists {
			continue
		}

		count := 1
		currentNum := num

		for {

			if _, exists := set[currentNum+1]; !exists {
				break
			}

			count++
			currentNum++
		}

		if count > long {
			long = count
		}

	}

	return long
}
