package main

import (
	"fmt"
)

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	hobbies := [3]string{"Reading", "Cooking", "Traveling"}
	fmt.Println("Hobbies:", hobbies)

	// print the first hobby
	fmt.Println("First hobby:", hobbies[0])

	// print second and third hobbies as a slice or new array
	fmt.Println("Second and third hobbies:", hobbies[1:3])

	// select the main hobby
	mainHobbies := hobbies[0:2]
	fmt.Println("Main hobbies:", mainHobbies)

	// main hobbies as a new slice containing the last two hobbies
	mainHobbies = mainHobbies[1:3]
	fmt.Println("Main hobbies (last two):", mainHobbies)

	// create course goals array
	courseGoals := []string{"Learn Go", "Build Projects", "Contribute to Open Source"}
	fmt.Println("Course Goals:", courseGoals)

	// change the second course goal
	courseGoals[1] = "Build Real-World Projects"
	fmt.Println("Updated Course Goals:", courseGoals)

	// append a new course goal
	courseGoals = append(courseGoals, "Network with Other Developers")
	fmt.Println("Final Course Goals:", courseGoals)

	// create an array of products
	products := []Product{
		{id: "p1", title: "Laptop", price: 999.99},
		{id: "p2", title: "Smartphone", price: 499.99},
		{id: "p3", title: "Headphones", price: 199.99},
	}
	fmt.Println("Products:", products)

	// create a new product and add it to the products array
	newProduct := Product{id: "p4", title: "Smartwatch", price: 299.99}
	products = append(products, newProduct)
	fmt.Println("Updated Products:", products)
}
