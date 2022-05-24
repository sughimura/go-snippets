package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	var n int
	if x <= y {
		n = x
	} else {
		n = y
	}
	for i := n; i >= 1; i-- {
		if x%i == 0 && y%i == 0 {
			fmt.Println(i)
			return
		}
	}
}
