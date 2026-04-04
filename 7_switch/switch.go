package main

import (
	"fmt"
	"time"
)

func main() {
	// switch statement is used to perform different actions based on different conditions
	// i := 7

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// case 4:
	// 	fmt.Println("four")
	// case 5:
	// 	fmt.Println("five")
	// default:
	// 	fmt.Println("other")
	// }

	// multiple condition in switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's the weekend!")
	// default:
	// 	fmt.Println("It's a weekday.")
	// }

	// type switch

	whoamI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Printf("I am an int and my value is %d\n", i)
		case string:
			fmt.Printf("I am a string and my value is %s\n", i)
		case bool:
			fmt.Printf("I am a bool and my value is %t\n", i)
		case time.Time:
			fmt.Printf("I am a time and my value is %s\n", i)
		default:
			fmt.Printf("I am of type %T and my value is %v\n", t, i)
		}
	}

	whoamI(3.5)

}
