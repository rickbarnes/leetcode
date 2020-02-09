package maxsubarray

import (
	"fmt"
	"testing"
)

func TestMaxSubarray(t *testing.T) {
	res := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	ans := 6
	if res != ans {
		t.Errorf("Received: %d, Expected: %d", res, ans)
	}
	fmt.Printf("Received %d, Expected: %d", res, ans)

}
