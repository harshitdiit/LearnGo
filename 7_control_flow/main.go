package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// if condition_that_evaluates_to_boolean {
	//      perform action1
	// } else if condition_that_evaluates_to_boolean {
	//      perform action2
	// } else {
	//      perform action3
	// }

	price := 100

	if price < 100 {
		fmt.Println("It's cheap!")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("It's Expensive!")
	}

	// simple (short) statement : initializes variables for if condition
	// i and err are variables scoped to the if statement only
	if i, err := strconv.Atoi("34"); err == nil {
		fmt.Println("No error, i is", i)
	} else {
		fmt.Println(err)
	}

	fmt.Println(os.Args[0])

	// printing numbers from 0 to 9
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// has the same effect as a while loop in other languages
	// there is no while loop in Go
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	// handling of multiple variables in a for loop
	for i, j := 0, 100; i < 10; i, j = i+1, j+1 {
		fmt.Printf("i = %v, j = %v\n", i, j)
	}

	// infinite loop
	// sum := 0
	// for {
	//  sum++
	// }

	// printing even numbers less than or equal to 10
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // skipping the remaining code in this iteration
		}
		fmt.Println(i)
	}

	count := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", i)
			count++
		}

		if count == 10 { //if 10 numbers were found, break!
			break //it breaks the current loop (inner loop if there are more loops)
		}
	}

	//** LABEL STATEMENT **//

	// There is no conflict name between label names and identifiers as they reside in separate places in memory
	// Used to define scopes and jump between pieces of codes. works with break and continue
	outer := 19
	_ = outer

	// declaring two arrays
	people := [5]string{"Helen", "Mark", "Brenda", "Antonio", "Michael"}
	friends := [2]string{"Mark", "Marry"}

	// searching for a single friend in a list of people.

outer: //label, it doesn't conflict with other names
	// iterating over the array.
	for index, name := range people { // range returns both the index and the elements of the array one by one
		for _, friend := range friends { //iterating over the second array
			if name == friend {
				fmt.Printf("FOUND A FRIEND: %q at index %d\n", friend, index)
				break outer //breaking outside the outer loop which terminates
			}
		}
	}

	fmt.Println("Next instruction after the break.")

	//the following piece of code creates a loop like a for statement does
	i := 0
loop: // label
	if i < 5 {
		fmt.Print(i)
		i++
		goto loop
	}

	//  goto todo //ERROR it's not permitted to jump over the declaration of a new variable
	//  x := 5
	// todo:
	//  fmt.Println("something here")

	/* SWITCH-CASE */

	language := "golang"

	switch language {
	case "Python": //values must be comparable (compare string to string)
		fmt.Println("You are learning Python! You don't use { } but indentation !! ")
		// an implicit break is added here
	case "Go", "golang": //compare language with "Go" OR "golang"
		fmt.Println("Good, Go for Go!. You are using {}!")
	default:
		// the default clause the equivalent of the else clause of an if statement
		// and gets executed if no testing condition is true.
		fmt.Println("Any other programming language is a good start!")
	}

	n := 5
	// comparing the result of an expression which is bool to another bool value
	switch true {
	case n%2 == 0:
		fmt.Println("Even!")
	case n%2 != 0:
		fmt.Println("Odd!")
	default:
		fmt.Println("Never here!")
	}

	//** Switch simple statement **//

	// Syntax: statement (n:=10), semicolon and a switch condition
	//(true in this case, we are comparing boolean expressions that return true)
	// we can remove the word "true" because it's the default
	switch n := 10; true {
	case n < 20:
		fmt.Println("Nice")
	case n < 15:
		fmt.Println("Wow")
	default:
		fmt.Println("Zero")
	}
}
