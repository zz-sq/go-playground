/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] Find Median from Data Stream
 *
 * https://leetcode.cn/problems/find-median-from-data-stream/description/
 *
 * algorithms
 * Hard (54.78%)
 * Likes:    1013
 * Dislikes: 0
 * Total Accepted:    155.4K
 * Total Submissions: 283.7K
 * Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n' +
  '[[],[1],[2],[],[3],[]]'
 *
 * The median is the middle value in an ordered integer list. If the size of
 * the list is even, there is no middle value, and the median is the mean of
 * the two middle values.
 *
 *
 * For example, for arr = [2,3,4], the median is 3.
 * For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
 *
 *
 * Implement the MedianFinder class:
 *
 *
 * MedianFinder() initializes the MedianFinder object.
 * void addNum(int num) adds the integer num from the data stream to the data
 * structure.
 * double findMedian() returns the median of all elements so far. Answers
 * within 10^-5 of the actual answer will be accepted.
 *
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
 * [[], [1], [2], [], [3], []]
 * Output
 * [null, null, null, 1.5, null, 2.0]
 *
 * Explanation
 * MedianFinder medianFinder = new MedianFinder();
 * medianFinder.addNum(1);    // arr = [1]
 * medianFinder.addNum(2);    // arr = [1, 2]
 * medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
 * medianFinder.addNum(3);    // arr[1, 2, 3]
 * medianFinder.findMedian(); // return 2.0
 *
 *
 *
 * Constraints:
 *
 *
 * -10^5 <= num <= 10^5
 * There will be at least one element in the data structure before calling
 * findMedian.
 * At most 5 * 10^4 calls will be made to addNum and findMedian.
 *
 *
 *
 * Follow up:
 *
 *
 * If all integer numbers from the stream are in the range [0, 100], how would
 * you optimize your solution?
 * If 99% of all integer numbers from the stream are in the range [0, 100], how
 * would you optimize your solution?
 *
 *
*/

// @lc code=start
type MedianFinder struct {
	maxHeap, minHeap *MaxHeap
	size             int
}

func Constructor() MedianFinder {
	return MedianFinder{maxHeap: &MaxHeap{}, minHeap: &MaxHeap{}, size: 0}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxHeap, num)
	heap.Push(this.minHeap, -heap.Pop(this.maxHeap).(int))
	this.size++
	if this.size%2 == 1 {
		heap.Push(this.maxHeap, -heap.Pop(this.minHeap).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.size%2 == 0 {
		return float64((this.maxHeap.Top() - this.minHeap.Top())) / 2
	} else {
		return float64(this.maxHeap.Top())
	}
}

type MaxHeap struct {
	nums []int
}

func (maxH *MaxHeap) Len() int {
	return len(maxH.nums)
}

func (maxH *MaxHeap) Swap(i, j int) {
	maxH.nums[i], maxH.nums[j] = maxH.nums[j], maxH.nums[i]
}

func (maxH *MaxHeap) Less(i, j int) bool {
	return maxH.nums[i] > maxH.nums[j]
}

func (maxH *MaxHeap) Push(x interface{}) {
	maxH.nums = append(maxH.nums, x.(int))
}

func (maxH *MaxHeap) Pop() interface{} {
	x := maxH.nums[maxH.Len()-1]
	maxH.nums = maxH.nums[:maxH.Len()-1]
	return x
}

func (maxH *MaxHeap) Top() int {
	return maxH.nums[0]
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

