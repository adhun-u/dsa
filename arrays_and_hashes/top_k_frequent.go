package arraysandhashes

func TopKFrequent(nums []int, k int) []int {

	if k == 0 {
		return []int{}
	} else if k >= len(nums) {
		return nums
	}

	frqCountMap := map[int]int{}
	bucket := make([][]int, len(nums)+1)
	output := []int{}

	for i := range nums {

		num := nums[i]
		frqCountMap[num]++

	}
	for key, value := range frqCountMap {

		array := bucket[value]

		array = append(array, key)

		bucket[value] = array
	}

outer:
	for i := len(bucket) - 1; i >= 0; i-- {

		array := bucket[i]

		for j := len(array) - 1; j >= 0; j-- {

			value := array[j]

			if len(output) < k {
				output = append(output, value)
			} else {
				break outer
			}

		}
	}
	return output
}
