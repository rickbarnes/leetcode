package twosum

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		for x := 1; x < len(nums)-1; x++ {
			if nums[x]+nums[i] == target {
				return []int{i, x}
			}
		}
	}
	return []int{0, 0}
}
