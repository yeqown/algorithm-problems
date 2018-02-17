/*
有一个机器人，给一串指令，L左转 R右转，F前进一步，B后退一步，问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。
可以输入重复指令n ： 比如 R2(LF) 这个等于指令 RLFLF。
问最后机器人的坐标是多少？
*/

/*
！！没有实现重复指令！！
*/
package array

import (
	"fmt"
)

const (
	Top   = iota // 0
	Left         // 1
	Down         // 2
	Right        // 3
)

type Point struct {
	X int
	Y int
}

func (p Point) String() {
	fmt.Printf("Point(%d,%d)\n", p.X, p.Y)
}

type Robot struct {
	Direct int    // 方向
	Step   int    // 是否需要前进一步
	Pos    *Point // 当前位置
}

func NewRobot() *Robot {
	return &Robot{
		Direct: Top,
		Step:   0,
		Pos:    &Point{0, 0},
	}
}

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

func (r *Robot) LoadOrder(order string) Point {
	for _, char := range order {
		r.judgeOrderChar(char)
		r.walk()
		// fmt.Printf("Char %v, Robot: %v\n", string(char), r)
	}
	return *r.Pos
}
