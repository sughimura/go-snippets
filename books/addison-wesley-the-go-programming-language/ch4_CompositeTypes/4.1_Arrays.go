package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	fmt.Println(q)
	var r [3]int = [3]int{1, 2}
	fmt.Println(r)

	q2 := [...]int{1, 2, 3}
	// 型を表示
	fmt.Printf("%T\n", q2)

	q3 := [3]int{1, 2, 3}
	//q3 = [4]int{1, 2, 3, 4} // エラー
	fmt.Println(q3)

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])

	r2 := [...]int{9: -1}
	fmt.Println(r2) // [0 0 0 0 0 0 0 0 0 -1]

	a2 := [2]int{1, 2}
	b2 := [...]int{1, 2}
	c2 := [2]int{1, 3}
	fmt.Println(a2 == b2) // true
	fmt.Println(a2 == c2) // false
	fmt.Println(b2 == c2) // false
	d := [3]int{1, 2}
	//fmt.Println(a2 == d) // invalid operation: a2 == d (mismatched types [2]int and [3]int)
	fmt.Println(d)
}
