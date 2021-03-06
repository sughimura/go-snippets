package main

import "fmt"

// intのスライスを逆順に並び替える
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 配列と異なりスライスは比較可能では無い
// バイトスライス([]byte)の比較にはbytes.Equalを使えるが
// それ以外の型のスライスの比較に == は使えないので、自分で比較が必要
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
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

	aa := [...]string{"a", "b", "c"}
	bb := [...]string{"a", "a", "a"}
	fmt.Println(equal(aa[:], bb[:]))

	//	4.2.1 append関数
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Println(runes)
	fmt.Printf("%q\n", runes)
}
