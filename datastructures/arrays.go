package datastructures

import "fmt"

/*
	Arrays store a list of values where each value
	is referenced by an index or a key.
*/
func Arrays() {
	fmt.Println("Array Functions")
	/*
		We'll start by creating an array that is 5 indexes in length
		and contains only integers. Meaning it is a homogeneous structure,
		or contains only one type.
	*/
	var array [5]int
	fmt.Println("Empty array: ", array)

	// We can set the individual indexes with specific values
	// Here we are setting the 4th ( last ) index with the integer 100
	array[4] = 100
	fmt.Println("Set array: ", array)
	// Here we'll print out the specific index reference
	fmt.Println("Get array: ", array[4])

}

// Redundant, but for testing we'll use this to return a new array
func ReturnNewArray(array []int) []int {
	return array
}
