package algorithms

import "fmt"

/*
	Binary Search
		- Start by defining the start and end indexes of our list
		- Start a loop if our first index is not greater than our last
		- In the loop determine our midpoint
		- Run a check against the current midpoint to the target
		- If it's not a match, set our new first and last indexes
		- Let the loop run until if narrows down enough to find the target
*/
func BinarySearch(list []int, target int) string {
	fmt.Println("==========================")
	fmt.Printf("Starting Binary Search..\n")
	fmt.Printf("Finding target: %d in list %d \n", target, list)

	// declare the first, and last indexes of the list
	first := 0
	last := len(list) - 1

	for first <= last {
		/*
			Find the midpoint of the list by adding the first and last items.
			Ex. [0..6] translate to ( 0 + 6 ) / 2 = 3. The midpoint = 3
		*/
		midpoint := (first + last) / 2

		fmt.Printf("%d..", list[midpoint])

		// If we hit our target on the first iteration, great we're done
		if list[midpoint] == target {
			fmt.Printf("\n")
			return fmt.Sprintf("Target %d found \n", target)
			// Otherwise if our midpoint is less than the target
			// set our new first index to our previous index + 1.
		} else if list[midpoint] < target {
			first = midpoint + 1
			// Our midpoint is less than our target, reduce the range of our last index by one.
		} else {
			last = midpoint - 1
		}

	}

	return fmt.Sprintf("Target %d not found \n", target)
}
