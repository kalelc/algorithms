# Algorithms

## Big O Notation, O(n)

Big O notation is special notation that tells you how fast an algorithm is. Who cares? Well, it turns out that you’ll use other people’s algorithms often and when you do, it’s nice to understand how fast or slow they are. In this section, I’ll explain what Big O notation is and give you a list of the most common running times for algorithms using it.

Big O doesn’t tell you the speed in seconds. _Big O notation lets you compare the number(n) of operations_. It tells you how fast the algorithm grows.

## Some common Big O run times

Here are five Big O run times that you’ll encounter a lot, sorted from fastest to slowest:

- O(1), constant time
- O(log n), also known as log time. Example: Binary search.
- O(n), also known as linear time. Example: Simple search.
- O(n log n). Example: A fast sorting algorithm, like quicksort.
- O(n^2). Example: A slow sorting algorithm, like selection sort.
- O(n!). Example: A really slow algorithm, like the traveling salesperson.


## Arrays and Linked Lists

|         | Array | List |
| Reading |  O(1) | O(n) |
|Insertion|  O(n) | O(1) |
|Deletion |  O(n) | O(1) |


## Algorithms

### Binary Search: O(log n)

It is an algorithm for finding a target value within a sorted array. It works by repeatedly dividing the search interval in half until the target value is found or the search interval is empty.

### Selection Sort: O(n^2)

### Quicksort: O(n log n)

### BubbleSort: O(n^2)

### Graphs

```
       -->        (1) <- Node
      |           / \
Graph |          /   \  <- Edge
      |         (2)  (3)
      |          /    \
       -->      (4)   (5)
```

- breadth-first search (BFS): O(Number of Edges) Finding the shortest path.
