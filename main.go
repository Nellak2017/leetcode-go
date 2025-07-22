package main

import (
	"fmt"
	"leetcode/twoPointers"
)

func main() {
	nums := []int{-1, -2, -3, -4, -5}
	target := -8
	result := twoPointers.TwoSum(nums, target)
	fmt.Printf("Result is %v", result)
}
