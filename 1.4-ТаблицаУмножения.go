package main

import "fmt"

func main() {

	for y := 1; y <= 10; y++ {
		for x := 1; x <= 10; x++ {
			fmt.Printf("%v x %v = ", y, x); fmt.Println(x*y)
		}
	}
}