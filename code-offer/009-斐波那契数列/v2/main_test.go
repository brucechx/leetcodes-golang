package v2

import "testing"

func TestFib(t *testing.T) {
	cases := []struct{
		input int
		output int
	}{
		{0, 0},
		{1, 1},
		{3, 2},
	}

	for _, cas := range cases{
		got := Fib2(cas.input)
		want := cas.output
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
}
