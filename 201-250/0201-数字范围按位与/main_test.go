package _201_数字范围按位与

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name:"t1",
			args:args{
				m: 5,
				n: 7,
			},
			want:4,
		},
		{
			name:"t2",
			args:args{
				m: 0,
				n: 1,
			},
			want:0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}