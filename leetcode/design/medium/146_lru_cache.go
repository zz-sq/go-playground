/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU Cache
 *
 * https://leetcode.cn/problems/lru-cache/description/
 *
 * algorithms
 * Medium (53.69%)
 * Likes:    3212
 * Dislikes: 0
 * Total Accepted:    665.7K
 * Total Submissions: 1.2M
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * Design a data structure that follows the constraints of a Least Recently
 * Used (LRU) cache.
 *
 * Implement the LRUCache class:
 *
 *
 * LRUCache(int capacity) Initialize the LRU cache with positive size
 * capacity.
 * int get(int key) Return the value of the key if the key exists, otherwise
 * return -1.
 * void put(int key, int value) Update the value of the key if the key exists.
 * Otherwise, add the key-value pair to the cache. If the number of keys
 * exceeds the capacity from this operation, evict the least recently used
 * key.
 *
 *
 * The functions get and put must each run in O(1) average time complexity.
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
 * Output
 * [null, null, null, 1, null, -1, null, -1, 3, 4]
 *
 * Explanation
 * LRUCache lRUCache = new LRUCache(2);
 * lRUCache.put(1, 1); // cache is {1=1}
 * lRUCache.put(2, 2); // cache is {1=1, 2=2}
 * lRUCache.get(1);    // return 1
 * lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
 * lRUCache.get(2);    // returns -1 (not found)
 * lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
 * lRUCache.get(1);    // return -1 (not found)
 * lRUCache.get(3);    // return 3
 * lRUCache.get(4);    // return 4
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= capacity <= 3000
 * 0 <= key <= 10^4
 * 0 <= value <= 10^5
 * At most 2 * 10^5 calls will be made to get and put.
 *
 *
*/

// @lc code=start
type LRUCache struct {
	Capacity int
	Size     int
	Data     map[int]*LinkListNode
	Head     *LinkListNode
	Tail     *LinkListNode
}

type LinkListNode struct {
	Key  int
	Val  int
	Pre  *LinkListNode
	Next *LinkListNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{}
	cache.Capacity = capacity
	cache.Data = map[int]*LinkListNode{}
	cache.Head = &LinkListNode{}
	cache.Tail = &LinkListNode{}
	cache.Head.Next = cache.Tail
	cache.Tail.Pre = cache.Head
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Data[key]; ok {
		this.moveToHead(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) moveToHead(node *LinkListNode) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) addNode(node *LinkListNode) {
	node.Pre = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Pre = node
	this.Head.Next = node
}

func (this *LRUCache) removeNode(node *LinkListNode) {
	pre := node.Pre
	next := node.Next
	pre.Next = next
	next.Pre = pre
}

func (this *LRUCache) popTail() *LinkListNode {
	tail := this.Tail.Pre
	this.removeNode(tail)
	return tail
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Data[key]; ok {
		node.Val = value
		this.moveToHead(node)
		return
	}
	newNode := &LinkListNode{}
	newNode.Val = value
	newNode.Key = key
	this.Data[key] = newNode
	this.addNode(newNode)
	this.Size++
	if this.Size > this.Capacity {
		removeNode := this.popTail()
		delete(this.Data, removeNode.Key)
		this.Size--
	}
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

