package main

import (
	"fmt"
	"math"
)

func main() {
	num := int64(600851475143)

	var maxPrime int64
	for i := int64(2); i <= int64(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			isPrime := true

			for j := int64(2); j <= int64(math.Sqrt(float64(i))); j++ {
				if i%j == 0 {
					isPrime = false
					break
				}
			}

			if isPrime {
				maxPrime = i
				fmt.Println(maxPrime)
			}
		}
	}

	fmt.Println(maxPrime)
}
