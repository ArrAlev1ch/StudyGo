package main

import "fmt"

type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() {
	fmt.Println(r.length * r.width)
}

func main() {
	rec := Rectangle{3, 6}
	rec.Area()
}