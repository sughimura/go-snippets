package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var sum int
	for _, si := range s {
		if si == '1' {
			sum++
		}
	}
	fmt.Println(sum)
}
