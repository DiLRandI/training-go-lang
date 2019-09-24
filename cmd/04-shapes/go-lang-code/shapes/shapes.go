package shapes

type Shape interface {
	GetArea() float64
}

type color struct {
	Color string
}
type Square struct {
	color
	width int
}

type Circle struct {
	color
	radius int
}
