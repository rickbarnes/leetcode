package maxsubarray

import "fmt"

func maxSubArray(nums []int) int {
	ans := nums[0]
	res := ans

	for i, num := range nums {
		if i != 0 {
			ans += num
		}

		if num > ans {
			ans = num
		}

		if ans > res {
			res = ans
		}

		fmt.Printf("Num: %v, Ans: %v, Res: %v \n", num, ans, res)
	}

	return res
}
