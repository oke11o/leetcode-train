# 0054. Spiral Matrix

    Medium

[go to solution](./0054-spiral-matrix.go)

# 073. Set Matrix Zeroes

    Medium 

[go to solution](./073-set-matrix-zeroes.go)

## 0079. Word Search

    Medium

[Backtracking]

[go to solution](./0079-word-search.go)

# 141. Linked List Cycle

    Easy 

[go to solution](./141-linked-list-cycle.go)

# 153. Find Minimum in Rotated Sorted Array

    Medium

[go to solution](./153-find-minimum-in-rotated-sorted-array.go)

Результат так себе. Но работает

# 238. Product of Array Except Self

    Medium

[go to solution](./238-product-of-array-except-self.go)

# 287. Find the Duplicate Number

    Medium

[go to solution](./287-find-the-duplicate-number.go)

# 442. Find All Duplicates in an Array

    Medium

[go to solution](./442-find-all-duplicates-in-an-array.go)

Классная визуализация для Python решения
[pythontutor.com/visualize.html](https://pythontutor.com/visualize.html?mode=display#code=def%20findDuplicates%28nums%29%3A%0A%20%20%20%20result%20%3D%20%5B%5D%0A%20%20%20%20%0A%20%20%20%20for%20num%20in%20nums%3A%0A%20%20%20%20%20%20if%20nums%5Babs%28num%29%20-%201%5D%20%3C%200%3A%0A%20%20%20%20%20%20%20%20result.append%28abs%28num%29%29%0A%20%20%20%20%20%20nums%5Babs%28num%29%20-%201%5D%20*%3D%20-1%0A%20%20%20%20%20%0A%20%20%20%20return%20result%0A%0AfindDuplicates%284,%203,%202,%207,%208,%202,%203,%201%29&cumulative=false&curInstr=0&heapPrimitives=nevernest&mode=display&origin=opt-frontend.js&py=3&rawInputLstJSON=%5B%5D&textReferences=false)