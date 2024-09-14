package main

import "fmt"

func main() {
	var x1 = []int{1, 2, 3}
	fmt.Println(x1)

	var x2 = []int{1, 5: 4, 10: 100, 15}
	fmt.Println(x2)

	var x3 [][]int
	fmt.Println(x3)

	var x4 = []int{1, 2, 3}
	fmt.Println(len(x4))

	var x5 []int
	fmt.Println(x5)
	x5 = append(x5, 1)
	fmt.Println(x5)
	x5 = append(x5, []int{1, 2, 3}...)
	fmt.Println(x5)

	var x6 []int
	fmt.Println(x6, len(x6), cap(x6))
	x6 = append(x6, []int{1, 2, 3}...)
	fmt.Println(x6, len(x6), cap(x6))
	x6 = append(x6, []int{1, 2, 3}...)
	fmt.Println(x6, len(x6), cap(x6))
	x6 = append(x6, 1)
	fmt.Println(x6, len(x6), cap(x6))
	x6 = append(x6, []int{1, 2, 3, 4, 5, 6, 7}...)
	fmt.Println(x6, len(x6), cap(x6))

	x7 := make([]int, 3)
	fmt.Println(x7, len(x7), cap(x7))

	x8 := make([]int, 5, 10)
	fmt.Println(x8, len(x8), cap(x8))

	x9 := make([]int, 0, 5)
	fmt.Println(x9, len(x9), cap(x9))
	x9 = append(x9, 1)
	fmt.Println(x9, len(x9), cap(x9))
}
