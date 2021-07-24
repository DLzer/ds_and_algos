package algorithms

import "fmt"

/*
	Recursive Binary Search
	- Similary to binary search we start with a list and target
	- The primary difference is through each iteration if we
	find that the target is not within the bounds we call the
	function again with a new list that removes the excess.
*/
func RecursiveBinarySearch(list []int, target int) string {

	fmt.Printf("Starting Recursive Binary Search..\n")
	fmt.Printf("Finding target: %d in list %d \n", target, list)

	low := 0
	high := len(list) - 1

	if low <= high {
		mid := ((high + low) / 2)

		if list[mid] > target {
			return RecursiveBinarySearch(list[:mid], target)
		} else if list[mid] < target {
			return RecursiveBinarySearch(list[mid+1:], target)
		} else {
			fmt.Printf("\n")
			return fmt.Sprintf("Target %d found \n", target)
		}
	}

	return fmt.Sprintf("Target %d not found \n", target)
}
