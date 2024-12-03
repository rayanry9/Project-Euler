package main

import (
	"fmt"
)

func main() {
	val := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			val += i
		}
	}
	fmt.Println(val)
}
