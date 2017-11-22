package main

import "fmt"

// First let’s define a person, and a method that check if the person knows a friend.

type Person struct {
	Name string
	Friends []string
}

func contains(s []string, c string) bool {
	for _, a := range s {
		if a == c {
			return true
		}
	}

	return false
}

func (p *Person) knows(m *Person) bool {
	return contains(p.Friends, m.Name)
}

// Okay, we have a good-enough `Person` struct.
// Now let’s think about our problem.
//
// Again we’ll try with a small subset of people to visualize our approach.
//
// Assume we have four people: A, B, C, and D
//
// If A knows B, then B may be a celebrity.
// If A does not know B, then A may be a celebrity.
// There is no other option; so we can pick only one candidate out of this first operation.
//
// Let us compare the candidate with C (say the former winner is B)
// if B knows C, then B cannot be a celebrity (because she knows someone)
// C, however, can be our new celebrity candidate.
//
// When we sweep the entire list like this, we will be left with a celebrity candidate.
//
// Next we need to go through the list once more and check if our candidate satisfies the
// celebrity condition:
// 1. Everyone should know him.
// 2. She should know no one.
//
// If our candidate satisfies the condition, then we have a celebrity.
// Otherwise there is no celebrity in the party.
//
// So what’s the complexity of this solution?
//
// Since we do two passes, it can look like O(n).
// However, keep in mind that `contains` function is not constant time,
// it has a loop inside it — so in an average network where everyone knows at
// least half of the party-goers, the time complexity will be more like O(n^2)

func main() {
	joe := Person{Name:"Joe", Friends: []string{"Jill"}}
	jack := Person{Name:"Jack", Friends: []string{"Jill"}}
	jill := Person{Name:"Jill", Friends: []string{"Rosetta", "Joel"}}

	people := []Person{joe, jack, jill}

	// A default person:
	candidate := Person{Name: "", Friends: []string{}}

	for i := 1; i < len(people); i++ {
		a := people[i - 1]
		b := people[i]

		// If the candidate has a name, it means
		// we have selected one; otherwise we are
		// yet to select a candidate.
		candidateSet := candidate.Name != ""

		if candidateSet {
			if candidate.knows(&b) {

				// If candidate knows someone, then she cannot be a celebrity:
				// Then our new candidate becomes b.
				candidate = b
			}

			// Otherwise, she still can be a celebrity.
		} else {
			// This is the initialization, that gets executed only in the first iteration.

			if a.knows(&b) {

				// If a knows b, then she cannot be celebrity; so our celebrity candidate is b.
				candidate = b
			} else {

				// If a does not know b, then she, in deed, can be a celebrity.
				candidate = a
			}
		}
	}

	fmt.Printf("Candidate found: “%s”. Proceeding…\n", candidate.Name)

	// After this first pass we will, for sure, have one `candidate`
	// that may or may not be a celebrity.
	//
	// Let’s iterate over the collection once over to see if she satisfies the conditions
	// of being a celebrity.
	for _, p := range people {

		// Cross check everyone but the candidate:
		if p.Name != candidate.Name {

			// If someone does not know our celebrity, then she is not a celebrity:
			if !p.knows(&candidate) {
				fmt.Println("There is no celebrity here!")
				fmt.Printf("%s does not know %s", p.Name, candidate.Name)

				return
			}

			// If our “celebrity” knows someone, then she’s not that much of
			// a celebrity either:
			if candidate.knows(&p) {
				fmt.Println("There is no celebrity here!")
				fmt.Printf("%s knows %s", candidate.Name, p.Name)

				return
			}
		}
	}

	fmt.Printf("Our star is %s", candidate.Name)
}