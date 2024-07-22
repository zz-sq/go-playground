/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 *
 * https://leetcode.cn/problems/minimum-genetic-mutation/description/
 *
 * algorithms
 * Medium (54.37%)
 * Likes:    296
 * Dislikes: 0
 * Total Accepted:    69.7K
 * Total Submissions: 128.2K
 * Testcase Example:  '"AACCGGTT"\n"AACCGGTA"\n["AACCGGTA"]'
 *
 * A gene string can be represented by an 8-character long string, with choices
 * from 'A', 'C', 'G', and 'T'.
 *
 * Suppose we need to investigate a mutation from a gene string startGene to a
 * gene string endGene where one mutation is defined as one single character
 * changed in the gene string.
 *
 *
 * For example, "AACCGGTT" --> "AACCGGTA" is one mutation.
 *
 *
 * There is also a gene bank bank that records all the valid gene mutations. A
 * gene must be in bank to make it a valid gene string.
 *
 * Given the two gene strings startGene and endGene and the gene bank bank,
 * return the minimum number of mutations needed to mutate from startGene to
 * endGene. If there is no such a mutation, return -1.
 *
 * Note that the starting point is assumed to be valid, so it might not be
 * included in the bank.
 *
 *
 * Example 1:
 *
 *
 * Input: startGene = "AACCGGTT", endGene = "AACCGGTA", bank = ["AACCGGTA"]
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: startGene = "AACCGGTT", endGene = "AAACGGTA", bank =
 * ["AACCGGTA","AACCGCTA","AAACGGTA"]
 * Output: 2
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= bank.length <= 10
 * startGene.length == endGene.length == bank[i].length == 8
 * startGene, endGene, and bank[i] consist of only the characters ['A', 'C',
 * 'G', 'T'].
 *
 *
 */

// @lc code=start
func minMutation(startGene string, endGene string, bank []string) int {
	type Pair struct {
		Gene string
		Step int
	}
	queue := []Pair{{startGene, 0}}
	visted := map[string]bool{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if distance(cur.Gene, endGene) == 0 {
			return cur.Step
		}
		for _, candidate := range bank {
			if distance(cur.Gene, candidate) == 1 && !visted[candidate] {
				visted[candidate] = true
				queue = append(queue, Pair{candidate, cur.Step + 1})
			}
		}
	}
	return -1
}

func distance(m, n string) int {
	d := 0
	for i := 0; i < len(m); i++ {
		if m[i] != n[i] {
			d++
		}
	}
	return d
}

// @lc code=end

