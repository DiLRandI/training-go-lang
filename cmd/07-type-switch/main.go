package main

import "fmt"

func main() {
	do(21)
	do("hello")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("This is a int : %v \n", v)
	case string:
		fmt.Printf("This is a string : %q \n", v)
	default:
		fmt.Printf("This is a different type %T!\n", v)
	}
}
