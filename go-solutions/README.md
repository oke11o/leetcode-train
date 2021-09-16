# All tasks

[Link](https://seanprashad.com/leetcode-patterns/)

# Get random task

[run](./get-random-task/get-random-task.go)

# 0046. Permutations

    Medium

**Tags:** [Backtracking]

[go to solution](./0046-permutations.go)

# 0047. Permutations II

    Medium

**Tags:** [Backtracking]

**Comment:** Need work

[go to solution](./0047-permutations-ii.go)

# 0054. Spiral Matrix

    Medium

**Tags:** [Matrix]

**Links:**

- [go to solution](./0054-spiral-matrix.go)
- [leetcode](https://leetcode.com/problems/spiral-matrix/)

# 0056. Merge Intervals

    Medium

**Tags:** [Array, Sorting]

**Links:**

- [go to solution](./0056-merge-intervals.go)
- [leetcode](https://leetcode.com/problems/merge-intervals/)

# 0070. Climbing Stairs

    Easy

**Tags:** [Math, Dynamic Programming, Memoization]

**Links:**

- [go to solution](./0070-climbing-stairs.go)
- [leetcode](https://leetcode.com/problems/climbing-stairs/)

# 0073. Set Matrix Zeroes

    Medium 

**Tags:** [Matrix]

**Links:**

- [go to solution](./0073-set-matrix-zeroes.go)
- [leetcode](https://leetcode.com/problems/set-matrix-zeroes/)

# 0078. Subsets

    Medium

**Tags:** [Backtracking]

**Links:**

- [go to solution](./0078-subsets.go)
- [leetcode](https://leetcode.com/problems/subsets/)

# 0079. Word Search

    Medium

**Tags:** [Backtracking]

[go to solution](./0079-word-search.go)

# 0090. Subsets II

    Medium

**Tags:** [Backtracking]

[go to solution](./0079-word-search.go)

# 0098. Validate Binary Search Tree

    Medium

**Comment:** -

**Tags:** [DFS]

**Links:**

- [go to solution](./0098-validate-binary-search-tree.go)
- [leetcode](https://leetcode.com/problems/validate-binary-search-tree/)

# 0100. Same Tree

    Easy

**Comment:** -

**Tags:** [DFS]

**Links:**

- [go to solution](./0100-same-tree.go)
- [leetcode](https://leetcode.com/problems/same-tree/)

# 0136. Single Number

    Easy

**Comment:** Just XOR all array items

**Tags:** [Array, Bit Manipulation]

**Links:**

- [go to solution](./0136-single-number.go)
- [leetcode](https://leetcode.com/problems/single-number/)

# 0141. Linked List Cycle

    Easy 

[go to solution](./0141-linked-list-cycle.go)

# 0153. Find Minimum in Rotated Sorted Array

    Medium

**Comment:** Результат так себе. Но работает

**Links:**

- [go to solution](./0153-find-minimum-in-rotated-sorted-array.go)

# 0198. House Robber

    Medium

**Tags:** [Dynamic Programming]

**Comment:** Need understand

**Links:**

- [go to solution](./0198-house-robber.go)
- [leetcode](https://leetcode.com/problems/house-robber/)

# 0202. Happy Number

    Medium

**Tags:** [Hash Table, Math, Two Pointers]

**Comment:**

**Links:**

- [go to solution](./0202-happy-number.go)
- [leetcode](https://leetcode.com/problems/happy-number/)

# 0206. Reverse Linked List

    Easy

[go to solution](./0206-reverse-linked-list.go)

# 0234. Palindrome Linked List

    Easy

[go to solution](./0234-palindrome-linked-list.go)

# 238. Product of Array Except Self

    Medium

*Tags:* [Arrays]

[task](https://leetcode.com/problems/product-of-array-except-self/)

- [go to solution](./238-product-of-array-except-self.go)

# 287. Find the Duplicate Number

    Medium

[go to solution](./287-find-the-duplicate-number.go)

# 0340. Longest Substring with At Most K Distinct Characters

    Medium

**Comment:** Similar as *0904. Fruit Into Baskets*

**Tags:** [Sliding Window]

**Links:**

- [Go solution](./0340-longest-substring-with-at-most-k-distinct-characters.go)
- [Problem](https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/)
- [Problem Lint Code](https://www.lintcode.com/problem/386/)

# 442. Find All Duplicates in an Array

    Medium

[go to solution](./442-find-all-duplicates-in-an-array.go)

# 0448. Find All Numbers Disappeared in an Array. Find All Duplicates in an Array

    Easy

*Tags:* [Arrays]

[go to solution](./0448-find-all-numbers-disappeared-in-an-array.go)

# 0784. Letter Case Permutation

    Medium

[go to solution](./0784-letter-case-permutation.go)

# 0876. Middle of the Linked List

    Easy

[go to solution](./0876-middle-of-the-linked-list.go)

# 0904. Fruit Into Baskets

    Medium

**Comments:** Similar as *0340. Longest Substring with At Most K Distinct Characters*

**Tags:** [Sliding Window]

**Links:**

- [go to solution](./0904-fruit-into-baskets.go)
- [leetcode](https://leetcode.com/problems/fruit-into-baskets/)

# 0917. Reverse Only Letters

    Easy

**Comments:** -

**Tags:** []

**Links:**

- [go to solution](./0917-reverse-only-letters.go)
- [leetcode](https://leetcode.com/problems/reverse-only-letters/)

# 0978. Longest Turbulent Subarray

    Medium

**Comments:**

**Tags:** [Array, Dynamic Programming, Sliding Window]

**Links:**

- [go to solution](./0978-longest-turbulent-subarray.go)
- [leetcode](https://leetcode.com/problems/longest-turbulent-subarray/)

# 1189. Maximum Number of Balloons

    Easy

**Comments:**

**Tags:** [Hash Table, String, Counting]

**Links:**

- [go to solution](./1189-maximum-number-of-balloons.go)
- [leetcode](https://leetcode.com/problems/maximum-number-of-balloons/)

Классная визуализация для Python решения
[pythontutor.com/visualize.html](https://pythontutor.com/visualize.html?mode=display#code=def%20findDuplicates%28nums%29%3A%0A%20%20%20%20result%20%3D%20%5B%5D%0A%20%20%20%20%0A%20%20%20%20for%20num%20in%20nums%3A%0A%20%20%20%20%20%20if%20nums%5Babs%28num%29%20-%201%5D%20%3C%200%3A%0A%20%20%20%20%20%20%20%20result.append%28abs%28num%29%29%0A%20%20%20%20%20%20nums%5Babs%28num%29%20-%201%5D%20*%3D%20-1%0A%20%20%20%20%20%0A%20%20%20%20return%20result%0A%0AfindDuplicates%284,%203,%202,%207,%208,%202,%203,%201%29&cumulative=false&curInstr=0&heapPrimitives=nevernest&mode=display&origin=opt-frontend.js&py=3&rawInputLstJSON=%5B%5D&textReferences=false)