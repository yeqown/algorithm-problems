/*
有一个机器人，给一串指令，L左转 R右转，F前进一步，B后退一步，问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。
可以输入重复指令n ： 比如 R2(LF) 这个等于指令 RLFLF。
问最后机器人的坐标是多少？
*/

package array

import (
	"fmt"
	// "time"
	// "strconv"
)

const (
	// Top 0
	Top = iota
	// Left 1
	Left
	// Down 2
	Down
	// Right 3
	Right
)

// Point ...
type Point struct {
	X int
	Y int
}

func (p Point) String() {
	fmt.Printf("Point(%d,%d)\n", p.X, p.Y)
}

// Robot ...
type Robot struct {
	Direct int    // 方向
	Step   int    // 是否需要前进一步
	Pos    *Point // 当前位置
}

// NewRobot ...
func NewRobot() *Robot {
	return &Robot{
		Direct: Top,
		Step:   0,
		Pos:    &Point{0, 0},
	}
}

/*
according to $orderChar to deside which direction to walk or step
*/
func (r *Robot) judgeOrderChar(orderChar rune) {
	switch orderChar {
	// case "L":
	case 76:
		r.Direct = (r.Direct + 1) % 4
	// case "R":
	case 82:
		r.Direct = (r.Direct + 3) % 4
	// case "F":
	case 70:
		r.Step = 1
	// case "B":
	case 66:
		r.Step = -1
	default:
		panic("shouldn't be in this case")
	}
}

// DecodeOrder decode `FR3(L2(RR)B)LLFRBLF` to `FR LLL RRRRRRRRRRRR BBB LLFRBLF`
func (r *Robot) DecodeOrder(order string) string {

	var (
		dupTime    = 1
		lastDupPos = -1
	)
Loop:
	for i, c := range order {

		if c == 40 {
			lastDupPos = i
		} else if c == 41 {
			dupStr := order[lastDupPos+1 : i]
			s := []byte{}
			s = append(s, order[:lastDupPos-1]...)
			for i := 0; i < dupTime; i++ {
				s = append(s, dupStr...)
			}
			s = append(s, order[i+1:]...)
			order = string(s)

			fmt.Println("new order: ", order)
			// time.Sleep(1 * time.Second)

			// reset
			lastDupPos = -1
			dupTime = 1

			// reloop
			goto Loop
		} else if c >= 48 && c <= 57 {
			dupTime = int(c) - 48
		}
		// if char is invalid then panic
	}
	return order
}

/*
 robot walk step and reset step
*/
func (r *Robot) walk() {
	switch r.Direct {
	case Top:
		r.Pos.Y += r.Step
	case Down:
		r.Pos.Y -= r.Step
	case Right:
		r.Pos.X += r.Step
	case Left:
		r.Pos.X -= r.Step
	}
	r.Step = 0
}

// LoadOrder ...
func (r *Robot) LoadOrder(order string) Point {
	for _, char := range order {
		r.judgeOrderChar(char)
		r.walk()
		// fmt.Printf("Char %v, Robot: %v\n", string(char), r)
	}
	return *r.Pos
}
