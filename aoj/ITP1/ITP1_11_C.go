package main

import "fmt"

type Dice struct {
	way1, way2, way3, way4, way5, way6 int
}

func (d *Dice) rollToLeft() {
	d.way1, d.way3, d.way6, d.way4 = d.way3, d.way6, d.way4, d.way1
}

func (d *Dice) rollToRight() {
	d.way1, d.way4, d.way6, d.way3 = d.way4, d.way6, d.way3, d.way1
}

func (d *Dice) rollToFront() {
	d.way1, d.way5, d.way6, d.way2 = d.way5, d.way6, d.way2, d.way1
}

func (d *Dice) rollToBack() {
	d.way1, d.way2, d.way6, d.way5 = d.way2, d.way6, d.way5, d.way1
}

func (d *Dice) turnToRight() {
	d.way2, d.way3, d.way5, d.way4 = d.way4, d.way2, d.way3, d.way5
}

func (d *Dice) top() int {
	return d.way1
}

func (d *Dice) bottom() int {
	return d.way6
}

func (d *Dice) left() int {
	return d.way4
}

func (d *Dice) right() int {
	return d.way3
}

func (d *Dice) front() int {
	return d.way2
}

func (d *Dice) back() int {
	return d.way5
}

func main() {
	v1 := make([]int, 6)
	v2 := make([]int, 6)
	for i := 0; i < 6; i++ {
		fmt.Scan(&v1[i])
	}
	for i := 0; i < 6; i++ {
		fmt.Scan(&v2[i])
	}
	d1 := Dice{v1[0], v1[1], v1[2], v1[3], v1[4], v1[5]}
	d2 := Dice{v2[0], v2[1], v2[2], v2[3], v2[4], v2[5]}
	if d1.top() == d2.front() {
		d2.rollToBack()
	}
	if d1.top() == d2.back() {
		d2.rollToFront()
	}
	for i := 0; i < 4; i++ {
		if d1.top() == d2.top() {
			break
		}
		d2.rollToRight()
	}
	if d1.top() != d2.top() {
		println("No")
		return
	}
	if d1.bottom() != d2.bottom() {
		println("No")
		return
	}
	for i := 0; i < 4; i++ {
		if d1.front() == d2.front() {
			break
		}
		d2.turnToRight()
	}
	if d1.front() == d2.front() && d1.back() == d2.back() && d1.right() == d2.right() && d1.left() == d2.left() {
		println("Yes")
	} else {
		println("No")
	}
}
