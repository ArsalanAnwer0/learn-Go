package main

import "fmt"

func main() {
	// age := 18

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	age := 16

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("Person is a teenager")
	}

	var role = "admin"
	var hasPermission = false

	if role == "admin" && hasPermission {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	// we can declare a variable inside if construct
	if age := 19; age > 18 {
		fmt.Println("Person is an adult", age)
	}

	// go does not have ternary operator, but we can achieve similar result using if else

}
