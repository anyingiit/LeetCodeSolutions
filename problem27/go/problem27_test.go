package main

import "testing"

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		name    string
		nums    []int
		val     int
		wantLen int
		wantOut []int
	}{
		{"leetcode example 1", []int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{"leetcode example 2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 0, 1, 3, 4}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			nums := append([]int(nil), tc.nums...)
			gotLen := removeElement(nums, tc.val)
			if gotLen != tc.wantLen {
				t.Fatalf("removeElement(%v, %d) length = %d, want %d", tc.nums, tc.val, gotLen, tc.wantLen)
			}
			got := nums[:gotLen]
			counts := map[int]int{}
			for _, v := range got {
				counts[v]++
			}
			want := map[int]int{}
			for _, v := range tc.wantOut {
				want[v]++
			}
			for k, v := range want {
				if counts[k] != v {
					t.Fatalf("removeElement(%v, %d) kept %v, want multiset %v", tc.nums, tc.val, got, tc.wantOut)
				}
			}
		})
	}
}
