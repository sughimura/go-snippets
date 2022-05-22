package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	for len(s) > 0 {
		if strings.HasPrefix(s, "dreameraser") {
			s = strings.TrimPrefix(s, "dreameraser")
		} else if strings.HasPrefix(s, "dreamerase") {
			s = strings.TrimPrefix(s, "dreamerase")
		} else if strings.HasPrefix(s, "dreamer") {
			s = strings.TrimPrefix(s, "dreamer")
		} else if strings.HasPrefix(s, "dream") {
			s = strings.TrimPrefix(s, "dream")
		} else if strings.HasPrefix(s, "eraser") {
			s = strings.TrimPrefix(s, "eraser")
		} else if strings.HasPrefix(s, "erase") {
			s = strings.TrimPrefix(s, "erase")
		} else {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
