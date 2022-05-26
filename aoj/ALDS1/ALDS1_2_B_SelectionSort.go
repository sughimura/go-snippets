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
	for i := 0; i < n; i++ {
		minj := i
		for j := i; j < n; j++ {
			if A[j] < A[minj] {
				minj = j
			}
		}
		if A[i] != A[minj] {
			A[i], A[minj] = A[minj], A[i]
			count++
		}
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
