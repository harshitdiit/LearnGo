package main

import "fmt"

func main() {

	// Println -> Concatenates arguements with a whitespace and prints with a new line at the end

	var name string = "Println"
	fmt.Println("This is an example", "of how", name, "works")

	// Printf -> Behaves like C printf
	// %d -> int, %f -> float, %s -> string, %q -> quoted string, %v -> any value, %T -> variable type, %t -> bool, %b -> binary representation
	// %x -> hexadecimal representation

	fmt.Printf("Variable name is of type: %T\n", name)
}
