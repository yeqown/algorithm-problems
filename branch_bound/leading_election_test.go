package branch_bound

import (
	"fmt"
	"reflect"
	"testing"
)

func run(persons []int, times []int, qs []int) []int {
	vc := Constructor(persons, times)
	fmt.Printf("got TopVotedCandidate: %v\n", vc)
	got := make([]int, len(qs))
	for idx, q := range qs {
		got[idx] = vc.Q(q)
	}
	return got
}

func TestTopVotedCandidate(t *testing.T) {
	type args struct {
		persons []int
		times   []int
		qs      []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 0",
			args: args{
				persons: []int{0, 0, 1, 1, 2},
				times:   []int{0, 67, 69, 74, 87},
				qs:      []int{4, 62, 100, 88, 70, 73, 22, 75, 29, 10},
			},
			// 0,0,1,1
			want: []int{0, 0, 1, 1, 0, 0, 0, 1, 0, 0},
		},
		{
			name: "case 1",
			args: args{
				persons: []int{0, 1, 1, 0, 0, 1, 0},
				times:   []int{0, 5, 10, 15, 20, 25, 30},
				qs:      []int{3, 12, 25, 15, 24, 8},
			},
			// 0,1,1,0,0,1
			want: []int{0, 1, 1, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.persons, tt.args.times, tt.args.qs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
