package array

import (
	"reflect"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
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
				[]int{4, 3, 2, 7, 8, 2, 3},
			},
			want: []int{1, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDisappearedNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
