package stack

type MyStack struct {
	queue1 []int
}

func StackConstructor() MyStack {
	return MyStack{queue1: []int{}}
}

func (this *MyStack) Push(x int) {

	this.queue1 = append(this.queue1, x)

	lastIndex := len(this.queue1) - 1
	index := 0
	for index != lastIndex {
		front := this.queue1[0]
		this.queue1 = this.queue1[1:]
		this.queue1 = append(this.queue1, front)
		index++
	}

}

func (this *MyStack) Pop() int {
	elem := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return elem
}

func (this *MyStack) Top() int {
	return this.queue1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}
