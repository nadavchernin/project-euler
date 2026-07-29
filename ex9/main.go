package main

import "fmt"

func main() {
	found := false
	for a := 1; a < 998; a++ {
		for b := 1; b < 998; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				found = true
				break
			}
		}
		if found {
			break
		}
	}
}
