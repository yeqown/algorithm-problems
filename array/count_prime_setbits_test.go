package array

import (
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name       string
		args       args
		wantBitNum int
	}{
		{
			name:       "case 0",
			args:       args{4},
			wantBitNum: 1,
		},
		{
			name:       "case 1",
			args:       args{15},
			wantBitNum: 4,
		},
		{
			name:       "case 2",
			args:       args{14},
			wantBitNum: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBitNum := countBits(tt.args.num); gotBitNum != tt.wantBitNum {
				t.Errorf("countBits() = %v, want %v", gotBitNum, tt.wantBitNum)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 0",
			args: args{4},
			want: false,
		},
		{
			name: "case 1",
			args: args{3},
			want: true,
		},
		{
			name: "case 2",
			args: args{14},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.n); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countPrimeSetBits(t *testing.T) {
	type args struct {
		L int
		R int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 0",
			args: args{L: 6, R: 10},
			want: 4,
		},
		{
			name: "case 1",
			args: args{L: 10, R: 15},
			want: 5,
		},
		{
			name: "case 2",
			args: args{L: 244, R: 269},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimeSetBits(tt.args.L, tt.args.R); got != tt.want {
				t.Errorf("countPrimeSetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
