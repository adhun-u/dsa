package stack

type MyStack struct {
	queue []int
}

func StackConstructor() MyStack {
	return MyStack{queue: []int{}}
}

func (this *MyStack) Push(x int) {

	this.queue = append(this.queue, x)
	lastIndex := len(this.queue) - 1
	currIndex := 0
	for currIndex < lastIndex {
		front := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, front)
		currIndex++
	}
}

func (this *MyStack) Pop() int {
	elem := this.queue[0]
	this.queue = this.queue[1:]
	return elem
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
