package shapes

import "math"

type IShape interface {
	GetArea() float64
}

type shape struct {
	color string
}

func (s shape) GetColor() string {
	return s.color
}

type Square struct {
	shape
	width int
}

func NewSquare(color string, width int) *Square {
	return &Square{
		shape: shape{color},
		width: width,
	}
}

func (s *Square) GetArea() float64 {
	return float64(s.width * s.width)
}

type Circle struct {
	shape
	radius int
}

func NewCircle(color string, radius int) *Circle {
	return &Circle{
		shape:  shape{color},
		radius: radius,
	}
}

func (c *Circle) GetArea() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2)
}
