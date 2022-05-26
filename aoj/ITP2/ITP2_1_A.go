package main

import "fmt"

func main() {
	var q int
	fmt.Scan(&q)
	var A []int
	for i := 0; i < q; i++ {
		var query, arg int
		fmt.Scan(&query)
		if query == 0 {
			fmt.Scan(&arg)
			A = append(A, arg)
		} else if query == 1 {
			fmt.Scan(&arg)
			fmt.Println(A[arg])
		} else if query == 2 {
			A = A[:len(A)-1]
		}
	}
}
