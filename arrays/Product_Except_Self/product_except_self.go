package productexceptself

func productExceptSelf(nums []int) []int {
	var res []int

	if len(nums) > 0 {
		prod := 1

		for i := 0; i < len(nums); i++ {
			res = append(res, prod)
			prod *= nums[i]
		}

		prod = 1
		for i := len(nums) - 1; i >= 0; i-- {
			res[i] *= prod
			prod *= nums[i]
		}
	}

	return res
}
