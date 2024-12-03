package main

import "fmt"

func main() {
	var val int64
	val = 2
	var f1 int64 = 1
	var f2 int64 = 2
	for {
		if f2 >= 4_000_000 {
			break
		}
		f1 = f1 + f2
		if f1%2 == 0 {
			val += f1
		}

		temp := f1
		f1 = f2
		f2 = temp
	}
	fmt.Println(val)
}
