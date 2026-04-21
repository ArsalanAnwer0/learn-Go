// compiled? -> interpreted?
// go lang is compiled.
// Go tool can run file directly, without VM
// Executables are different for OS.
// what and where can I use go lang? system apps to web app - Cloud
// object oriented -> Yes and No -> Structs -> no operating loading.
// prev -> setup path, now -> run code from anywhere..
// What is Go Path..
// Lexer? ->
// types -> private or public,
// data types -> string, bool, integer, floating, complex, array, slices, maps, structs and
// pointers
// Function and Channels

package main

import (
	"bufio"
	"fmt"
	"os"
)

const LoginToken string = "blah" // capital L -> public variable

func main() {
	// fmt.Println("Arsalan")

	// variables:
	var username string = "Arsalan"
	var isLoggedIn bool = true
	fmt.Println(username)
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", username)

	var smallFloat float64 = 255.455444353
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases -> always zero, initialized and not assigned
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type -> if not mentioned, i will decide for you (lexer)
	var website = "learncodeonline"
	fmt.Println(website)

	// no var style, not allowed outside
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

	// user input using package bufio
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	// comma ok | error ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of rating is %T", input)

}
