package maxproduct

import (
	"log"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	nums := []int{2, 3, -2, 4}
	ans := 6
	res := maxProduct(nums)

	if ans != res {
		log.Fatalf("Test 1 Fail: Received: %d, Expected: %d\n", res, ans)
	}
}

func TestMaxProduct2(t *testing.T) {
	nums := []int{-2, 0, -1}
	ans := 0
	res := maxProduct(nums)

	if ans != res {
		log.Fatalf("Test 2 Fail: Received: %d, Expected: %d\n", res, ans)
	}
}

func TestMaxProduct3(t *testing.T) {
	nums := []int{2, -5, -2, -4, 3}
	ans := 24
	res := maxProduct(nums)

	if ans != res {
		log.Fatalf("Test 3 Fail: Received: %d, Expected: %d\n", res, ans)
	}
}
