/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.cn/problems/min-stack/description/
 *
 * algorithms
 * Medium (60.00%)
 * Likes:    1786
 * Dislikes: 0
 * Total Accepted:    604.2K
 * Total Submissions: 1M
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 *
 * Implement the MinStack class:
 *
 *
 * MinStack() initializes the stack object.
 * void push(int val) pushes the element val onto the stack.
 * void pop() removes the element on the top of the stack.
 * int top() gets the top element of the stack.
 * int getMin() retrieves the minimum element in the stack.
 *
 *
 * You must implement a solution with O(1) time complexity for each
 * function.
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 *
 * Output
 * [null,null,null,null,-3,null,0,-2]
 *
 * Explanation
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin(); // return -3
 * minStack.pop();
 * minStack.top();    // return 0
 * minStack.getMin(); // return -2
 *
 *
 *
 * Constraints:
 *
 *
 * -2^31 <= val <= 2^31 - 1
 * Methods pop, top and getMin operations will always be called on non-empty
 * stacks.
 * At most 3 * 10^4 calls will be made to push, pop, top, and getMin.
 *
 *
*/

// @lc code=start
type MinStack struct {
	data   []int
	minVal []int
}

func Constructor() MinStack {
	res := MinStack{}
	res.data = []int{}
	res.minVal = []int{}
	return res
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if len(this.minVal) == 0 || this.minVal[len(this.minVal)-1] > val {
		this.minVal = append(this.minVal, val)
	} else {
		this.minVal = append(this.minVal, this.minVal[len(this.minVal)-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.data) > 0 {
		this.data = this.data[:len(this.data)-1]
		this.minVal = this.minVal[:len(this.minVal)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if len(this.minVal) > 0 {
		return this.minVal[len(this.minVal)-1]
	}
	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

