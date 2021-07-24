package main

import (
	"fmt"

	"github.com/DLzer/ds-and-algos/algorithms"
)

func main() {

	/*
		Linear Search
		 - Start at beginning
		 - Compare current value to target
		 - Move Sequentially
		 - Reach end of list
	*/
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	response := algorithms.LinearSearch(items, 58)
	fmt.Println(response)

	/*
		Binary Search
		 - Start by defining the start and end indexes of our list
		 - Start a loop if our first index is not greater than our last
		 - In the loop determine our midpoint
		 - Run a check against the current midpoint to the target
		 - If it's not a match, set our new first and last indexes
		 - Let the loop run until if narrows down enough to find the target
	*/
	items = []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	response = algorithms.BinarySearch(items, 63)
	fmt.Println(response)

	/*
		Recursive Binary Search
		- Similary to binary search we start with a list and target
		- The primary difference is through each iteration if we
		find that the target is not within the bounds we call the
		function again with a new list that removes the excess.
	*/
	items = []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	response = algorithms.RecursiveBinarySearch(items, 63)
	fmt.Println(response)

}
