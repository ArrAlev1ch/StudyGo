package main

import "fmt"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль запрещено")
	}
	return a / b, nil
}

func main() {
	result, error := Divide(10, 0)

	if error != nil {
		fmt.Println("Ошибка:", error)
		return
	}
	fmt.Println("Ответ:", result)
}