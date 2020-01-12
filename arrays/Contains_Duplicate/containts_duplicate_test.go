package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	res := containsDuplicate([]int{1, 2, 3, 1})
	ans := true

	if res != ans {
		t.Errorf("Received: %v, Expected: %v", res, ans)
	}
}

func TestContainsDuplicateTwo(t *testing.T) {
	res := containsDuplicate([]int{1, 2, 3, 4})
	ans := false

	if res != ans {
		t.Errorf("Received: %v, Expected: %v", res, ans)
	}
}

func TestContainsDuplicateThree(t *testing.T) {
	res := containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	ans := true

	if res != ans {
		t.Errorf("Received: %v, Expected: %v", res, ans)
	}
}
