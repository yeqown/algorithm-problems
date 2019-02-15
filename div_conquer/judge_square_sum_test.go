package div_conquer

import "testing"

func Test_judgeSquareSum(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 0",
			args: args{c: 5},
			want: true,
		},
		{
			name: "case 1",
			args: args{c: 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeSquareSum(tt.args.c); got != tt.want {
				t.Errorf("judgeSquareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
