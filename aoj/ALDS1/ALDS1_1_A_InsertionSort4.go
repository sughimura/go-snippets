package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	A := make([]int, n)
	printA := func() {
		for i := 0; i < n; i++ {
			fmt.Print(A[i])
			if i == n-1 {
				fmt.Print("\n")
			} else {
				fmt.Print(" ")
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	for i := 1; i < n; i++ {
		printA()
		v := A[i]
		j := i - 1
		for j >= 0 && A[j] >= v {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = v
	}
	printA()
}
