# All tasks

[Link](https://seanprashad.com/leetcode-patterns/)

# Get random task

[run](./get-random-task/get-random-task.go)

# 0303. Range Sum Query - Immutable

    Easy

**Comment:**

**Tags:** [Array, Design, Prefix Sum]

**Links:**

- [Go solution](./0303-range-sum-query-immutable.go)
- [Problem](https://leetcode.com/problems/range-sum-query-immutable/)

# 0340. Longest Substring with At Most K Distinct Characters

    Medium

**Comment:** Similar as *0904. Fruit Into Baskets*

**Tags:** [Sliding Window]

**Links:**

- [Go solution](./0340-longest-substring-with-at-most-k-distinct-characters.go)
- [Problem](https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/)
- [Problem Lint Code](https://www.lintcode.com/problem/386/)

# 0350. Intersection of Two Arrays II

    Easy

**Comment:**

**Tags:** [Array, Hash Table, Two Pointers, Binary Search, Sorting]

**Links:**

- [Go solution](./0350-intersection-of-two-arrays-ii.go)
- [Problem](https://leetcode.com/problems/intersection-of-two-arrays-ii/)

# 0704. Binary Search

    Easy

**Tags:** [Array, Binary Search]

**Links:**

- [go solution](./0704-binary-search.go)
- [leetcode](https://leetcode.com/problems/binary-search/)

# 0784. Letter Case Permutation

    Medium

[go to solution](./0784-letter-case-permutation.go)

# 0876. Middle of the Linked List

    Easy

[go to solution](./0876-middle-of-the-linked-list.go)

# 0904. Fruit Into Baskets

    Medium

**Comments:** Similar as *0340. Longest Substring with At Most K Distinct Characters*

**Tags:** [Array, Hash Table, Sliding Window]

**Links:**

- [go to solution](./0904-fruit-into-baskets.go)
- [leetcode](https://leetcode.com/problems/fruit-into-baskets/)

# 0917. Reverse Only Letters

    Easy

**Comments:** -

**Tags:** [Two Pointers, String]

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

# Алгоритм подготовки



1. Читаем условие задачи, ни в коем случае не пытаемся придумать решение до того, как условие прочитано до конца. Это важно!!! Мозг пытается найти похожую задачу, решение которой он знает, и выдать за требуемое.
2. Пытаемся придумать уточняющие вопросы. Пример 1: есть задача, в которой нужно как-то трансформировать строку. Что спрашивать? — Какие символы могут быть в строке — ASCII или Unicode? Могут ли рядом стоять несколько пробелов? Могут ли быть пробелы в начале или конце строки? Есть ли спецсимволы типа -,.^/ ? Есть ли разница для анализа между большой и маленькой буквами? Насколько длинная входная строка? Помещается ли она в память машины? Пример 2: есть массив из Integer, в нем нужно что-то найти. Вопросы: есть ли повторяющиеся элементы? Есть ли отрицательные числа? Что если в результате подсчета мы получим больше, чем Integer.MAX_VALUE?
3. Рисуем примеры, лучше парочку — один классический, второй с corner cases. После этого мы +/- должны быть уверены в том, что задачу мы поняли правильно.
4. Придумываем решение «в лоб» и оцениваем его сложность. Сложность решения нужно уметь определить всегда.
5. Придумываем более оптимальное решение, оцениваем его сложность.
6. Разрабатываем API решения — какие будут методы (приватные и публичные).
7. Пишем код в тетрадке.
8. Дебажим код по тетрадке на новом примере. Не нужно брать один из примеров, который мы рисовали в начале. В этом случае очень высока вероятность, что мы написали решение именно для этого случая, а не для всех возможных. Лучше взять новый пример с corner-кейсом, такой, чтобы потенциально мог решение сломать.
9. Перебиваем код в любимую IDE, при этом не смотрим в бумажку. Таким образом, мы повторяем решение два раза.
10. Копируем код из IDE в LeetCode и запускаем. В случае идеального выполнения должно cработать правильно с первого раза. У меня такое получалось в 10% случаев.