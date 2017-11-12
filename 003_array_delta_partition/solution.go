package main

import (
	"sort"
	"fmt"
)

// Writing the code is easy, you can even intuitively think about a solution.
// Formalizing the solution, however, takes some afterthought.
//
// Let’s put all the numbers on a line in increasing order.
// I will take numbers 1 to 10 to simplify the discussion; though these can be any
// numbers.
//
// <--1--2--3--4--5--6--7--8--9--10-->
//
// Then picking a tuple is equivalent to creating a line segment.
//
// Let us pick five line segments at random.
//
//          |-----------|
//       |-----|     |-----------|
//    |-----------|        |--|
// <--1--2--3--4--5--6--7--8--9--10-->  (case 1)
//    a1 b1 c1 b2 a2 d1 c2 e1 e2 d2
//
// Well that looks pretty random. Since we are seeking an extremum (i.e. maximum, or minimum)
// to our solution, we are (most probably) looking for a more orderly orientation.
//
// So let’s reorder the segments:
//
//    |--|  |--|  |--|  |--|  |--|
// <--1--2--3--4--5--6--7--8--9--10-->  (case 2)
//    a1 a2 b1 b2 c1 c2 d1 d2 e1 e2
//
// That looks pretty much “ordered”.
//
// Let us compare the sums in either case:
//
// case 1: min({1, 5}) + min({2, 4}) + min({3, 7}) + min({6, 10}) + min({8, 9}) = 20
// case 2: min({1, 2}) + min({3, 4}) + min({5, 6}) + min({7, 8}) + min({9, 10}) = 25
//
// Clearly, case 2 gives the higher sum, but why?
//
// Well, since we are reordering line segments, it might be related to the length of those
// segments.
//
// Lets sum the lengths of them:
//
// TotalLength = (a2-a1) + (b2-b1) + (c2-c1) + (d2-d1) + (e2-e1)
//             = (a2+b2+c2+d2+e2) - (a1+b1+c1+d1+e1)
//
// We are asked to maximize `a1+b1+c1+d1+e1` (the second part of the equation)
// That, by association, means that we are asked to *minimize* the total length of
// the line segments. [0]
//
// So how do we minimize the total length of those line segment?
// By making sure that they do **not** overlap.
//
// (<aside>
//  Tangentially-related: See also “interval scheduling” (https://en.wikipedia.org/wiki/Interval_scheduling)
//  that has some nice challenges about things like how to merge overlapping intervals
// </aside>) 
//
// Why?
//
// Assume we have at least two overlapping segments as in `case 1`:
//
//          |-----------|
//       |-----|
//
// If you swap the overlapping ends of those segments so that they don’t overlap anymore,
// you will decrease the lengths of *both* of those segments and hence decreasing `TotalLength`:
//
//             |--------|
//       |--|
//
// Therefore, if there are overlapping segments, then `TotalLength` is NOT minimized
// (since the swap makes the sum less already).
//
// Therefore to minimize `TotalLength`, it is necessary (but maybe not sufficient) that
// no segments shall overlap. [1]
//
// Also, any arrangement other than `case 2` will result in at least two segments to overlap. [2]
//
// So from [1] and [2], it follows that the only way to minimize the sum of the total length of
// the line segments is to arrange them as in `case 2`.
//
// Minimizing the `TotalLength` means maximizing `sum(min({ai,bi})) : i from 1 to n` (from [0]).
//
// Therefore, solving this problem is equivalent to
//
// * Sorting the array in increasing order.
// * Selecting the odd-indexed items.
// * Summing them up.
//
func sum(nums []int) int {

	// Note that sorting is in-place; it changes the given slice,
	// it does not return a new one.
	sort.Ints(nums)

	result := 0

	for index, number := range nums {
		if index%2 == 0 {
			result += number
		}
	}

	return result
}

func main() {
	fmt.Println(sum([]int{11, 22, -22, 356, 192, 25, 54}));
}
