/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] Interleaving String
 *
 * https://leetcode.cn/problems/interleaving-string/description/
 *
 * algorithms
 * Medium (45.06%)
 * Likes:    1021
 * Dislikes: 0
 * Total Accepted:    154.5K
 * Total Submissions: 342.8K
 * Testcase Example:  '"aabcc"\n"dbbca"\n"aadbbcbcac"'
 *
 * Given strings s1, s2, and s3, find whether s3 is formed by an interleaving
 * of s1 and s2.
 *
 * An interleaving of two strings s and t is a configuration where s and t are
 * divided into n and m substrings respectively, such that:
 *
 *
 * s = s1 + s2 + ... + sn
 * t = t1 + t2 + ... + tm
 * |n - m| <= 1
 * The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 +
 * t3 + s3 + ...
 *
 *
 * Note: a + b is the concatenation of strings a and b.
 *
 *
 * Example 1:
 *
 *
 * Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
 * Output: true
 * Explanation: One way to obtain s3 is:
 * Split s1 into s1 = "aa" + "bc" + "c", and s2 into s2 = "dbbc" + "a".
 * Interleaving the two splits, we get "aa" + "dbbc" + "bc" + "a" + "c" =
 * "aadbbcbcac".
 * Since s3 can be obtained by interleaving s1 and s2, we return true.
 *
 *
 * Example 2:
 *
 *
 * Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
 * Output: false
 * Explanation: Notice how it is impossible to interleave s2 with any other
 * string to obtain s3.
 *
 *
 * Example 3:
 *
 *
 * Input: s1 = "", s2 = "", s3 = ""
 * Output: true
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= s1.length, s2.length <= 100
 * 0 <= s3.length <= 200
 * s1, s2, and s3 consist of lowercase English letters.
 *
 *
 *
 * Follow up: Could you solve it using only O(s2.length) additional memory
 * space?
 *
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m := len(s1), len(s2)
	if n+m != len(s3) {
		return false
	}
	dp := make([]bool, m+1)
	dp[0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				dp[j] = dp[j] || dp[j-1] && s2[j-1] == s3[p]
			}
		}
	}
	return dp[m]
}

func isInterleaveDfs(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if i == m {
			return s2[j:] == s3[k:]
		}
		if j == n {
			return s1[i:] == s3[k:]
		}
		if visited[i][j] != 0 {
			return visited[i][j] == 1
		}
		res := false
		res = s1[i] == s3[k] && dfs(i+1, j, k+1)
		if !res {
			res = s2[j] == s3[k] && dfs(i, j+1, k+1)
		}
		if res {
			visited[i][j] = 1
		} else {
			visited[i][j] = 2
		}
		return res
	}
	return dfs(0, 0, 0)
}

// @lc code=end

