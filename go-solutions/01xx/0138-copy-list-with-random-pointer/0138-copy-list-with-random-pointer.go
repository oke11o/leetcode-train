package _138_copy_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 0138. Copy List with Random Pointer
// https://leetcode.com/problems/copy-list-with-random-pointer/
// https://www.educative.io/blog/crack-amazon-coding-interview-questions
// Amazon: 4. Copy linked list with arbitrary pointer
func copyRandomList(head *Node) *Node {
	m := map[*Node]*Node{}

	result := &Node{}
	current := result
	for head != nil {
		newNode := &Node{Val: head.Val, Random: head.Random}
		m[head] = newNode
		current.Next = newNode
		current = current.Next
		head = head.Next
	}

	current = result.Next
	for current != nil {
		newRandom := m[current.Random]
		if newRandom != nil {
			current.Random = newRandom
		} else {
			current.Random = nil
		}
		current = current.Next
	}

	return result.Next
}
