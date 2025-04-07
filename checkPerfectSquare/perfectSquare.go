package main

import "fmt"

func main() {
	isPerfectSquare(16)
	isPerfectSquare(24)

}

func isPerfectSquare(num int) bool {
	initialIndex := num

	if num > 1 {
		initialIndex = num / 2
	}

	for i := initialIndex; i > 0; i-- {
		fmt.Printf("%v * %v = %v | target = %v\n", i, i, i*i, num)
		if i*i == num {
			return true
		}
	}

	return false
}
