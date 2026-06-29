package arraysandhashes

func ProductExceptSelf(nums []int) []int {

	rightProducts := []int{}
	out := []int{}

	for i := len(nums) - 1; i >= 0; i-- {

		num := nums[i]

		if i == len(nums)-1 {
			rightProducts = append(rightProducts, num)
		} else {
			prevIndex := (len(nums) - 1) - (i + 1)
			rightProducts = append(rightProducts, rightProducts[prevIndex]*num)
		}

	}

	left := nums[0]
	for i := range nums {

		num := nums[i]
		rightProductIndex := (len(nums) - 1) - (i + 1)

		if rightProductIndex >= 0 {

			if i == 0 {
				out = append(out, rightProducts[rightProductIndex])
			} else {
				rightProduct := rightProducts[rightProductIndex]

				currentProduct := left * rightProduct
				out = append(out, currentProduct)
				left *= num
			}
		} else {
			out = append(out, left)
		}
	}

	return out
}
