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
	hour  = []int{1, 2, 4, 8}
	min   = []int{1, 2, 4, 8, 16, 32}
	lenH  = 4
	times = []string{}
)

func help(h, m, num, startPoint int) {

	if num == 0 {
		// fmt.Println(h, m, num, startPoint)
		times = append(times, fmt.Sprintf("%d:%02d", h, m))
		return
	}

	for i := startPoint; i < len(hour)+len(min); i++ {
		if i < lenH {
			h += hour[i]
			if h < 12 {
				help(h, m, num-1, i+1)
			}
			h -= hour[i]
		} else {
			m += min[i-lenH]
			if m < 60 {
				help(h, m, num-1, i+1)
			}
			m -= min[i-lenH]
		}
	}
}

// ReadBinayWatch ...
func ReadBinayWatch(num int) []string {
	help(0, 0, num, 0)
	return times
}
