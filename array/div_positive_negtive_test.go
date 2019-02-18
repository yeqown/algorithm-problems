package array

import (
	"reflect"
	"testing"
)

func Test_divPositiveNegative(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 0",
			args: args{
				nums: []int{-1, -99, 90, 23, 55, -8, 23, 12, -39, -80, -92, 12},
			},
			want: []int{12, 12, 90, 23, 55, 23, -99, -8, -39, -80, -92, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divPositiveNegative(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divPositiveNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}
