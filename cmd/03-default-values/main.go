package main

//https://tour.golang.org/basics/12
import "fmt"

//Interface ...
type Interface interface{}

//Struct ...
type Struct struct{}

//Person ...
type Person struct {
	Name string
	Age  int
}

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var intf Interface
	var stuc Struct
	var p Person
	var pp *Person
	fmt.Printf("%v %v %v %q %v %v %v %v\n", i, f, b, s, intf, stuc, p, pp)
}
