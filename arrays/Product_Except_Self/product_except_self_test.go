package productexceptself

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	res := productExceptSelf([]int{1, 2, 3, 4})
	ans := []int{24, 12, 8, 6}

	if !reflect.DeepEqual(res, ans) {
		t.Errorf("Expected: %d, Received: %d", ans, res)
	}
}
