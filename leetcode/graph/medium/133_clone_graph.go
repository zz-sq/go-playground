/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] Clone Graph
 *
 * https://leetcode.cn/problems/clone-graph/description/
 *
 * algorithms
 * Medium (71.03%)
 * Likes:    726
 * Dislikes: 0
 * Total Accepted:    151.4K
 * Total Submissions: 213.1K
 * Testcase Example:  '[[2,4],[1,3],[2,4],[1,3]]'
 *
 * Given a reference of a node in a connected undirected graph.
 *
 * Return a deep copy (clone) of the graph.
 *
 * Each node in the graph contains a value (int) and a list (List[Node]) of its
 * neighbors.
 *
 *
 * class Node {
 * ⁠   public int val;
 * ⁠   public List<Node> neighbors;
 * }
 *
 *
 *
 *
 * Test case format:
 *
 * For simplicity, each node's value is the same as the node's index
 * (1-indexed). For example, the first node with val == 1, the second node with
 * val == 2, and so on. The graph is represented in the test case using an
 * adjacency list.
 *
 * An adjacency list is a collection of unordered lists used to represent a
 * finite graph. Each list describes the set of neighbors of a node in the
 * graph.
 *
 * The given node will always be the first node with val = 1. You must return
 * the copy of the given node as a reference to the cloned graph.
 *
 *
 * Example 1:
 *
 *
 * Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
 * Output: [[2,4],[1,3],[2,4],[1,3]]
 * Explanation: There are 4 nodes in the graph.
 * 1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val =
 * 4).
 * 2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val =
 * 3).
 * 3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val =
 * 4).
 * 4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val =
 * 3).
 *
 *
 * Example 2:
 *
 *
 * Input: adjList = [[]]
 * Output: [[]]
 * Explanation: Note that the input contains one empty list. The graph consists
 * of only one node with val = 1 and it does not have any neighbors.
 *
 *
 * Example 3:
 *
 *
 * Input: adjList = []
 * Output: []
 * Explanation: This an empty graph, it does not have any nodes.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the graph is in the range [0, 100].
 * 1 <= Node.val <= 100
 * Node.val is unique for each node.
 * There are no repeated edges and no self-loops in the graph.
 * The Graph is connected and all nodes can be visited starting from the given
 * node.
 *
 *
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	nodeMap := map[*Node]*Node{}
	return dfs(nodeMap, node)
}

func dfs(nodeMap map[*Node]*Node, node *Node) *Node {
	if newNode, ok := nodeMap[node]; ok {
		return newNode
	}
	root := &Node{}
	root.Val = node.Val
	nodeMap[node] = root
	for _, n := range node.Neighbors {
		root.Neighbors = append(root.Neighbors, dfs(nodeMap, n))
	}
	return root
}

// @lc code=end

