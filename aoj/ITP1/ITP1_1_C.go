package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scanln(&a, &b)
	fmt.Println(a*b, 2*a+2*b)
}
