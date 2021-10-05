# All tasks

[Link](https://seanprashad.com/leetcode-patterns/)

# 0733. Flood Fill

    Easy

**Tags**: [Array, DFS, BFS, Matrix]

**Links**:

- [go solution](./0733-flood-fill.go)
- [leetcode](https://leetcode.com/problems/flood-fill/)

# 0740. Delete and Earn

    Medium

**Tags:** [Array, Hash Table, Dynamic Programming]

**Links:**

- [go solution](./0740-delete-and-earn.go)
- [leetcode](https://leetcode.com/problems/delete-and-earn/)


# 0744. Find Smallest Letter Greater Than Target

    Easy

**Comments:**

**Tags:** [Array, Binary Search]

**Links:**

- [go solution](./0744-find-smallest-letter-greater-than-target.go)
- [leetcode](https://leetcode.com/problems/find-smallest-letter-greater-than-target/)

# 0746. Min Cost Climbing Stairs

    Easy

**Comments:**

**Tags:** [Array, Dynamic Programming]

**Links:**

- [go solution](./0744-min-cost-climbing-stairs.go)
- [leetcode](https://leetcode.com/problems/min-cost-climbing-stairs/)

# 0782. Transform to Chessboard

    Hard

**Tags:** [Array, Math, Bit Manipulation, Matrix]

**Links:**

- [go solution](./0782-transform-to-chessboard.go)
- [leetcode](https://leetcode.com/problems/transform-to-chessboard/)

Классная визуализация для Python решения
[pythontutor.com/visualize.html](https://pythontutor.com/visualize.html?mode=display#code=def%20findDuplicates%28nums%29%3A%0A%20%20%20%20result%20%3D%20%5B%5D%0A%20%20%20%20%0A%20%20%20%20for%20num%20in%20nums%3A%0A%20%20%20%20%20%20if%20nums%5Babs%28num%29%20-%201%5D%20%3C%200%3A%0A%20%20%20%20%20%20%20%20result.append%28abs%28num%29%29%0A%20%20%20%20%20%20nums%5Babs%28num%29%20-%201%5D%20*%3D%20-1%0A%20%20%20%20%20%0A%20%20%20%20return%20result%0A%0AfindDuplicates%284,%203,%202,%207,%208,%202,%203,%201%29&cumulative=false&curInstr=0&heapPrimitives=nevernest&mode=display&origin=opt-frontend.js&py=3&rawInputLstJSON=%5B%5D&textReferences=false)