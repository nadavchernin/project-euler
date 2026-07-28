package main

import "fmt"

func solution(n int) int {
	s := 0
	ss := 0
	for i := 1; i <= n; i++ {
		s += i
		ss += i * i
	}
	return s*s - ss
}

func main() {
	fmt.Println(solution(100))
}
