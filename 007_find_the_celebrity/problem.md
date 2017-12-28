## Question

There is a party that N people are attending.

In this party, there **may or may not be** a celebrity.

* A celebrity is a person where everybody knows.
* A celebrity knows no one (*other than himself*) in the party.

Devise an efficient algorithm to find this celebrity.

## Assumptions

* You are only allowed to talk to individual persons.
* You can only ask “Do you know him/her” questions to a person (*i.e. “Hey John,
  do you know Jill?*)
* People will be truthful in their answers, and they have perfect memories; so
  if they know someone, they will know for sure.
* Keep in mind that there might not be any celebrity in the party at all.

## Example

```text
Given --> depicts a “knows” relationship

Of
    Alice
    Bob
    Charlie
    David

If
    Alice --> Bob
    Charlie --> Bob
    David --> Bob
Then Bob is the celebrity.

If
    Alice --> Charlie
    Bob --> Charlie
    Charlie --> Alice
    Bob --> David
Then there is no celebrity in the party.
```
