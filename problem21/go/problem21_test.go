package main

import "testing"

func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func listToSlice(n *ListNode) []int {
	var out []int
	for n != nil {
		out = append(out, n.Val)
		n = n.Next
	}
	return out
}

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		name   string
		l1, l2 []int
		want   []int
	}{
		{"both non-empty", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{"first empty", nil, []int{0}, []int{0}},
		{"both empty", nil, nil, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := listToSlice(mergeTwoLists(buildList(tc.l1), buildList(tc.l2)))
			if len(got) != len(tc.want) {
				t.Fatalf("mergeTwoLists(%v, %v) = %v, want %v", tc.l1, tc.l2, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("mergeTwoLists(%v, %v) = %v, want %v", tc.l1, tc.l2, got, tc.want)
				}
			}
		})
	}
}
