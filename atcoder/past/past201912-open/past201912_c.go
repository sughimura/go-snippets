// 16:29 - 16:33

package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 6
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println(numbers[2])
}
