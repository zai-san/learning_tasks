package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
	perimetr() float64
}

type MultiShape struct {
	shapes []Shape
}

//func circleArea(x, y, r float64) float64 {
//  return math.Pi * r*r
//}

//func circleArea(c *Circle) float64 {
//	return math.Pi * c.r * c.r
//}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimetr() float64 {
	return math.Pi * c.r * 2
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

//func rectangleArea(x1, y1, x2, y2 float64) float64 {
//    l := distance(x1, y1, x1, y2)
//    w := distance(x1, y1, x2, y1)
//    return l * w
//}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimetr() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2*l + 2*w
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	//	var rx1, ry1 float64 = 0, 0
	//	var rx2, ry2 float64 = 10, 10
	r := Rectangle{0, 0, 10, 10}
	//	var cx, cy, cr float64 = 0, 0, 5
	c := Circle{0, 0, 5}
	//fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(r.area())
	fmt.Println(r.perimetr())
	//	fmt.Println(circleArea(cx, cy, cr))
	//fmt.Println(circleArea(&c))
	fmt.Println(c.area())
	fmt.Println(c.perimetr())
	fmt.Println(totalArea(&c, &r))
}
