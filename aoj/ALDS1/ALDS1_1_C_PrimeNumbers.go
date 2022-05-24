package main

import "fmt"

func isPrimeSimple(x int) bool {
	if x <= 1 {
		return false
	}
	for i := 2; i < x-1; i++ {
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
		if isPrimeSimple(number) {
			count++
		}
	}
	fmt.Println(count)
}
