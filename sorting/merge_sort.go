package sorting

func MergeSort(nums []int) []int {

	left := 0
	midd := len(nums) / 2
	right := midd

	if left >= right {
		return nums
	}

	MergeSort(nums[:right])
	MergeSort(nums[right:])
	return mergeArray(nums, nums[:right], nums[right:])

}

func mergeArray(nums []int, leftArr []int, rightArr []int) []int {

	temp := make([]int, len(leftArr)+len(rightArr))

	i := 0
	j := 0
	k := 0
	for i < len(leftArr) && j < len(rightArr) {

		if leftArr[i] < rightArr[j] {
			temp[k] = leftArr[i]
			k++
			i++
		} else {
			temp[k] = rightArr[j]
			k++
			j++
		}
	}

	for i < len(leftArr) {
		temp[k] = leftArr[i]
		i++
		k++
	}
	for j < len(rightArr) {
		temp[k] = rightArr[j]
		j++
		k++
	}

	copy(nums, temp)

	return nums
}
