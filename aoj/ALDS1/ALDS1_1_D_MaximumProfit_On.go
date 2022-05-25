package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	r := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		r[i], _ = strconv.Atoi(scanner.Text())
	}
	maxV := r[1] - r[0]
	minV := r[0]
	for j := 1; j < n; j++ {
		maxV = int(math.Max(float64(maxV), float64(r[j]-minV)))
		minV = int(math.Min(float64(minV), float64(r[j])))
	}
	fmt.Println(maxV)
}
