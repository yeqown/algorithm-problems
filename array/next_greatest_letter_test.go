package array

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "case 0",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'a',
			},
			want: 'c',
		},
		{
			name: "case 1",
			args: args{
				letters: []byte{'a', 'b'},
				target:  'z',
			},
			want: 'a',
		},
		{
			name: "case 2",
			args: args{
				letters: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'},
				target:  'y',
			},
			want: 'z',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byte(t *testing.T) {
	for b := 'a'; b <= 'z'; b++ {
		t.Logf("%v", b)
	}
	t.Fail()
}
