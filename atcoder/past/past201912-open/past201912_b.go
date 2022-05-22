//16:12 - 16:23

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i < n; i++ {
		diff := a[i] - a[i-1]
		if diff == 0 {
			fmt.Println("stay")
		} else if diff > 0 {
			fmt.Printf("up %d\n", diff)
		} else {
			fmt.Printf("down %d\n", diff*(-1))
		}
	}
}
