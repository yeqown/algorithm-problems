// 二进制手表

/*
* Binay Watch
* Q: A binary watch has 4 LEDs on the top which represent the hours (0-11), and the 6 LEDs on the bottom represent the minutes (0-59).
Each LED represents a zero or one, with the least significant bit on the right.
Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times the watch could represent.
*/

package backtracking

import (
	"fmt"
)

/*
* @num 手表上LED亮着的个数
* @times 可能的时间
 */

var (
	HOUR  = []int{1, 2, 4, 8}
	MIN   = []int{1, 2, 4, 8, 16, 32}
	LEN_H = 4
	LEN_M = 6
	times = []string{}
)

func help(h, m, num, start_point int) {

	if num == 0 {
		// fmt.Println(h, m, num, start_point)
		times = append(times, fmt.Sprintf("%d:%02d", h, m))
		return
	}

	for i := start_point; i < len(HOUR)+len(MIN); i++ {
		if i < LEN_H {
			h += HOUR[i]
			if h < 12 {
				help(h, m, num-1, i+1)
			}
			h -= HOUR[i]
		} else {
			m += MIN[i-LEN_H]
			if m < 60 {
				help(h, m, num-1, i+1)
			}
			m -= MIN[i-LEN_H]
		}
	}
}

func ReadBinayWatch(num int) []string {
	help(0, 0, num, 0)
	return times
}
