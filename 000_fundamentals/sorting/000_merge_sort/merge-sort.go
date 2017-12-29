package main

// Mergesort is a divide-and conquer algorithm that is essentially
// based on the idea of splitting the collection to sort into halves
// and recursively mergin them.
//
// See [Merge Sort](https://en.wikipedia.org/wiki/Merge_sort) for additional
// details.

// Sorts all the items between `l` and `r` indices (inclusive) of the array.
// This implementation of `mergeSort` mutates the original array.
func mergeSort(ar []int, l int, r int) {
	// If we are sorting one or less items, then we can return since
	// a single item is already sorted by itself.
	if l >= r {
		return
	}

	// Find the middle point
	m := (l + (r - 1)) / 2

	// Sort the left half
	mergeSort(ar, l, m)

	// Sort the right half
	mergeSort(ar, m+1, r)

	// Merge the sorted parts back into the array.
	merge(ar, l, m, r)
}

// Add some details about the algorithm.

// The `mergeSort` function is not “that” complicated since the meat
// of the algorithm is about how we merge things, rather than how we
// sort things. So the `merge` part it a bit more involved.
func merge(ar []int, l int, m int, r int) {
	n1 := m - l + 1
	n2 := r - m

	// Create two temp arrays for left and right halves
	left := make([]int, 0)
	right := make([]int, 0)

	// Note that there are also “in place” `merge` implementations; however
	// it makes the algorithm relatively harder to read, so we’ll stick
	// with using temporary storage. — In an interview setup, this
	// current implementation of merge will be adequate for you.
	// That said, you might want to look up the “in place” merge algorithm
	// (ref: http://penguin.ewu.edu/cscd300/Topic/AdvSorting/MergeSorts/InPlace.html)
	// so that at least you can have an idea about it, if you are asked in an interview.
	//
	// For the “in place” merge logic (which is **not** the logic we use here), as usual,
	// you have a subscript into the left array segment (x[left]) and the right array
	// segment (x[right]).  If x[left] tests as less than or equal to x[right], then it is already
	// in place within the sorted array segment, so just increment left.  Otherwise,
	// the array element in x[right] needs to move down into the space currently occupied by x[left],
	// and to accommodate this, the entire array segment from x[left] through x[right-1] needs to move up by one —
	// effectively you need to rotate that little segment of the array.  In the process,
	// everything moves up by one, including the end of the left segment (mid).
	//
	// Anyways, back to our code…

	// Initialize `left` and `right` slices.
	//
	// n1: Size of the left temp array.
	// n2: Size of the right temp array.
	for i := 0; i < n1; i++ {
		left = append(left, ar[l+i])
	}
	for i := 0; i < n2; i++ {
		right = append(right, ar[m+1+i])
	}

	i := 0 // Left index
	j := 0 // Right index
	k := l // Current index
	for i < n1 && j < n2 {

		if left[i] <= right[j] {

			// If left is smaller than right, put the left
			// to the current position; increment left index.
			ar[k] = left[i]
			i++
		} else {

			// Otherwise put right to the current position and
			// increment the right index.
			ar[k] = right[j]
			j++
		}

		// Move the current index.
		k++
	}

	// Copy all remaining items in the left temp array.
	// Note that left temp array is already sorted.
	for i < n1 {
		ar[k] = left[i]
		i++
		k++
	}

	// Copy all remaining items in the right temp array.
	// Note that the right temp array is already sorted.
	for j < n2 {
		ar[k] = right[j]
		j++
		k++
	}
}
