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
	cnt := 0
	insertionSort := func(n int, g int) {
		for i := g; i < n; i++ {
			v := A[i]
			j := i - g
			for j >= 0 && A[j] > v {
				A[j+g] = A[j]
				j = j - g
				cnt++
			}
			A[j+1] = v
		}
	}

}
