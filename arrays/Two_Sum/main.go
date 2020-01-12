package twosum

import (
	"errors"
)

func twoSum(nums []int, target int) ([]int, error) {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		comp := target - nums[i]

		val, ok := m[comp]
		if ok {
			return []int{val, i}, nil
		}
		m[nums[i]] = i
	}
	return nil, errors.New("Cannot be solved")
}
