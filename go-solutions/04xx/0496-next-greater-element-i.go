package _4xx

/**
https://leetcode.com/problems/next-greater-element-i/
496. Next Greater Element I
Easy

Тут просто создаем мапу ответов. Но как ее создать?
Используем monotonous stack. Убывающий, так как нам нужно следующие большее значение.
Если наш элемент меньше, добавляем на стек, для последующей обработки.
Если больше - это наш парень - следующий больший элемент. Наверху стека лежит - как раз для кого он является большим.
У нас есть вершина стека и текущий элемент. Вот с ними можно, что-то делать.
Когда один элемент снимем со стека, то обрабатываем. И пробуем снять второй.

Важно понимать, что на стек мы кладем для следующий обработки.
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, 0, len(nums1))
	m := make(map[int]int, len(nums2))
	stack := []int{}
	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			m[stack[len(stack)-1]] = v
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, v)
	}
	for len(stack) > 0 {
		m[stack[len(stack)-1]] = -1
		stack = stack[0 : len(stack)-1]
	}
	for _, n := range nums1 {
		result = append(result, m[n])
	}

	return result
}
