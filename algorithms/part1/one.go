package part1

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
// 示例:
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]
func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i, v := range nums {
		for j := i + 1; j < length; j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// 示例：
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var addNode func(*ListNode, *ListNode, int) *ListNode

	addNode = func(v1, v2 *ListNode, sum int) *ListNode {
		t := v1.Val + v2.Val + sum
		c := t % 10
		e := t / 10
		result := &ListNode{Val: c}
		if v1.Next != nil || v2.Next != nil || e > 0 {
			vv1, vv2 := v1.Next, v2.Next
			if vv1 == nil {
				vv1 = &ListNode{}
			}
			if vv2 == nil {
				vv2 = &ListNode{}
			}
			result.Next = addNode(vv1, vv2, e)
		}
		return result
	}

	return addNode(l1, l2, 0)
}

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

// 示例 1:

// 输入: "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:

// 输入: "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:

// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

// 3 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	m, length, index := map[rune]int{}, 0, -1

	for i, v := range s {
		pre, ok := m[v]
		if !ok {
			pre = -1
		}
		if pre > index {
			index = pre
		}
		if i-index > length {
			length = i - index
		}
		m[v] = i
	}
	return length
}

// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

// 你可以假设 nums1 和 nums2 不会同时为空。

// 示例 1:
// [2], [3] i:0 j:0 c:1 o:f
// nums1 = [1, 3]
// nums2 = [2]   i:0 j:0 c:1 o:t

// 则中位数是 2.0
// 示例 2:

// nums1 = [1, 2]
// nums2 = [3, 4] i:1 j:0 c:2 o:f

// 则中位数是 (2 + 3)/2 = 2.5
