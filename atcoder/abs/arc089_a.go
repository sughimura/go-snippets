package main

import (
	"fmt"
	"math"
)

type Point struct {
	t, x, y int
}

func main() {
	var n int
	fmt.Scan(&n)
	points := make([]Point, n+1)
	points[0] = Point{t: 0, x: 0, y: 0}
	for i := 1; i <= n; i++ {
		var t, x, y int
		fmt.Scan(&t, &x, &y)
		points[i] = Point{t: t, x: x, y: y}
	}
	for i := 0; i <= n-1; i++ {
		step := points[i+1].t - points[i].t
		xDistance := math.Abs(float64(points[i+1].x - points[i].x))
		yDistance := math.Abs(float64(points[i+1].y - points[i].y))
		rest := step - int(xDistance+yDistance)
		if rest < 0 || rest%2 != 0 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
