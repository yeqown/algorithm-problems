package array_test

import (
	"reflect"
	"testing"

	"github.com/yeqown/algorithm-problems/array"
)

func TestFindPairInArray(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []array.Pair
	}{
		{
			name: "case 0",
			args: args{
				arr:    []int{-3, -1, 0, 1, 3, 6, 9},
				target: 6,
			},
			want: []array.Pair{
				{NumA: -3, NumB: 9},
				{NumA: 0, NumB: 6},
			},
		},
		{
			name: "case 1",
			args: args{
				arr:    []int{-3, -1, 0, 1, 3, 6, 6, 9},
				target: 6,
			},
			want: []array.Pair{
				{NumA: -3, NumB: 9},
				{NumA: 0, NumB: 6},
				{NumA: 0, NumB: 6},
			},
		},
		{
			name: "case 2",
			args: args{
				arr:    []int{-3, -3, -1, 0, 1, 3, 6, 6, 9},
				target: 6,
			},
			want: []array.Pair{
				{NumA: -3, NumB: 9},
				{NumA: -3, NumB: 9},
				{NumA: 0, NumB: 6},
				{NumA: 0, NumB: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.FindPairInArray(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPairInArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
