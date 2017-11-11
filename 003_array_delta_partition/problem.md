## Definitions

* [Array data type](array)
* [Tuple data type](tuple)

## Question

We have an array `A` of 2n integers.

Assume you arbitrarily partition this array in to `n` tuples of `{ai,bi}`.

Let the sum `S` be `sum(min({ai,bi})) : i from 1 to n`.

Write a function that takes this array `A` as an input argument, then computes and returns the maximum of such sums.

Or more formally, find… 

```text
max(
    sum(
        min( {ai,bi} )
    ) : i from 1 to n
)
```

## Assumptions

* You can assume that the array has an even number of elements.
* Assume the array has 2 to 10,000 items.
* These are integers, they can be negative too.

## Example

```text
given 
A: [1,5,2,11] 

The optimum tuple (of all possible tuples) is: 
T: {11,5}, {1,2} 

because
min(11, 5) + min(1, 2) = 7

where, 7 is the maximum of all possible of such sums.  
```

[Array]: https://en.wikipedia.org/wiki/Array "Array"
[Tuple]: https://en.wikipedia.org/wiki/Tuple "Tuple"