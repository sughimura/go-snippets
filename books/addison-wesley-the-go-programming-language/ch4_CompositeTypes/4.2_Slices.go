package main

import "fmt"

// intのスライスを逆順に並び替える
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	fmt.Println(months)
	Q2 := months[4:7]
	fmt.Println(Q2)
	summer := months[6:9]
	fmt.Println(summer)

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)
	fmt.Printf("%T\n", a)    // 配列
	fmt.Printf("%T\n", a[:]) // スライス
}
