package main

import (
	"fmt"
	"math"
)

func isPrime(num int64) bool {
	for i := int64(2); i <= int64(math.Floor(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	mult := int64(1)
	for i := int64(2); i <= 20; i++ {
		if !isPrime(i) {
			continue
		}
		prod := i
		for prod <= 20 {
			prod *= i
		}
		prod /= i
		mult *= prod
	}
	fmt.Println(mult)
}
