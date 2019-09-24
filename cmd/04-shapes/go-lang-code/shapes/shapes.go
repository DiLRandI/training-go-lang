package shapes

import "math"

type Shape interface {
	GetArea() float64
}

type stringColor string

func (c stringColor) GetColor() string {
	return string(c)
}

type Square struct {
	stringColor
	width int
}

func NewSquare(color string, width int) *Square {
	return &Square{
		stringColor: stringColor(color),
		width:       width,
	}
}

func (s *Square) GetArea() float64 {
	return float64(s.width * s.width)
}

type Circle struct {
	stringColor
	radius int
}

func NewCircle(s string, radius int) *Circle {
	return &Circle{
		stringColor: stringColor(s),
		radius:      radius,
	}
}

func (c *Circle) GetArea() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2)
}
