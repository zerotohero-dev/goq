package main

// Add some details about the algorithm.

func merge(ar []int, l int, m int, r int) {
	n1 := m - l + 1
	n2 := r - m

	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < n1; i++ {
		left = append(left, ar[l+i])
	}

	for i := 0; i < n2; i++ {
		right = append(right, ar[m+1+i])
	}

	i := 0
	j := 0
	k := l
	for i < n1 && j < n2 {
		if left[i] <= right[j] {
			ar[k] = left[i]
			i++
		} else {
			ar[k] = right[j]
			j++
		}
		k++
	}

	for i < n1 {
		ar[k] = left[i]
		i++
		k++
	}

	for j < n2 {
		ar[k] = right[j]
		j++
		k++
	}
}

func mergeSort(ar []int, l int, r int) {
	if l >= r {
		return
	}

	m := (l + (r - 1)) / 2

	mergeSort(ar, l, m)
	mergeSort(ar, m+1, r)

	merge(ar, l, m, r)
}
