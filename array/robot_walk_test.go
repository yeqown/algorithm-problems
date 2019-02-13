package array

import (
	"strings"
	"testing"
)

func Test_RobotWalk(t *testing.T) {
	order := "FRRB"
	exceptedPoint := Point{0, 2}

	robot := NewRobot()
	resultPoint := robot.LoadOrder(order)
	if exceptedPoint.Y != resultPoint.X && exceptedPoint.Y != resultPoint.Y {
		t.Fatal("not right answer", resultPoint)
	}
}

func Test_RobotWalk_case1(t *testing.T) {
	order := "RLRRLRLRLRLRLRR"
	exceptedPoint := Point{0, 0}

	robot := NewRobot()
	resultPoint := robot.LoadOrder(order)
	if exceptedPoint.Y != resultPoint.X && exceptedPoint.Y != resultPoint.Y {
		t.Fatal("not right answer", resultPoint)
	}
}

func Test_RobotWalk_case2(t *testing.T) {
	order := "FFFFFFFF"
	exceptedPoint := Point{0, 8}

	robot := NewRobot()
	resultPoint := robot.LoadOrder(order)
	if exceptedPoint.Y != resultPoint.X && exceptedPoint.Y != resultPoint.Y {
		t.Fatal("not right answer", resultPoint)
	}
}

func Test_RobotWalkDecodeOrder_case3(t *testing.T) {
	order := "FR3(L2(RR)B)LLFRBLF"
	decodedOrder := strings.Replace("FR L RRRR B L RRRR B L RRRR B LLFRBLF", " ", "", -1)

	robot := NewRobot()

	if d := robot.DecodeOrder(order); d != decodedOrder {
		t.Fatalf("error decoder, want (%s), have (%s)", decodedOrder, d)
	}

}
