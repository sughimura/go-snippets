package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	dm := make(map[int]int)
	for i := 0; i < n; i++ {
		var di int
		fmt.Scan(&di)
		dm[di] = di
	}
	fmt.Println(len(dm))
}
