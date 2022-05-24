package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&r[i])
	}
	maxV := r[1] - r[0]
	minV := r[0]
	for j := 1; j < n; j++ {
		maxV = int(math.Max(float64(maxV), float64(r[j]-minV)))
		minV = int(math.Min(float64(minV), float64(r[j])))
	}
	fmt.Println(maxV)
}
