package week01

// https://leetcode-cn.com/problems/design-circular-deque/

type MyCircularDeque struct {
	q []int
	length int
	capacity int
}

func Constructor(k int) MyCircularDeque {
	q := make([]int, 0)
	return MyCircularDeque{
		q: q,
		length: 0,
		capacity: k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.length >= this.capacity{
		return false
	}
	this.q = append([]int{value}, this.q...)
	this.length++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.length >= this.capacity{
		return false
	}
	this.q = append(this.q, value)
	this.length++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.length < 1{
		return false
	}
	this.q = this.q[1:]
	this.length--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.length < 1{
		return false
	}
	this.q = this.q[:this.length-1]
	this.length--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.length < 1{
		return -1
	}
	elem := this.q[0]
	// this.q = this.q[1:]
	// this.length--
	return elem
}

func (this *MyCircularDeque) GetRear() int {
	if this.length < 1{
		return -1
	}
	elem := this.q[this.length-1]
	// this.q = this.q[:this.length-1]
	// this.length--
	return elem
}

func (this *MyCircularDeque) IsEmpty() bool {
	if this.length < 1{
		return true
	}
	return false
}

func (this *MyCircularDeque) IsFull() bool {
	if this.length >= this.capacity{
		return true
	}
	return false
}

