package tree

import (
	"reflect"
	"testing"
)

// {1,2,3,4,5} to be:
//
// \--  15
//     |--  6
//     |   |--  3
//     |   \--  3
//     |       |--  1
//     |       \--  2
//     \--  9
//         |--  4
//         \--  5

func Test_construtor(t *testing.T) {
	type args struct {
		weights []int
	}
	tests := []struct {
		name string
		args args
		want *huffmanTreeNode
	}{
		{
			name: "case 0",
			args: args{
				weights: []int{1, 2, 3, 4, 5},
			},
			want: &huffmanTreeNode{
				Weight: 15,
				Left: &huffmanTreeNode{
					Weight: 6,
					Left: &huffmanTreeNode{
						Weight: 3,
						Left:   nil,
						Right:  nil,
					},
					Right: &huffmanTreeNode{
						Weight: 3,
						Left: &huffmanTreeNode{
							Weight: 1,
							Left:   nil,
							Right:  nil,
						},
						Right: &huffmanTreeNode{
							Weight: 2,
							Left:   nil,
							Right:  nil,
						},
					},
				},
				Right: &huffmanTreeNode{
					Weight: 9,
					Left: &huffmanTreeNode{
						Weight: 4,
						Left:   nil,
						Right:  nil,
					},
					Right: &huffmanTreeNode{
						Weight: 5,
						Left:   nil,
						Right:  nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := construtor(tt.args.weights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("construtor() = %v, want %v", got, tt.want)

				got.print()
				tt.want.print()
			}
		})
	}
}
