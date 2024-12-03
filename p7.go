package main

import (
	"fmt"
	"math"
)

func isPrimee(num int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	i := 0
	for j := 2; ; j++ {
		if isPrimee(j) {
			i++
		}
		if i == 10001 {
			fmt.Println(j)
			break
		}
	}
}
