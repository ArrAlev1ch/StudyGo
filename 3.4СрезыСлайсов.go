package main

import "fmt"

func Subslice(s []int, start, end int) []int {
	if  start > end {
		return []int{}
	}
	if  end > len(s) {
		return []int{}
	}
	return s[start:end]
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sub := Subslice(slice, 0, 3)
	fmt.Println(sub)
}