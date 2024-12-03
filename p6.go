package main

import "fmt"

func main() {
	sumOfSquares := 100 * 101 * 201 / 6
	squareOfSums := 50 * 101 * 50 * 101
	fmt.Println(squareOfSums - sumOfSquares)
}
