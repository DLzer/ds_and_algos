package algorithms

import "fmt"

func LinearSearch(list []int, target int) string {
	fmt.Printf("Starting Linear Search..\n")
	fmt.Printf("Finding target: %d in list %d \n", target, list)

	// Start by iterating through our list
	for _, item := range list {

		fmt.Printf("%d..", item)

		// If the current item in the list is equal to our target, we're done!
		// Otherwise continue iterating through each item until we match.
		if item == target {
			fmt.Printf("\n")
			return fmt.Sprintf("Target of %d found \n", target)
		}
	}

	return fmt.Sprintf("Target %d not found", target)
}
