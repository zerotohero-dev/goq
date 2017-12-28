package main

import "fmt"

func isCircular(moves string) bool {

	// Count all the moves
	left := 0
	right := 0
	up := 0
	down := 0
	for _, m := range moves {
		switch m {
		case 'L':
			left++
		case 'R':
			right++
		case 'U':
			up++
		case 'D':
			down++
		}
	}

	// You can come back to your original point only when all your movement
	// vectors cancel out. That means your left movement should equal your
	// right movement, and your top movement should equal your bottom movement.
	return left == right && up == down
}

func main() {
	fmt.Println(isCircular("TRUFFLEBOTTOMLEFT"))
}