package main

import "fmt"

func printArray(A []int) {
	n := len(A)
	for i := 0; i < n; i++ {
		fmt.Printf("%d", A[i])
		if i == n-1 {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	for i := 1; i < n; i++ {
		printArray(A)
		v := A[i]
		j := i - 1
		for j >= 0 && A[j] > v {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = v
	}
	printArray(A)
}
