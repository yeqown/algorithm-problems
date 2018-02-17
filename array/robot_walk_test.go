package array

import (
	"testing"
)

func Test_RobotWalk(t *testing.T) {
	order := "FRRB"
	excepted_p := Point{0, 2}

	robot := NewRobot()
	result_p := robot.LoadOrder(order)
	if excepted_p.Y != result_p.X && excepted_p.Y != result_p.Y {
		t.Fatal("not right answer", result_p)
	}
}

func Test_RobotWalk_case1(t *testing.T) {
	order := "FRRBLFRB"
	excepted_p := Point{0, 2}

	robot := NewRobot()
	result_p := robot.LoadOrder(order)
	if excepted_p.Y != result_p.X && excepted_p.Y != result_p.Y {
		t.Fatal("not right answer", result_p)
	}
}

func Test_RobotWalk_case2(t *testing.T) {
	order := "FRRRBLLFRBLF"
	excepted_p := Point{0, 2}

	robot := NewRobot()
	result_p := robot.LoadOrder(order)
	if excepted_p.Y != result_p.X && excepted_p.Y != result_p.Y {
		t.Fatal("not right answer", result_p)
	}
}
