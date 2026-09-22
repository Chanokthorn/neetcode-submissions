/*
1,2,7,4,6,5

2,3,1,4
2
2 3
2 3 1
2 3 1 4

*/
type MinStack struct {
    data []int
    currMin []int
}

func Constructor() MinStack {
    return MinStack{}
}

func (this *MinStack) Push(val int) {
    this.data = append(this.data, val) 
    if len(this.currMin) == 0  {
        this.currMin = append(this.currMin, val)
        return
    }
    top := this.currMin[len(this.currMin) - 1]
    if top < val {
        this.currMin = append(this.currMin, top)
        return
    }
    this.currMin = append(this.currMin, val)
}

func (this *MinStack) Pop() {
    this.data = this.data[:len(this.data) - 1]
    this.currMin = this.currMin[:len(this.currMin) - 1]
}

func (this *MinStack) Top() int {
    return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
    return this.currMin[len(this.currMin)-1]
}
