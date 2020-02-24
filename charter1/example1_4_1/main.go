package main

import "fmt"

func main() {
	fmt.Println("------- Integer ------")
	a := Integer{4}
	b := Integer{8}
	fmt.Printf("compare结果: %v\n", a.compare(b))

	fmt.Println("------- Point ------")
	//var point Point
	point := new(Point)
	point.setXY(2.1, 3.5)
	x, y := point.getXY()
	fmt.Printf("x: %f, y: %f\n", x, y)
}

type Integer struct {
	value int32
}

func (a Integer) compare(b Integer) bool {
	return a.value < b.value
}

type Point struct {
	x float32
	y float32
}

func (point *Point) setXY(x, y float32) {
	point.x = x
	point.y = y
}

func (point *Point) getXY() (x, y float32)  {
	return point.x, point.y
}