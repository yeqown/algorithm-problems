package branch_bound

import (
	"fmt"
	"reflect"
	"testing"
)

func debugPrint() {
	for k, v := range powOf2 {
		fmt.Printf("%d, %v\n", k, v)
	}
}

func Test_convert(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want numNode
	}{
		{
			name: " case 0",
			args: args{n: 123456},
			want: numNode{len: 6, nums: map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1}},
		},
		{
			name: " case 1",
			args: args{n: 112233456},
			want: numNode{len: 9, nums: map[int]int{1: 2, 2: 2, 3: 2, 4: 1, 5: 1, 6: 1}},
		},
		{
			name: " case 2",
			args: args{n: 1},
			want: numNode{len: 1, nums: map[int]int{1: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reorderedPowerOf2(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 0",
			args: args{N: 8},
			want: true,
		},
		{
			name: "case 1",
			args: args{N: 46},
			want: true,
		},
		{
			name: "case 0",
			args: args{N: 8},
			want: true,
		},
		{
			name: "case 2",
			args: args{N: 536870912},
			want: true,
		},
		{
			name: "case 3",
			args: args{N: 538607921},
			want: true,
		},
	}

	debugPrint()
	t.Fail()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderedPowerOf2(tt.args.N); got != tt.want {
				t.Errorf("reorderedPowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numNode_equal(t *testing.T) {
	type fields struct {
		nums map[int]int
		len  int
	}
	type args struct {
		node numNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: " case 0",
			fields: fields{
				nums: map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1},
				len:  6,
			},
			args: args{
				node: numNode{
					len:  6,
					nums: map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1},
				},
			},
			want: true,
		},
		{
			name: " case 0",
			fields: fields{
				nums: map[int]int{1: 1, 2: 1, 7: 1, 4: 1, 5: 1, 6: 1},
				len:  6,
			},
			args: args{
				node: numNode{
					len:  6,
					nums: map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := numNode{
				nums: tt.fields.nums,
				len:  tt.fields.len,
			}
			if got := n.equal(tt.args.node); got != tt.want {
				t.Errorf("numNode.equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
