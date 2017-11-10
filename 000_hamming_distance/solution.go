package main

import "fmt"

func hammingDistance(x int, y int) int {

	// The bitwise or will flag the changed bits as `1`.
	tmp := x ^ y

	// Result
	dis := 0

	for {

		// Is there a non-zero number to reduce? [1]
		anyNonZeroBits := tmp != 0

		// Everything processed; terminate the loop.
		if !anyNonZeroBits {
			break
		}

		// Shift left, then shift right.
		// If the last bit is `1`, this operation will swallow the last bit.
		// If the last bit is `0`, it will remain `0`.
		tmp1 := (tmp >> 1) << 1
		hasMatch := tmp1 != tmp

		if hasMatch {
			dis++
		}

		// Shift the number until there’s nothing left ([1])
		tmp >>= 1
	}

	return dis
}

func main() {
	fmt.Println(hammingDistance(42, 21))
}
