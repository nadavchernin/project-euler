package main

import (
	"fmt"
	"slices"
)

func isPalindrom(n int) bool {
	var dgts []int
	for {
		rem := n % 10
		n = n / 10
		dgts = append(dgts, rem)
		if n == 0 {
			break
		}
	}
	reversedDgts := slices.Clone(dgts)

	slices.Reverse(reversedDgts)
	// fmt.Println(dgts)
	// fmt.Println(reversedDgts)
	return slices.Equal(dgts, reversedDgts)
}

func main() {
	maxPalindrom := 1
	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			if i*j > maxPalindrom && isPalindrom(i*j) {
				maxPalindrom = i * j
			}
		}
	}
	fmt.Println(maxPalindrom)
}
