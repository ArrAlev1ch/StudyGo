package main

import "fmt"

func main() {
	
	var i int
	
	fmt.Print("Введите число: "); fmt.Scan(&i)
	
	if i % 2 == 1 {
		fmt.Println("Нечетное")
	}
	
	if i % 2 == 0 {
		fmt.Println("Четное")
	}
}