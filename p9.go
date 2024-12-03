package main

import "fmt"

func eqPythagorean(a int, b int) bool {
	return a*b+500_000 == 1000*(a+b)
}

func main() {
	for a := 1; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			if eqPythagorean(a, b) {
				fmt.Println(a * b * (1000 - a - b))
			}
		}
	}
}
