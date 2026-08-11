package main

import "fmt"

type User struct {
	Name string
	Age int
	Email string
}

func main() {
	u := User{"Вася", 21, "Lox228@gmail.com"}
	fmt.Printf("%+v", u)
}