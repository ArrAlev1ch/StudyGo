package main

import "fmt"

type User struct{
	Name string
}

type ValidationError struct {
	Field string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("Ошибка валидации: пустое поле %s", v.Field)
}

func ValidateUser(u User) error {
	if u.Name == "" {
		return ValidationError{Field: "Name"}
	}
	return nil
}

func main() {
	user1 := User{Name: ""}
	err := ValidateUser(user1)
	if err != nil {
		fmt.Println(err)
	} 
}