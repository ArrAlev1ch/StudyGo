package main

import "fmt"

func Symbols(s string) map[string]int {
	res := make(map[string]int)
	for _, i := range s {
		if i != ' ' {
			res[string(i)]++
		}
	}
	return res
}

func main() {
	str := "hello world"
	str_symbols := Symbols(str)
	
	for i, j := range str_symbols {
		fmt.Printf("%s:%d, ", i, j)
	}
}