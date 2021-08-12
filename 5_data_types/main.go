package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*
		Built in types :
		int8, int16, int32, int64
		uint8, uint16, uint32, uint64 (int refers to either int32/int64 depending upon arch)
		float32, float64
		complex64, complex128
		byte (alias for unit8) (used for storing characters, as go does not have a char data type)
		rune (alias for int32)
		bool
		string
		array
		map
		struct
		pointer
		func (function)
		interfaces
		channel
	*/

	var c rune = 'a'
	fmt.Printf("%T %v\n", c, c)

	var s string = "Hello go"
	fmt.Printf("%T\n", s)

	// array -> fixed size, slice -> flexible size
	var arr = [2]int{4, 5}
	var slc = []int{5, 6, 7}
	fmt.Printf("%T %T\n", arr, slc)

	var m = map[int]string{
		1: "Konichiwa",
		2: "Arigato",
	}
	fmt.Printf("%T, %v\n", m, m[2])

	type Person struct {
		name string
		age  uint
	}

	var p1 Person
	fmt.Printf("%T %v %v\n", p1, p1.name, p1.age)

	var x int = 2
	var ptr *int = &x
	fmt.Printf("%T %v\n", ptr, ptr)

	// For type conversion, use type functions, for eg -> f=2.4. int64(f) -> conversion
	// For string to int, float, bool use strconv package functions
	// For int, float, bool to string, use fmt.Sprintf()

	f1, err := strconv.ParseFloat("56.77", 64)
	_ = err
	fmt.Println(f1)

	// Defined types :

	type age uint8
	var myage age = 24
	fmt.Printf("%T\n", myage)

	var yourAge uint8 = uint8(myage)

	var ourAge age = age(yourAge)
	_ = ourAge

	// Aliases (byte and rune are also aliases)
	type height = float32
}
