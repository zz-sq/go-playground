/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 *
 * https://leetcode.cn/problems/evaluate-reverse-polish-notation/description/
 *
 * algorithms
 * Medium (54.67%)
 * Likes:    917
 * Dislikes: 0
 * Total Accepted:    388.5K
 * Total Submissions: 710.5K
 * Testcase Example:  '["2","1","+","3","*"]'
 *
 * You are given an array of strings tokens that represents an arithmetic
 * expression in a Reverse Polish Notation.
 *
 * Evaluate the expression. Return an integer that represents the value of the
 * expression.
 *
 * Note that:
 *
 *
 * The valid operators are '+', '-', '*', and '/'.
 * Each operand may be an integer or another expression.
 * The division between two integers always truncates toward zero.
 * There will not be any division by zero.
 * The input represents a valid arithmetic expression in a reverse polish
 * notation.
 * The answer and all the intermediate calculations can be represented in a
 * 32-bit integer.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: tokens = ["2","1","+","3","*"]
 * Output: 9
 * Explanation: ((2 + 1) * 3) = 9
 *
 *
 * Example 2:
 *
 *
 * Input: tokens = ["4","13","5","/","+"]
 * Output: 6
 * Explanation: (4 + (13 / 5)) = 6
 *
 *
 * Example 3:
 *
 *
 * Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
 * Output: 22
 * Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
 * = ((10 * (6 / (12 * -11))) + 17) + 5
 * = ((10 * (6 / -132)) + 17) + 5
 * = ((10 * 0) + 17) + 5
 * = (0 + 17) + 5
 * = 17 + 5
 * = 22
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= tokens.length <= 10^4
 * tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the
 * range [-200, 200].
 *
 *
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			res := calculate(stack[len(stack)-2], stack[len(stack)-1], token)
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		} else {
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		}
	}
	return stack[0]
}

func calculate(a int, b int, oper string) int {
	switch oper {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return -1
	}
}

// @lc code=end

