package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	root := int(math.Sqrt(float64(n)))
	for i := 2; i <= root; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func divider(n int) int {
	root := int(math.Sqrt(float64(n)))
	maxDivider := 1
	for i := 2; i <= root; i++ {
		if isPrime(i) && n%i == 0 {
			maxDivider = i
			n = n / maxDivider
			fmt.Printf("Found divider %d, new value = %d\n", maxDivider, n)
			if n == 1 {
				break
			}
		}
	}
	return maxDivider
}

func main() {
	fmt.Println(divider(12))

}
