package main

import "fmt"

func main() {
	i := 10
	fmt.Println("Value of i:", i)
	p := &i
	fmt.Println("Value of p:", p)
	fmt.Println("Value of *p:", *p)

	change(i)
	fmt.Println("Change value of i:", i)

	changePointer(p)
	fmt.Println("Using pointer change value of i:", i)
}

func change(x int) {
	x = 20
}

func changePointer(x *int) {
	*x = 25
}
