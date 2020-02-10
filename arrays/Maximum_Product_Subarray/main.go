package maxproduct

func maxProduct(nums []int) int {
	res := nums[0]
	ans := res

	for i, num := range nums {
		if i != 0 {
			res *= num
		}

		if num >= ans {
			ans = num
		}

		if res > ans {
			ans = res
		}
	}

	return ans
}
