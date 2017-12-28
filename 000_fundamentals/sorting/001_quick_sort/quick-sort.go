package main

// Add some details about the algorithm.

func partition(ar []int, low, high int) int {
	i := low - 1

	for j := low; j <= high-1; j++ {
		if ar[j] <= ar[high] {
			i++
			ar[i], ar[j] = ar[j], ar[i]
		}
	}
	ar[i+1], ar[high] = ar[high], ar[i+1]

	return i + 1
}

func quickSort(ar []int, low, high int) {
	if low >= high {
		return
	}

	pi := partition(ar, low, high)

	quickSort(ar, low, pi-1)
	quickSort(ar, pi+1, high)
}
