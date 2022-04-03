package _3xx

/**
https://leetcode.com/problems/moving-average-from-data-stream/
346. Moving Average from Data Stream
Easy
*/

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
type MovingAverage struct {
	maxSize int
	queue   []int
	sum     int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		maxSize: size,
		queue:   make([]int, 0, size),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.queue) == this.maxSize {
		this.sum -= this.queue[0]
		this.queue = this.queue[1:]
	}

	this.queue = append(this.queue, val)
	this.sum += val
	return float64(this.sum) / float64(len(this.queue))
}

/**
With cicle queue
*/
type MovingAverage_2 struct {
	maxSize int
	queue   []int
	sum     int
	pointer int
}

func Constructor_2(size int) MovingAverage_2 {
	return MovingAverage_2{
		maxSize: size,
		queue:   make([]int, 0, size),
	}
}

func (this *MovingAverage_2) Next(val int) float64 {
	if len(this.queue) < this.maxSize {
		this.sum += val
		this.queue = append(this.queue, val)
		this.pointer++
		return float64(this.sum) / float64(len(this.queue))
	}
	this.pointer %= this.maxSize
	this.sum -= this.queue[this.pointer]
	this.queue[this.pointer] = val
	this.pointer++
	this.sum += val
	return float64(this.sum) / float64(len(this.queue))
}
