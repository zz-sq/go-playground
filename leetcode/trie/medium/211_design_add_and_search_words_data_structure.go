/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] Design Add and Search Words Data Structure
 *
 * https://leetcode.cn/problems/design-add-and-search-words-data-structure/description/
 *
 * algorithms
 * Medium (50.16%)
 * Likes:    572
 * Dislikes: 0
 * Total Accepted:    91.8K
 * Total Submissions: 182.9K
 * Testcase Example:  '["WordDictionary","addWord","addWord","addWord","search","search","search","search"]\n' +
  '[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]'
 *
 * Design a data structure that supports adding new words and finding if a
 * string matches any previously added string.
 *
 * Implement the WordDictionary class:
 *
 *
 * WordDictionary() Initializes the object.
 * void addWord(word) Adds word to the data structure, it can be matched
 * later.
 * bool search(word) Returns true if there is any string in the data structure
 * that matches word or false otherwise. word may contain dots '.' where dots
 * can be matched with any letter.
 *
 *
 *
 * Example:
 *
 *
 * Input
 *
 * ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
 * [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
 * Output
 * [null,null,null,null,false,true,true,true]
 *
 * Explanation
 * WordDictionary wordDictionary = new WordDictionary();
 * wordDictionary.addWord("bad");
 * wordDictionary.addWord("dad");
 * wordDictionary.addWord("mad");
 * wordDictionary.search("pad"); // return False
 * wordDictionary.search("bad"); // return True
 * wordDictionary.search(".ad"); // return True
 * wordDictionary.search("b.."); // return True
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= word.length <= 25
 * word in addWord consists of lowercase English letters.
 * word in search consist of '.' or lowercase English letters.
 * There will be at most 2 dots in word for search queries.
 * At most 10^4 calls will be made to addWord and search.
 *
 *
*/

// @lc code=start
type WordDictionary struct {
	children [26]*WordDictionary
	isWord   bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this
	for _, ch := range word {
		idx := ch - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = &WordDictionary{}
		}
		cur = cur.children[idx]
	}
	cur.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	type Pair struct {
		Wd *WordDictionary
		i  int
	}
	queue := []Pair{{this, 0}}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		i := curr.i
		cur := curr.Wd
		if i > len(word)-1 {
			break
		}
		ch := word[i]
		next := []Pair{}
		if ch == '.' {
			for _, v := range cur.children {
				if v != nil {
					next = append(next, Pair{v, i + 1})
				}
			}
		} else {
			idx := ch - 'a'
			if cur.children[idx] != nil {
				next = append(next, Pair{cur.children[idx], i + 1})
			}
		}
		if i == len(word)-1 && len(next) > 0 {
			for _, v := range next {
				if v.Wd.isWord {
					return true
				}
			}
		}
		queue = append(queue, next...)
		i++
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

