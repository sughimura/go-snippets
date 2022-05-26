package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var q int
	fmt.Scan(&q)
	var A []int
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < q; i++ {
		scanner.Scan()
		query := strings.Split(scanner.Text(), " ")
		if query[0] == "0" {
			x, _ := strconv.Atoi(query[1])
			A = append(A, x)
		} else if query[0] == "1" {
			p, _ := strconv.Atoi(query[1])
			fmt.Println(A[p])
		} else if query[0] == "2" {
			A = A[:len(A)-1]
		}
	}
}
