package string

import (
	"reflect"
	"testing"
)

func TestSunday(t *testing.T) {
	type args struct {
		str    string
		substr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				str:    "abaac",
				substr: "aac",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				str:    "substring searching alg",
				substr: "search",
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				str:    "searcsearchsearcsearc",
				substr: "search",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sunday(tt.args.str, tt.args.substr); got != tt.want {
				t.Errorf("Sunday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sundayHelper(t *testing.T) {
	type args struct {
		substr string
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
		{
			name: "case 1",
			args: args{
				substr: "abacaaa",
			},
			want: map[rune]int{
				'a': 1,
				'b': 6,
				'c': 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sundayHelper(tt.args.substr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sundayHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}
