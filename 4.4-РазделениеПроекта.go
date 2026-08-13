package main

import (
	"fmt"
	"TEST/mathutils"
)

func main() {
	fmt.Println(utils.Add(3, 4))
	fmt.Println(utils.Multiply(3, 4))
}

/* в папке src/TEST/mathutils, файл math.go

package utils

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

*/
