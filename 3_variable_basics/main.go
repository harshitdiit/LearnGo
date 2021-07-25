package main

import "fmt"

// Package scoped variables can be declared and not used, but shorthand declaration is not allowed
var global_var int = 8

func main() {
	// Statically declared variable
	var age int = 10
	// Dynamically declare variable
	rank := 4
	// Assign to blank to avoid compiile time error, if declared variable not used. (Use with caution)
	name := "Harshit"
	_ = name

	fmt.Printf("%d %d \n", age, rank)

	// Can skip type in variables declared with var
	var car = "Audi"
	// Multiple declaration (atleast one value on the left should be new)
	car, bike, cycle := "BMW", "Husqvarna", "BMX"
	opened, fruit := true, "mango"
	_, _ = opened, fruit

	fmt.Println(car, bike, cycle)

	// Multiple declaration with var
	var (
		salary    float64
		firstName string
		gender    bool
	)

	var a, b, c int

	fmt.Println(salary, firstName, gender)
	fmt.Printf("%d %d %d\n", a, b, c)

	// Multiple assignment
	var i, j int
	i, j = 0, 1
	i, j = j, i

	fmt.Printf("%d %d\n", i, j)

	// Expressions work as well
	sum := i + 5.0

	fmt.Println(sum, global_var)
}
