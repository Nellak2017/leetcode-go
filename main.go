package main

import (
	"fmt"
	"leetcode/twoPointers"
	"os"
)

var showcases = map[string]func(){
	"1": twoPointers.TwoSumShowcase,
}

func main() {
	problemID := "1"
	if len(os.Args) > 1 {
		problemID = os.Args[1]
	}
	showcaseFunc, ok := showcases[problemID]
	if ok {
		showcaseFunc()
	} else if fallback, hasFallback := showcases["1"]; hasFallback {
		fmt.Printf("Invalid problem number '%s'. Falling back to problem 1.\n\n", problemID)
		fallback()
	} else {
		fmt.Println("No valid problem found and no fallback available. You have no problems solved.")
	}
}
