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
	"strconv"
	"strings"
	"time"
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

	// conversions
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader_two := bufio.NewReader(os.Stdin)
	input_two, _ := reader_two.ReadString('\n')

	fmt.Println("Thanks for rating, ", input_two)
	numRating, errr := strconv.ParseFloat(strings.TrimSpace(input_two), 64)

	if errr != nil {
		fmt.Println(errr)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

	// handling time
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.November, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

	// memory management allocation and deallocation happens automatically
	// two methods
	// new() -> allocate memory but no INIT, get memory address, zeroed storage -> no data
	// make()-> allocate memory and INIT, get memory address, non-zeroed storage -> data
	// Garbage Collection happens automatically,
	// what is runtime package -> CPU available..

	// Pointers
	fmt.Println("Welcome to a class on pointers")

	var num int = 5
	var ptr *int = &num
	var ptr2 = &num
	*ptr = *ptr * 2
	fmt.Println("Value of pointer is: ", *ptr)
	fmt.Println("Value of pointer is: ", *ptr2)
	fmt.Println("Value of num is: ", num)

	// Array
	fmt.Println("Welcome to array in golangs")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[3] = "Peach"

	fmt.Println("Fruit List is: ", fruitlist)
	fmt.Println("Fruit List is: ", len(fruitlist))

	var vegList = [3]string{"Potato", "beans", "mushroom"}
	fmt.Println("Vegy list is ", vegList)
	fmt.Println("Vegy list is ", len(vegList))

}
