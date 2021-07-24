package algorithms

/*
	MergeSort sorts a list in ascending order
	Returns a new sorted list
*/
func MergeSort(list []int) []int {
	/*
		Steps:
		- Divide the midpoint of the list and divide into sublists
		- Recursively sort the sublists created in previous step
		- Merge the sorted sublists created in previous step
	*/

	if len(list) <= 1 {
		return list
	}
	first := MergeSort(list[:len(list)/2])
	second := MergeSort(list[len(list)/2:])
	return Merge(first, second)

}

/*
	This is where the magic really happens. At the lowest level
	of recursion, the two 'sorted' lists will each have a length
	of 1. Those single elemnt lists will be merged into a sorted
	list of length two, and we can build up from there.
*/
func Merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
