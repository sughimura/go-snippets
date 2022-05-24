package main

import (
	"fmt"
	"math"
)

func isPrime(x int) bool {
	if x == 1 {
		return false
	}
	if x == 2 {
		return true
	}
	if x < 2 || x%2 == 0 {
		return false
	}
	for i := 3; float64(i) <= math.Pow(float64(x), 0.5); i = i + 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)
	var count int
	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		if isPrime(number) {
			count++
		}
	}
	fmt.Println(count)
}
