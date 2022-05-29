package main

import "fmt"

// 空文字列では無い文字列を保持するスライスを返す
// 基底配列は呼び出し中に修正される
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	// 4.2.2 スライス内での技法
	data := []string{"one", "", "three"}
	fmt.Println(nonempty(data))
	fmt.Println(data) // [one three three]

	data2 := []string{"one", "", "three"}
	fmt.Println(nonempty2(data2))
	fmt.Println(data2) // [one three three]
}
