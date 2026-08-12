package main

import "fmt"

func ReverseSlice(rs []int) []int {
	res := make([]int, len(rs))
	for i := range rs {
		res[len(rs)-(i+1)] = rs[i]
	}
	return res
}

func main() {
	s := []int{1, 2, 3, 4}
	rev_s := ReverseSlice(s)
	fmt.Println(rev_s)
}