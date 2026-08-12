package main

import "fmt"

func RemoveDuplicate(s []int) []int {

	m := make(map[int]bool)
	res := []int{}

	for _, i := range s {

		if m[i] == false {
			res = append(res, i)
			m[i] = true
		} 
	}
	return res
}

func main() {

	s := []int{1, 2, 2, 3, 1, 4}
	rem_s := RemoveDuplicate(s)
	fmt.Println(rem_s)

}
