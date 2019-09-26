package main

import "fmt"

func main() {
	defer fmt.Println("Done printing")
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	defer fmt.Println("Start printing")
	fmt.Println("done")
}
