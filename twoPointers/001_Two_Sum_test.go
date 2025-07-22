package twoPointers

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		result := TwoSum(tt.nums, tt.target)

		// Since order might matter, check with reflect.DeepEqual
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("twoSum(%v, %d) = %v; expected %v", tt.nums, tt.target, result, tt.expected)
		}
	}
}
