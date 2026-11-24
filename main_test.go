package main

import "testing"

func TestCountSubarraysXor(t *testing.T) {
	cases := []struct {
		in   []int
		k    int
		want int
	}{
		{[]int{4, 2, 2, 6, 4}, 6, 4},
		{[]int{5, 6, 7, 8, 9}, 5, 2},
		{[]int{1, 1, 1}, 0, 2}, // [1,1] twice
		{[]int{}, 3, 0},
		{[]int{3}, 3, 1},
		{[]int{0, 0}, 0, 3},
	}
	for _, c := range cases {
		if got := CountSubarraysXor(c.in, c.k); got != c.want {
			t.Errorf("CountSubarraysXor(%v,%d) = %d, want %d", c.in, c.k, got, c.want)
		}
	}
}
