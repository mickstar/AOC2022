package main

import "testing"

func Test_calculateWeight(t *testing.T) {
	type args struct {
		a uint8
		b uint8
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{'a', 'c'},
			want: 2,
		},
		{
			name: "test 1",
			args: args{'a', 'b'},
			want: 1,
		},
		{
			name: "test 1",
			args: args{'b', 'a'},
			want: -1,
		},
		{
			name: "test 1",
			args: args{'S', 'a'},
			want: 1,
		},
		{
			name: "test 1",
			args: args{'x', 'E'},
			want: -1,
		},
		{
			name: "test 1",
			args: args{'x', 'x'},
			want: 1,
		},
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWeight(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("calculateWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
