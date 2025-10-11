package no155

type MinStack struct {
	List []int
	Min  []int
}

func Constructor() MinStack {
	return MinStack{
		List: make([]int, 0),
		Min:  make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.List = append(this.List, val)

	if len(this.Min) == 0 || this.Min[len(this.Min)-1] >= val {
		this.Min = append(this.Min, val)
	}
}

func (this *MinStack) Pop() {
	val := this.List[len(this.List)-1]

	this.List = this.List[:len(this.List)-1]

	if val == this.Min[len(this.Min)-1] {
		this.Min = this.Min[:len(this.Min)-1]
	}
}

func (this *MinStack) Top() int {
	return this.List[len(this.List)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
