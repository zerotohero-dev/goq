package main

// The [Quicksort algorithm](https://en.wikipedia.org/wiki/Quicksort) is
// about finding a pivot and sorting the collection around that pivot.
//
// When implemented well, on average, it can be notably faster than its
// competitors (such as Mergesort and Heapsort), especially when the data to
// sort is randomized enough.

func quickSort(ar []int, low, high int) {

	// If there is a single item, it is already sorted;
	// exit recursion.
	if low >= high {
		return
	}

	// Find the partition index.
	// `low`, and `high` are lower and higher indices of
	// the subarray that we are sorting.
	pi := partition(ar, low, high)

	// Sort the left half.
	quickSort(ar, low, pi-1)

	// Sort the right half.
	quickSort(ar, pi+1, high)
}

// Again, the clever part is the “partition” logic, since
// the actual `quickSort` function is straightforward enough.
//
// There are many different ways to partition an array (including
// selecting the pivot item to partition around totally at random
// — which, surprisingly, works well)
//
// In this current implementation, we are using the last item of the
// array as the pivot.
func partition(ar []int, low, high int) int {

	// i points to the end of the right partiton.
	i := low - 1

	pivot := ar[high]

	// Since last item is the pivot, we iterate from the lowest index
	// the the next-to-last item.
	for j := low; j <= high-1; j++ {

		// If the current item is greater than the pivot, it can stay where it is.
		// If it less than the pivot then we swap it with the first item of the
		// left partition, which makes it the last item of the right partition.
		//
		// i points to the end of the left partition.
		//
		//     j
		//     |
		// L L L S S S P
		//
		//         j
		//         v
		// L L L S S S P
		// *     *        : * are swapped.
		//
		//         j
		//         v
		// S L L L S S P
		// i
		//
		// S S L L L S P
		//   i
		//
		// S S S L L L P
		//     i
		//
		// When the loop ends `i+1` will point to the first element
		// of the right partition. — Swapping the pivot with it will
		// nicely move the pivot to in its place where the items to the
		// left of the pivot as less than the pivot; and the items to
		// the right of the pivot are greater than the pivot.
		//
		// S S S P L L L
		//       |
		//
		if ar[j] <= pivot {
			// move i to the first item of the right partition.
			i++

			// swap items, so that ar[j] is now at the end of the left
			// partition, and i points to the end of the left partition.
			ar[i], ar[j] = ar[j], ar[i]
		}
	}

	// ar[i+1] is the first item in the right partition.
	// Swapping it with the pivot will move the pivot to its proper place.
	ar[i+1], ar[high] = ar[high], ar[i+1]

	// Return the pivot’s index.
	return i + 1
}
