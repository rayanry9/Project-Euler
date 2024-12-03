package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(num int) bool {
	str := strconv.Itoa(num)
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	highest := 999_999

	for {
		if isPalindrome(highest) {
			for i := 100; i < 1000; i++ {
				if highest%i == 0 {
					if highest/i < 1000 && highest/i >= 100 {
						fmt.Println(highest)
						return
					}
				}
			}
		}
		highest--
	}
}
