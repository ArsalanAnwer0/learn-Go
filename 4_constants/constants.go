package main

import "fmt"

const year = 2026

func main() {
	const name string = "go lang"
	const age = 30

	fmt.Println(year)

	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)

}
