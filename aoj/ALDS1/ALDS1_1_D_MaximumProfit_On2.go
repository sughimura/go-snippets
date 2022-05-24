package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&r[i])
	}
	maxV := r[1] - r[0]
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if maxV < r[j]-r[i] {
				maxV = r[j] - r[i]
			}
		}
	}
	fmt.Println(maxV)
}
