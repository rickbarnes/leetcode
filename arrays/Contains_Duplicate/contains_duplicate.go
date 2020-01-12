package containsduplicate

func containsDuplicate(nums []int) bool {
	if len(nums) > 0 {
		seen := make(map[int]bool)

		for i := 0; i < len(nums); i++ {
			_, ok := seen[nums[i]]
			if ok {
				return true
			}
			seen[nums[i]] = true
		}
	}

	return false
}
