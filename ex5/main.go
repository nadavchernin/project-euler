package main

import (
	"fmt"
	"math"
)

func pow(k int, v int) int {
	power := 1
	for i := 0; i < v; i++ {
		power *= k
	}
	return power
}

func dividers(n int) map[int]int {
	dgts := make(map[int]int)
	root := int(math.Sqrt(float64(n)))
	for i := 2; i <= root; i++ {
		for {
			if n%i == 0 {
				dgts[i] += 1
				n /= i
			} else {
				break
			}
		}
	}
	if n > 1 {
		dgts[n] += 1
	}
	return dgts
}

func main() {
	dgts := make(map[int]int)
	for i := 2; i <= 20; i++ {
		for k, v := range dividers(i) {
			dgts[k] = max(dgts[k], v)
		}
	}
	s := 1
	for k, v := range dgts {
		s *= pow(k, v)
	}
	fmt.Println(s)
}
