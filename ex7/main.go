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
func main() {

	count := 0
	for i := 2; true; i++ {
		if isPrime(i) {
			count += 1
		}
		if count == 10001 {
			fmt.Println(i)
			break
		}

	}

}
