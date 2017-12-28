## Definitions

[The Sieve of Erathosthenes][sieve] is a very old algorithm for identifying prime numbers.

In the normaBeginning at 2, the algorithm iterates upward: For the current number,
if the number has not been marked, we identify it as a prime, and then mark all
multiples of that number up to our target number.

## Question

Implement The Sieve of Erathosthenes.

## Assumptions

* Assume that memory is not a problem;
  you have enough space to store temporary lookup info.
* You are **not** expected to optimize the algorithm, since anyone
  who knows how to read Wikipedia can come up with an optimized solution.
* This is **not** a trick question; it is a simple algorithm implementation.

## Example

So to find all primes up to 12:

* We would start at 2, note 2 as prime, and mark 4, 6, 8, 10 and 12 as non-prime
* Move to 3, note 3 as prime, and mark 6, 9, 12 as non-prime
* move to 4, see that it is marked, and skip it
* Found prime numbers so far: 2, 3, 5, 7, 11

[sieve]: http://example.com/
