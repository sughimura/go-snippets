package main

//
//import "fmt"
//
//type Dice struct {
//	way1, way2, way3, way4, way5, way6 int
//}
//
//func (d *Dice) rollToW() {
//	d.way1, d.way3, d.way6, d.way4 = d.way3, d.way6, d.way4, d.way1
//}
//
//func (d *Dice) rollToE() {
//	d.way1, d.way4, d.way6, d.way3 = d.way4, d.way6, d.way3, d.way1
//}
//
//func (d *Dice) rollToS() {
//	d.way1, d.way5, d.way6, d.way2 = d.way5, d.way6, d.way2, d.way1
//}
//
//func (d *Dice) rollToN() {
//	d.way1, d.way2, d.way6, d.way5 = d.way2, d.way6, d.way5, d.way1
//}
//
//func (d *Dice) top() int {
//	return d.way1
//}
//
//func (d *Dice) bottom() int {
//	return d.way6
//}
//
//func (d *Dice) left() int {
//	return d.way4
//}
//
//func (d *Dice) right() int {
//	return d.way3
//}
//
//func (d *Dice) front() int {
//	return d.way2
//}
//
//func (d *Dice) back() int {
//	return d.way5
//}
//
//func main() {
//	v := make([]int, 6)
//	for i := 0; i < 6; i++ {
//		fmt.Scan(&v[i])
//	}
//	d := Dice{v[0], v[1], v[2], v[3], v[4], v[5]}
//	var n int
//	fmt.Scan(&n)
//	for i := 0; i < n; i++ {
//		var t, f int
//		fmt.Scan(&t, &f)
//		if d.left() == f {
//			d.rollToW()
//		} else if d.right() == f {
//			d.rollToE()
//		}
//		for d.front() != f {
//			d.rollToS()
//		}
//		for d.top() != t {
//			d.rollToE()
//		}
//		fmt.Println(d.right())
//	}
//}
