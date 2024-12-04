# 🎄 Advent of Code Solutions

My solutions for [Advent of Code 2024](https://adventofcode.com/), a series of daily programming puzzles released during December. 


## Progress

- [x] Day 1 [Solution](https://github.com/Crowley723/advent-2024/blob/main/day01/main.go)
- [x] Day 2 [Solution](https://github.com/Crowley723/advent-2024/blob/main/day02/main.go)
- [ ] Day 3
- [ ] Day 4

[...]


## Journal
### Day 1
#### Part 1
- Load and Parse the input file into two different integer arrays.
- Sort the two arrays.
- Calculate the difference between each pair of values, from smallest to largest
- Sum the differences.
#### Part 2
- Load, Parse, and Sort (same methods from above)
- Find the number of occurrences of each value in the right array.
- Multiply each number in the left array by the occurrences in the right array.
- Sum

---
### Day 2
#### Part 1
- Load and Parse using a different parsing function from Day one since the input data is formatted differently.
- Write test function to validate a row's trend as well as make sure that the jumps between different numbers are within the specified distance.
- Test all rows, count rows that meet criteria.
#### Part 2
- Same as above.
- Now we also check every row minus each of the elements. [1, 2, 3] -> [2, 3], [1,3], [1,2] to see if one of the sub-slices works.
- Count rows that meet new criteria.

---