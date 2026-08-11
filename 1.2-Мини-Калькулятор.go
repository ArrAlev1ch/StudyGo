package main

import "fmt"

func main() {

	var x, y int
	fmt.Println("Введите два числа:")
	fmt.Scan(&x, &y)
	fmt.Printf("%v + %v = ", x, y); fmt.Println(x + y)

}