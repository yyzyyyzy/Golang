package main

type MyCircularQueue struct {
	values   []int
	len      int
	start    int
	capacity int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		values:   make([]int, k),
		capacity: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.values[(this.start+this.len)%this.capacity] = value
	this.len += 1
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	v := this.Front()
	if v == -1 {
		return false
	}
	this.start = (this.start + 1) % this.capacity
	this.len -= 1
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.values[this.start]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.values[(this.start+this.len-1)%this.capacity]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.len == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.len == this.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
