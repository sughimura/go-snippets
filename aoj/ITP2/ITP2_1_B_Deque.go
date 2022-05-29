package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func push(queue []int, pushPosition int, value int, head int, tail int) ([]int, int, int) {
	if pushPosition == 0 {
		head--
		queue[head] = value
	} else {
		queue[tail] = value
		tail++
	}
	return queue, head, tail
}

func pop(popPosition int, head int, tail int) (int, int) {
	if popPosition == 0 {
		head++
	} else {
		tail--
	}
	return head, tail
}

func main() {
	var q int
	fmt.Scan(&q)
	A := make([]int, q*2)
	head := q
	tail := head
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < q; i++ {
		scanner.Scan()
		query := strings.Split(scanner.Text(), " ")
		if query[0] == "0" {
			d, _ := strconv.Atoi(query[1])
			x, _ := strconv.Atoi(query[2])
			A, head, tail = push(A, d, x, head, tail)
		} else if query[0] == "1" {
			p, _ := strconv.Atoi(query[1])
			fmt.Println(A[head+p])
		} else if query[0] == "2" {
			d, _ := strconv.Atoi(query[1])
			head, tail = pop(d, head, tail)
		}
	}
}
