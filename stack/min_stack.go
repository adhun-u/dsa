package stack

type MinStack struct {
	values []int
	mins   []int
	min    *int
	top    int
}

func Constructor() MinStack {

	return MinStack{
		values: []int{},
		min:    nil,
		top:    0,
		mins:   []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.values = append(this.values, val)

	if this.min == nil {
		this.min = &val
	} else if val < *(this.min) {
		this.mins = append(this.mins, *(this.min))
		this.min = &val
	} else if val == *(this.min) {
		this.mins = append(this.mins, val)
	}

	this.top = val
}

func (this *MinStack) Pop() {

	if this.top == *(this.min) {

		if len(this.mins) != 0 {
			this.min = &this.mins[len(this.mins)-1]
			this.mins = this.mins[:len(this.mins)-1]
		} else {
			this.min = nil
		}
	}
	this.values = this.values[:len(this.values)-1]

	if len(this.values) != 0 {
		this.top = this.values[len(this.values)-1]
	}
}

func (this *MinStack) Top() int {
	return this.top
}

func (this *MinStack) GetMin() int {
	return *(this.min)
}
