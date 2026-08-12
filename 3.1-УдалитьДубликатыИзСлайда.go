package main

import "fmt"

func main() {

	s := []int{1, 2, 2, 3, 1, 4}
	m := make(map[int]bool)
	res := []int{}

	for i := range s {

		if m[s[i]] == false {
			res = append(res, s[i])
			m[s[i]] = true
		} 

	}

	fmt.Println(res)

}