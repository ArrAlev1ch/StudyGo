package main

import "fmt"

type User struct {
	Name string
	Age int
	Email string
}

func (u User) IsAdult() bool {
	IsTrue := false
	if u.Age >= 18 {
		IsTrue = true
		fmt.Println(IsTrue)
	} else {
		fmt.Println(IsTrue)
	}
	return IsTrue
}

func main() {
	u := User{"Вася", 18, "lox228@gmail.com"}
	u.IsAdult()
}