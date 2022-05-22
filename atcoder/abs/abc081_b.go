package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var sum int
	isAllEven := true
	for {
		for i := 0; i < n; i++ {
			if a[i]%2 != 0 {
				isAllEven = false
			}
			a[i] /= 2
		}
		if !isAllEven {
			break
		} else {
			sum++
		}
	}
	fmt.Println(sum)
}
