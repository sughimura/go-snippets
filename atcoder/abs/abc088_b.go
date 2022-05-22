package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	var alice, bob int
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			alice += a[i]
		} else {
			bob += a[i]
		}
	}
	fmt.Println(alice - bob)
}
