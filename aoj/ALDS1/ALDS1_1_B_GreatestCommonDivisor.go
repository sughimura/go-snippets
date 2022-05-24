package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if y > x {
		x, y = y, x
	} else if x == y {
		fmt.Println(x)
		return
	}
	xpy := x % y
	var yDivisors []int
	var xpyDivisors []int
	for i := 1; i <= y; i++ {
		if y%i == 0 {
			yDivisors = append(yDivisors, i)
		}
	}
	for i := 1; i <= xpy; i++ {
		if xpy%i == 0 {
			xpyDivisors = append(xpyDivisors, i)
		}
	}
	for i := len(yDivisors) - 1; i >= 0; i-- {
		for j := 0; j < len(xpyDivisors); j++ {
			if xpyDivisors[j] == yDivisors[i] {
				fmt.Println(xpyDivisors[j])
				return
			}
		}
	}
}
