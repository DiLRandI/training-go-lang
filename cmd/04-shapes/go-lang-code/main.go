package main

import (
	"github.com/dilrandi/training-go-lang/cmd/04-shapes/go-lang-code/shapes"
)

func main() {

}

func AreaCalculator(shapes []shapes.Shape) float64 {
	var ta float64
	for _, s := range shapes {
		ta += s.GetArea()
	}
	return ta
}
