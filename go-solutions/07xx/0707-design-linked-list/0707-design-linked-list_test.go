package _707_design_linked_list

// https://leetcode.com/problems/design-linked-list/
// 707. Design Linked List
type Node struct {
	val  int
	next *Node
}

func NewNode(val int) *Node {
	return &Node{
		val: val,
	}
}

type MyLinkedList struct {
	root *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.root
	for index > 0 && cur != nil {
		cur = cur.next
		index--
	}
	if index > 0 || cur == nil {
		return -1
	}
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	n := &Node{val: val}
	n.next = this.root
	this.root = n
}

func (this *MyLinkedList) AddAtTail(val int) {
	n := &Node{val: val}
	if this.root == nil {
		this.root = n
		return
	}
	cur := this.root
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = n
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	cur := this.root
	for index > 1 && cur != nil {
		index--
		cur = cur.next
	}
	if index == 1 && cur != nil {
		n := &Node{val: val}
		n.next = cur.next
		cur.next = n
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.root = this.root.next
		return
	}
	cur := this.root
	for index > 1 && cur != nil {
		index--
		cur = cur.next
	}
	if index == 1 && cur != nil && cur.next != nil {
		cur.next = cur.next.next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

/**
Test cases
*/
// ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
// [[],[1],[3],[1,2],[1],[1],[1]]
// [null,null,null,null,2,null,3]

// ["MyLinkedList","addAtHead","get","addAtHead","addAtHead","deleteAtIndex","addAtHead","get","get","get","addAtHead","deleteAtIndex"]
// [[],[4],[1],[1],[5],[3],[7],[3],[3],[3],[1],[4]]
// [null,null,-1,null,null,null,null,4,4,4,null,null]

// ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get","get","deleteAtIndex","deleteAtIndex","get","deleteAtIndex","get"]
// [[],[1],[3],[1,2],[1],[1],[1],[3],[3],[0],[0],[0],[0]]
// [null,null,null,null,2,null,3,-1,null,null,3,null,-1]

// ["MyLinkedList","addAtHead","addAtTail","addAtIndex"]
// [[],[1],[3],[3,2]]
// [null,null,null,null]
