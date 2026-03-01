package main

import "fmt"

func main() {
	age := 30

	var agePtr *int
	agePtr = &age

	fmt.Println("Age pointer derefencing:", *agePtr)
	fmt.Println("Age:", age)

	editAge(agePtr)

	fmt.Println("Age pointer derefencing:", *agePtr)
	fmt.Println("Age:", age)
}

func editAge(agePtr *int) {
	*agePtr = *agePtr - 12
}
