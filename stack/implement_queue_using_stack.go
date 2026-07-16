package stack

type MyQueue struct {
	in  []int
	out []int
}

func QueueConstructor() MyQueue {
	return MyQueue{in: []int{}, out: []int{}}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) InsertToOut() {

	if len(this.out) == 0 {
		for len(this.in) != 0 {
			last := this.in[len(this.in)-1]
			this.out = append(this.out, last)
			this.in = this.in[:len(this.in)-1]
		}
	}

}
func (this *MyQueue) Pop() int {
	this.InsertToOut()
	lastElem := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return lastElem
}

func (this *MyQueue) Peek() int {
	this.InsertToOut()
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {

	return len(this.out) == 0 && len(this.in) == 0
}
