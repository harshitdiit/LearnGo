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

	// Go always initializes variables with default values to avoid uninitialized errors
	// int->0, float64->0, pointer->nil, string->""

	/*
		Naming converntions in go :
		Keywords :
		break        default      func         interface    select
		case         defer        go           map          struct
		chan         else         goto         package      switch
		const        fallthrough  if           range        type
		continue     for          import       return       var

		Variable Naming :
		Maximum Value -> mv, or maxValue
		Go prefers non-verbose code
	*/

	// To declare constants we use const keyword
	// You can use the program without using a constant
	// Constants cannot be uninitialized

	const days int = 7 // Typed constant
	const x, y int = 7, 8
	const z = 9.81 // Untyped constant

	// In grouped constants, value is picked from previous const
	const (
		min1 = 400
		min2
	)

	fmt.Println(min1, min2)

	// Iota generates incremental integer constant values within a scope
	const (
		c1 int = iota
		c2
		c3
	)

	fmt.Println(c1, c2, c3)

	// Iota usage example
	const (
		c11 int = iota*2 + 1
		c22
		c33
	)

	fmt.Println(c11, c22, c33)

	// Skipping values
	const (
		f1 int = -(iota + 2)
		_
		f2
		f3
	)

	fmt.Println(f1, f2, f3)
}
