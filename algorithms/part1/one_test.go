package part1

import "testing"

import "reflect"

func TestTwoSum(t *testing.T) {
	type arg struct {
		nums   []int
		target int
		want   []int
	}
	args := []arg{
		arg{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
	}
	for _, v := range args {
		result := twoSum(v.nums, v.target)
		if !reflect.DeepEqual(result, v.want) {
			t.Errorf("want: %v, result: %v\n", v.want, result)
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	type arg struct {
		node []*ListNode
		want *ListNode
	}
	args := []arg{
		arg{
			[]*ListNode{&ListNode{2, &ListNode{4, &ListNode{Val: 3}}}, &ListNode{5, &ListNode{6, &ListNode{Val: 4}}}},
			&ListNode{7, &ListNode{0, &ListNode{Val: 8}}},
		},
		arg{
			[]*ListNode{&ListNode{2, &ListNode{4, &ListNode{Val: 3}}}, &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{Val: 6}}}}},
			&ListNode{7, &ListNode{0, &ListNode{8, &ListNode{Val: 6}}}},
		},
	}
	for _, v := range args {
		r := addTwoNumbers(v.node[0], v.node[1])
		if !reflect.DeepEqual(r, v.want) {
			t.Errorf("want: %v, result: %#v\n", v.want, r)
		}
	}

}

func TestLengthOfLongestSubstring(t *testing.T) {
	args := map[string]int{
		"abcabcbb": 3,
		"bbbb":     1,
		"pwwkew":   3,
		"":         0,
		"a":        1,
		"ab":       2,
		"dvdf":     3,
	}
	for str, want := range args {
		r := lengthOfLongestSubstring(str)
		if r != want {
			t.Errorf("string: %s, want: %d, result: %d", str, want, r)
		}
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("abcabcbb")
	}
}
