package string

import "fmt"

// https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/
// 395. Longest Substring with At Least K Repeating Characters
// Find the length of the longest substring T of a given string (consists of lowercase letters only) such that every character in T appears no less than k times.

// Example 1:
// Input: 	s = "aaabb", k = 3
// Output:	3
// The longest substring is "aaa", as 'a' is repeated 3 times.

// Example 2:
// Input:  s = "ababbc", k = 2
// Output: 5
// The longest substring is "ababb", as 'a' is repeated 2 times and 'b' is repeated 3 times.

// LongestSubstring .
//
//         |- dp[i-k] + v.end-v.start + 1
// v.cnt = |- dp[v.last_end] + i - v.last_end
//         |- 0
//
func LongestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}

	if k == 1 {
		return len(s)
	}

	/*
	* origin string is 'abab', k = 2
	* _memo {a: {cnt: 2, start: 0, end: 2}, b: 3: {cnt: 2, start: 1, end: 3}}
	* _dp   [0,0,3,4]
	* result _dp[length-1]
	 */
	_memo := make(map[byte]*help)
	_dp := make([]int, len(s))
	_max := -1

	maxHelp := func(slice []int, start, end int) (maxPos, maxV int) {

		if start < 0 {
			start = 0
		}
		if end > len(slice) {
			end = len(slice)
		}
		maxV = 0
		maxPos = -1
		for i := start; i <= end; i++ {
			if slice[i] > maxV {
				maxV = slice[i]
				maxPos = i
			}
		}
		println(start, end, maxV, maxPos)
		return
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		v, ok := _memo[c]
		if !ok {
			// true: first encouter
			v = &help{
				cnt:     0,
				start:   i,
				end:     -1,
				lastEnd: -1,
			}
			_memo[c] = v
		}

		v.cnt++
		v.lastEnd, v.end = v.end, i

		switch {
		case v.cnt == k:
			pos, maxv := maxHelp(_dp, i-k, i-1)
			if pos == -1 {
				_dp[i] = v.end - v.start + 1
			} else {
				_dp[i] = maxv + (i - pos)
			}
		case v.cnt > k:
			if v.lastEnd >= 0 {
				_dp[i] = _dp[v.lastEnd] + i - v.lastEnd
			} else {
				_dp[i] = i - v.lastEnd
			}
		default:
			_dp[i] = 0
		}

		if _dp[i] > _max {
			_max = _dp[i]
		}
	}

	fmt.Printf("%s,%d == %v\n", s, k, _dp)

	return _max
}

type help struct {
	cnt     int // char count
	start   int // first char index
	end     int // last char index
	lastEnd int // the second last char index
}
