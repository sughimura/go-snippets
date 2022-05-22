package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	number, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(number * 2)
	}
}
