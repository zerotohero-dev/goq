## Definitions

* A [binary tree][btree] is a data structure of connected nodes, where each node can have at most two child nodes.
* The **root node** of the tree is the node that does not have any parent nodes.
* A tree is typically identified by its root node.

## Question

You have two binary trees. Imagine moving one of them on top the other. When that happens some nodes of the top tree will overlap with the bottom tree, while some will not.

Write a function to merge these two binary trees into a third tree.

## Assumptions

* A `merge` operation is defined as adding the values of the nodes when they overlap, otherwise using whichever node that is not `NULL` as the result.
* Assume the trees contain positive integer values.
* Assume there is no numeric overflow after the merge operation.
* Start the merging process from the root of the trees.

## Example

```text
First Tree
      1
     / \
    3   2
   /
  5
Second Tree                  
      1
     / \
    6   4
     \   \
      4   7
Merge Result       
      3
     / \
    9   6
   / \   \ 
  5   4   7
```

[btree]: https://en.wikipedia.org/wiki/Binary_tree "Binary Tree"
