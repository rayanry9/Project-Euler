package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	limit := 2_000_000
	sum := 0

	for i := 2; i < limit; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
