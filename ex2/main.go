package main

import "fmt"

func fib(till int) {
	a1 := 1
	a2 := 2
	s := a1 + a2
	for {
		a1, a2 = a2, a1+a2
		if a2%2 == 0 {
			s += a2
		}
		if s >= till {
			break
		}
	}
	fmt.Println(s)

}

func main() {
	fib(4e+6)
}
