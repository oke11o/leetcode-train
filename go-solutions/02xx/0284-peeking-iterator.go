package _2xx

/*   Below is the interface for Iterator, which is already defined for you.
 *
 */
type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	return false // Returns true if the iteration has more elements.
}

func (this *Iterator) next() int {
	return 0 // Returns the next element in the iteration.
}

/**
https://leetcode.com/problems/peeking-iterator/
284. Peeking Iterator
Medium
*/
type PeekingIterator struct {
	iter      *Iterator
	peeked    int
	hasPeeked bool
}

func Constructor_(iter *Iterator) *PeekingIterator {
	p := &PeekingIterator{iter: iter}
	if iter.hasNext() {
		p.peeked = iter.next()
		p.hasPeeked = true
	}

	return p
}

func (this *PeekingIterator) hasNext() bool {
	return this.hasPeeked || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	result := this.peeked
	if this.iter.hasNext() {
		this.peeked = this.iter.next()
	} else if this.hasPeeked {
		this.hasPeeked = false
	} else {
		this.peeked = 0
	}
	return result
}

func (this *PeekingIterator) peek() int {
	return this.peeked
}
