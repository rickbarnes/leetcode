package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	a, err := twoSum([]int{2, 7, 11, 15}, 9)
	if err != nil {
		t.Errorf("Cannot solve")
	}
	b := []int{0, 1}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected: %d, Received: %d", b, a)
	}
}
