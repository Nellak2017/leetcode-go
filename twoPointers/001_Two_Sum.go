package twoPointers

import (
	"fmt"
	"leetcode/internal"
	"slices"
	"sort"
)

func TwoSum(nums []int, target int) []int {
	// 1. Sort the nums with a copy O(n) + O(n log n)
	sortedCopy := slices.Clone(nums)
	sort.Ints(sortedCopy)

	// 2. Apply 2 pointer technique O(n)
	left, right := 0, len(sortedCopy)-1
	for left < right {
		sum := sortedCopy[left] + sortedCopy[right]
		if sum == target {
			break
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	// 3. Find what goes to what index and return those O(n)
	leftIndex := internal.IndexOf(nums, sortedCopy[left], 0)
	rightIndex := internal.IndexOf(nums, sortedCopy[right], leftIndex+1)
	if leftIndex >= rightIndex {
		return []int{internal.IndexOf(nums, sortedCopy[right], 0), internal.IndexOf(nums, sortedCopy[left], leftIndex+1)}
	} // overall O(n log n)
	return []int{leftIndex, rightIndex}
}

func TwoSumShowcase() {
	cases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
		{[]int{-1, -2, -3, -4, -5}, -8},
	}
	for _, input := range cases {
		result := TwoSum(input.nums, input.target)
		fmt.Printf("List is %v, Target is %v, Indices of items in sum %v\n", input.nums, input.target, result)
	}
}
