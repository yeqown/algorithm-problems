package string_test

import (
	"testing"

	spkg "github.com/yeqown/algorithm-problems/string"
)

func TestLongestSubstring(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "abababc",
				k: 4,
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				s: "abac",
				k: 1,
			},
			want: 4,
		},
		{
			name: "case 3",
			args: args{
				s: "abaa",
				k: 2,
			},
			want: 4,
		},
		{
			name: "case 4",
			args: args{
				s: "abaacc",
				k: 3,
			},
			want: 4,
		},
		{
			name: "case 5",
			args: args{
				s: "baaccac",
				k: 3,
			},
			want: 6,
		},
		{
			name: "case 5.5",
			args: args{
				s: "aaabbb",
				k: 3,
			},
			want: 6,
		},
		{
			name: "case 6",
			args: args{
				s: "aaacbbb",
				k: 3,
			},
			want: 3,
		},
		{
			name: "case 7",
			args: args{
				s: "ababab",
				k: 3,
			},
			want: 6,
		},
		{
			name: "case 8",
			args: args{
				s: "abdaabbc",
				k: 3,
			},
			want: 7,
		},
		{
			name: "case 9",
			args: args{
				s: "ababacb",
				k: 3,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spkg.LongestSubstring(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("LongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
