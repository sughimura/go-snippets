package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX = 1000000

type Stack struct {
	vec [MAX]int
	top int
}

func (st *Stack) isEmpty() bool {
	return st.top == 0
}

func (st *Stack) isFull() bool {
	return st.top >= MAX-1
}

func (st *Stack) push(x int) {
	if st.isFull() {
		return
	}
	st.top++
	st.vec[st.top] = x
}

func (st *Stack) pop() int {
	if st.isEmpty() {
		return -1
	}
	st.top--
	return st.vec[st.top+1]
}

func main() {
	st := Stack{}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputArray := strings.Split(scanner.Text(), " ")
	for _, s := range inputArray {
		if s == "+" {
			y := st.pop()
			x := st.pop()
			st.push(x + y)
		} else if s == "-" {
			y := st.pop()
			x := st.pop()
			st.push(x - y)
		} else if s == "*" {
			y := st.pop()
			x := st.pop()
			st.push(x * y)
		} else {
			x, _ := strconv.Atoi(s)
			st.push(x)
		}
	}
	fmt.Println(st.pop())
}
