## Definitions

N/A

## Question

A robot is programmed to take an input string and do a sequence of movements by reading the characters in order.
(i.e., the movement sequence is a string, and each move is a character.)

The valid characters that the robot understands are “U”: Up, “D”: Down, “L”: Left, and “R”: Right.

> For example given the sequence “UDL” the robot sequentially moves 1 unit up, 1 unit down, and 1 unit left.

Given an input sequence string, write a function that returns a boolean; `true` if the sequence of movement is circular, and `false` otherwise.

## Assumptions

* Assume there is a robot at the origin of an Cartesian coordinate plane.
* You can assume that the robot skips any character that it does not understand.
* A “*circular movement*” can be defined by any sequence of characters that takes the robot to its original starting position.
    * “*Circular*”ly moving the robot does not necessarily create a geometric circle; it only means that the robot came back to its starting position at the end of the sequence.

## Example

```text
"UD" -> true   — The robot moves back to its original start position.
"LLU" -> false — The robot is away from its original start position. 
```
