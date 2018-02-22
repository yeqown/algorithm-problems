package array

import (
	"strings"
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
	order := "RLRRLRLRLRLRLRR"
	excepted_p := Point{0, 0}

	robot := NewRobot()
	result_p := robot.LoadOrder(order)
	if excepted_p.Y != result_p.X && excepted_p.Y != result_p.Y {
		t.Fatal("not right answer", result_p)
	}
}

func Test_RobotWalk_case2(t *testing.T) {
	order := "FFFFFFFF"
	excepted_p := Point{0, 8}

	robot := NewRobot()
	result_p := robot.LoadOrder(order)
	if excepted_p.Y != result_p.X && excepted_p.Y != result_p.Y {
		t.Fatal("not right answer", result_p)
	}
}

func Test_RobotWalkDecodeOrder_case3(t *testing.T) {
	order := "FR3(L2(RR)B)LLFRBLF"
	decoded_order := strings.Replace("FR L RRRR B L RRRR B L RRRR B LLFRBLF", " ", "", -1)

	robot := NewRobot()

	if d := robot.DecodeOrder(order); d != decoded_order {
		t.Fatalf("error decoder, want (%s), have (%s)", decoded_order, d)
	}

}
