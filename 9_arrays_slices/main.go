package main

import (
	"fmt"
)

func main() {

	// declaring an array with four elements of type int
	var numbers [4]int

	// array zero value is zeroed value elements
	fmt.Printf("%v\n", numbers)  // -> [0 0 0 0]
	fmt.Printf("%#v\n", numbers) // -> [4]int{0, 0, 0, 0}

	// declaring an array and initialize it using an array literal
	var a1 = [4]float64{}                           //initialized with defaults (0)
	var a2 = [3]int{5, -3, 5}                       //initialized with the given values
	a3 := [4]string{"Dan", "Diana", "Paul", "John"} //short declaration syntax
	a4 := [4]string{"x", "y"}                       //initializing only the first 2 elements

	// the ellipsis operator (...) finds out automatically the length of the array
	a5 := [...]int{1, 4, 5}
	fmt.Println("The length of a5 is: ", len(a5)) // len is 3

	// declare an array on multiple lines for better readability
	a6 := [...]int{1,
		2,
		3,
	} //the ending comma is mandatory when initializing the array on multiple lines and the closing curly brace is on its own line

	_, _, _, _, _, _ = a1, a2, a3, a4, a5, a6

	// changing an array
	// we can't add or remove elements from the array (it's fixed length)
	numbers[0] = 7              //changing first element (index 0)
	fmt.Printf("%v\n", numbers) // -> [7 0 0 0]

	// compile-time error
	// numbers[5] = 8  // invalid array index 5 (out of bounds for 4-element array)

	// getting an element
	x := numbers[0]
	fmt.Println("x is ", x) // => x is  7

	// iterating over an array (2-ways)
	for i, v := range numbers { // range is a language keyword used for iteration
		fmt.Println("index:", i, "value: ", v)

	}

	// iterating over an array (C/C++, Java Style)
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index:", i, "value: ", numbers[i])
	}

	// declaring a multi-dimensional arrays (array of arrays or matrix)
	balances := [2][3]int{
		[3]int{5, 6, 7}, //[3]int is optional
		{8, 9, 10},
	}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println("")
	}

	//  = operator creates a copy of an array.
	// the arrays are not connected and are saved in different memory locations
	m := [3]int{1, 2, 3}
	mcp := m //n is a copy of m

	fmt.Println("n is equal to m: ", mcp == m) // => true
	m[0] = -1                                  //only m is modified
	fmt.Println("n is equal to m: ", mcp == m) // => false

	// each key corresponds to an index of the array
	grades := [3]int{ //the keyed elements can be in any order
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println(grades) // -> [5 10 7]

	accounts := [3]int{
		2: 50,
	}
	fmt.Println(accounts) //[0 0 50]

	names2 := [...]string{
		4: "Dan",
	}
	fmt.Println(len(names2)) // -> 5

	// un unkeyed element gets its index from the last keyed element
	cities := [...]string{
		5:        "Paris",
		"London", // this is at index 6
		1:        "NYC",
	}
	fmt.Printf("%#v\n", cities) // -> [7]string{"", "NYC", "", "", "", "Paris", "London"}

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend) // => [false false false false false true true]

	// declaring a string slice, by default is initialized with nil or uninitialized
	var cities2 []string

	fmt.Println("cities is equal to nil: ", cities2 == nil) // -> cities is equal to nil:  true
	fmt.Printf("cities: %#v\n", cities)                     // -> cities: []string(nil)

	// we can not assign elements to nil slice:
	// cities[0] = "Paris" // -> runtime error

	// declaring a slice using a slice literal
	numbers2 := []int{2, 3, 4, 5} // on the right hand-side of the equal sign is a slice literal
	fmt.Println(numbers2)         // => [2 3 4 5]

	// creating a slice using the make() built-in function
	// creating a slice with 2 int elements initialized with zero.
	nums := make([]int, 2) // the same as []int{0, 0}.
	fmt.Println(nums)      // => [0 0]

	// declaring a slice using a slice literal
	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	// getting an element from the slice
	x2 := numbers2[0]
	fmt.Println("x is", x2) // => x is 2

	// modifying an element of the slice
	numbers2[1] = 200
	fmt.Printf("%#v\n", numbers2) // => []int{2, 200, 4, 5}

	// iterating over a slice
	for index, value := range numbers2 {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	//iterating over a slice (C/C++, Java Style)
	for i := 0; i < len(numbers2); i++ {
		fmt.Printf("index: %v, value: %v\n", i, numbers2[i])

	}

	// slices with the same element type can be assigned to each other
	var n []int
	n = numbers2
	_ = n

	//** COMPARING SLICES **//
	// a Go slice can only be compared to nil

	// uninitialized slice, equal to nil
	var nn []int
	fmt.Println(nn == nil) // true

	// empty slice but initialized, not equal to nil
	mm := []int{}
	fmt.Println(mm == nil) //false

	// we can not compare slices using the equal (=) operator
	// fmt.Println(nn == mm) //error -> slice can only be compared to nil

	// to compare 2 slices use a for loop to iterate over the slices and compare element by element
	a, b := []int{1, 2, 3}, []int{1, 2, 3}

	var eq bool = true

	for i, valueA := range a {
		if valueA != b[i] {
			eq = false // don't check further, break!
			break
		}
	}

	// it's also necessary to check the length of slices (if a is nil it doesn't enter the for loop)
	if len(a) != len(b) {
		eq = false
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}

	// arrays, slices and strings are sliceable
	// slicing doesn't modify the array or the slice, it returns a new one

	// declaring an [5]int array
	a7 := [5]int{1, 2, 3, 4, 5}

	// a slice expression is formed by specifying a start or a low bound and a stop or high bound like  a[start:stop].
	// this selects a range of elements which includes the element at index start, but excludes the element at index stop.

	// slicing an array returns a slice, not an array
	b2 := a7[1:3]                                 // 1 is called start (included), 3 is called stop (excluded)
	fmt.Printf("Type: %T , Value: %#v\n", b2, b2) // => Type: []int , Value: []int{2, 3}

	// declaring a slice
	s1 := []int{1, 2, 3, 4, 5, 6}

	s2 := s1[1:3]   //start included, stop excluded
	fmt.Println(s2) //[2 3]

	//for convenience, any of the indexis may be omitted.
	// a missing low index defaults to zero; a missing high index defaults to the length of the sliced operand.
	fmt.Println(s1[2:])       // => [3 4 5 6], same as s1[2:len(s1)]
	fmt.Println(s1[:3])       // => [1 2 3], same as s1[0:3]
	fmt.Println(s1[:])        // => [1 2 3 4 5 6], same with s1[0:len(s1)]
	fmt.Println(s1[:len(s1)]) // => => [1 2 3 4 5 6], returns the entire slice
	// fmt.Println(s1[:45])   //panic: runtime error: slice bounds out of range

	s1 = append(s1[:4], 100) // adds 100 after index 4 (excluded)
	fmt.Println(s1)          // -> [1 2 3 4 100]

	// a slice expression doesn't create a new backing array. The original and the returned slice are connected!
	s3 := []int{10, 20, 30, 40, 50}
	s4, s5 := s3[0:2], s3[1:3] //s4, s5 share the same backing array with s3

	s4[1] = 600     // modifying the backing array so s3, s4 and s5 are in fact modified!!
	fmt.Println(s3) // -> [10 600 30 40 50]
	fmt.Println(s5) // -> [600 30]

	// when a slice is created by slicing an array, that array becomes the backing array of the new slice.
	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	arr1[1] = 2                       // modifying the array
	fmt.Println(arr1, slice1, slice2) // -> [10 2 30 40] [10 2] [2 30]

	// append() function creates a complete new slice from an existing slice
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	// newCars doesn't share the same backing array with cars
	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"                              // only cars is modified
	fmt.Println("cars:", cars, "newCars:", newCars) // => cars: [Nissan Honda Audi Range Rover] newCars: [Ford Honda]

	// slice is a struct with three fields -> length, capacity and a pointer to the backing array

	numbers3 := []int{2, 3}

	// append() returns a new slice after appending a value to its end
	numbers3 = append(numbers3, 10)
	fmt.Println(numbers) //-> [2 3 10]

	// appending more elements at once
	numbers3 = append(numbers3, 20, 30, 40)
	fmt.Println(numbers) //-> [2 3 10 20 30 40]

	// appending all elements of a slice to another slice
	n2 := []int{100, 200, 300}
	numbers3 = append(numbers3, n2...) // ... is the ellipsis operator
	fmt.Println(numbers3)              // -> [2 3 10 20 30 40 100 200 300]

	//** Slice's Length and Capacity **//

	nums2 := []int{1}
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums2), cap(nums2)) // Length: 1, Capacity: 1

	nums2 = append(nums2, 2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums2), cap(nums2)) // Length: 2, Capacity: 2

	nums2 = append(nums2, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums2), cap(nums2)) // Length: 3, Capacity: 4
	// the capacity of the new backing array is now larger than the length
	// to avoid creating a new backing array when the next append() is called.

	nums2 = append(nums2, 4, 5)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums2), cap(nums2)) // Length: 5, Capacity: 8

	// copy() function copies elements into a destination slice from a source slice and returns the number of elements copied.
	// if the slices don't have the same no of elements, it copies the minimum of length of the slices
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn2 := copy(dst, src)
	fmt.Println(src, dst, nn2) // => [10 20 30] [10 20 30] 3
}
