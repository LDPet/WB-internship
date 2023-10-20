package main

import "math"

type Point struct {
	// имена с маленькой буквы видны только в рамках пакета
	x, y float64
}

// NewPoint конструктор
func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p *Point) dist(p2 Point) float64 {
	return math.Sqrt(math.Pow(p.x-p2.x, 2) + math.Pow(p.y-p2.y, 2))
}

func main() {
	p1 := NewPoint(1, 1)
	p2 := NewPoint(4, 6)

	println(p1.dist(p2))
}
