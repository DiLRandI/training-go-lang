package main

import (
	"fmt"

	"github.com/dilrandi/training-go-lang/cmd/09-shapes/go-lang-code/shapes"
)

func main() {
	var shapeList []shapes.IShape

	c := shapes.NewCircle("Red", 2)
	s := shapes.NewSquare("Green", 2)

	shapeList = append(shapeList, c)
	shapeList = append(shapeList, s)

	fmt.Println(AreaCalculator(shapeList))
}

func AreaCalculator(shapes []shapes.IShape) float64 {
	var ta float64
	for _, s := range shapes {
		ta += s.GetArea()
	}
	return ta
}
