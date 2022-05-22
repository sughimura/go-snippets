// 16:39 - 17:05

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

	sort.Ints(a)

	var before int
	for i := 0; i < n; i++ {
		if a[i] != (i + 1) {
			before = i + 1
			break
		}
	}

	counts := make(map[int]int)
	for i := 0; i < n; i++ {
		counts[a[i]]++
	}

	var after int
	for i := 1; i <= n; i++ {
		if counts[i] == 2 {
			after = i
		}
	}

	if before == 0 {
		fmt.Println("Correct")
	} else {
		fmt.Printf("%d %d\n", after, before)
	}
}
