package main

import "fmt"

func main() {
	var x1 [3]int
	fmt.Println(x1)

	var x2 = [3]int{1, 2}
	fmt.Println(x2)

	var x3 = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(x3)

	var x4 = [...]int{1, 2, 3}
	fmt.Println(x4)
	var y4 = [3]int{1, 2, 3}
	fmt.Println(x4)
	fmt.Println(x4 == y4)

	var x5 [2][3]int
	fmt.Println(x5)

	var x6 [3]int
	x6[0] = 100
	fmt.Println(x6)
}
