package main

import "fmt"

func primeSieve(n int) []int {
	primes := make([]bool, n+1)
	for i := range primes {
		primes[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if primes[p] {
			for i := p * 2; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	result := make([]int, 0)

	for p := 2; p <= n; p++ {
		if primes[p] {
			result = append(result, p)
		}
	}

	return result
}

func main() {
	fmt.Println(primeSieve(1000))
}
