package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	a := twoSum([]int{2, 7, 11, 15}, 9)
	b := []int{0, 1}
	if !reflect.DeepEqual(a, b) {
		t.Error("Failure")
	}
}
