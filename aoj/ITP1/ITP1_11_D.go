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

func isSameDice(d1 Dice, d2 Dice) bool {
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
		return false
	}
	if d1.bottom() != d2.bottom() {
		return false
	}
	for i := 0; i < 4; i++ {
		if d1.front() == d2.front() {
			break
		}
		d2.turnToRight()
	}
	if d1.front() == d2.front() && d1.back() == d2.back() && d1.right() == d2.right() && d1.left() == d2.left() {
		return true
	} else {
		return false
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	dices := make([]Dice, n)
	for i := 0; i < n; i++ {
		v := make([]int, 6)
		for j := 0; j < 6; j++ {
			fmt.Scan(&v[j])
		}
		dices[i] = Dice{v[0], v[1], v[2], v[3], v[4], v[5]}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				if isSameDice(dices[i], dices[j]) {
					fmt.Println("No")
					return
				}
			}
		}
	}
	fmt.Println("Yes")
}
