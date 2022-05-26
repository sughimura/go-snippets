package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	var count int
	flag := true
	i := 0
	for flag {
		flag = false
		for j := n - 1; j >= i+1; j-- {
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
				count++
				flag = true
			}
		}
		i++
	}
	for i := 0; i < n; i++ {
		fmt.Print(A[i])
		if i == n-1 {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println(count)
}
