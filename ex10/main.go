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
	s := 0
	for i := 2; i < 2000000; i++ {
		if isPrime(i) {
			s += i
		}
	}
	fmt.Println(s)

}
