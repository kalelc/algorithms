# Algorithms In Golang

## Big O Notation, O(n)

Big O notation is special notation that tells you how fast an algorithm is. Who cares? Well, it turns out that you’ll use other people’s algorithms often and when you do, it’s nice to understand how fast or slow they are. In this section, I’ll explain what Big O notation is and give you a list of the most common running times for algorithms using it.

Big O doesn’t tell you the speed in seconds. _Big O notation lets you compare the number(n) of operations_. It tells you how fast the algorithm grows.

## Some common Big O run times

Here are five Big O run times that you’ll encounter a lot, sorted from fastest to slowest:

- O(log n), also known as log time. Example: Binary search.
- O(n), also known as linear time. Example: Simple search.
- O(n log n). Example: A fast sorting algorithm, like quicksort.
- O(n^2). Example: A slow sorting algorithm, like selection sort.
- O(n!). Example: A really slow algorithm, like the traveling salesperson.

## Algorithms

- Binary Search: O(log n)
- Selection Sort: O(n^2)
- Quicksort: O(n log n)
- BubbleSort: O(n^2)

### Graphs

```
       -->          (1) <- Node
      |           / \
Graph |          /   \  <- Edge
      |         (2)  (3)
      |          /    \
       -->      (4)   (5)
```

- breadth-first search (BFS): O(Number of Edges) Finding the shortest path.
