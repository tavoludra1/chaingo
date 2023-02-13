package main

import "fmt"

func main() {
	var fname string
	fmt.Println("Enter first name: ")
	fmt.Scanln(&fname)

	var lname string
	fmt.Println("Enter last name: ")
	fmt.Scanln(&lname)

	fmt.Printf("Full name is: %s %s", fname, lname)

}
