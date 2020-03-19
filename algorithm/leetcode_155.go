package algorithm

/**
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
*/

type MinStack struct {
	stk []int
	min []int
	i   int
}

/** initialize your data structure here. */
func ConstructorMiniStack() MinStack {
	return MinStack{
		stk: []int{},
		min: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stk = append(this.stk, x)
	if this.i > 0 && x > this.min[this.i-1] {
		return
	}
	this.min = append(this.min, x)
	this.i++
}

func (this *MinStack) Pop() {
	if len(this.stk) == 0 {
		return
	}
	t := this.stk[len(this.stk)-1]
	this.stk = this.stk[:len(this.stk)-1]
	if this.min[this.i-1] == t {
		this.i--
		this.min = this.min[:this.i]
	}
}

func (this *MinStack) Top() int {
	if len(this.stk) == 0 {
		return 0
	}
	return this.stk[len(this.stk)-1]
}

func (this *MinStack) GetMin() int {
	if this.i > 0 {
		return this.min[this.i-1]
	}
	return 0
}
