package main

import "fmt" 

type Address struct {
	Город, Улица string
}

type User struct {
	Имя string
	Address
}

func main() {
	u := User{"Nilyaz", Address{"Kazan", "Main"}}
	fmt.Printf("Имя: %s, Город: %s, Улица: %s", u.Имя, u.Город, u.Улица)
}