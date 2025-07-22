package twoPointers

import (
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
	if left >= right {
		return nil
	}
	// 3. Find what goes to what index and return those O(n)
	leftIndex := internal.IndexOf(nums, sortedCopy[left], 0)
	rightIndex := internal.IndexOf(nums, sortedCopy[right], leftIndex+1)
	if leftIndex < rightIndex {
		return []int{leftIndex, rightIndex}
	} else {
		leftIndex = internal.IndexOf(nums, sortedCopy[right], 0)
		rightIndex = internal.IndexOf(nums, sortedCopy[left], leftIndex+1)
		return []int{leftIndex, rightIndex}
	} // overall O(n log n)
}
