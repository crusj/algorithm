package review

import (
	"reflect"
	"testing"

	. "github.com/crusj/algorithm/common"
)

func Test_printListFromTailToHead(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				head: NewSingleList([]int{1, 2, 3}),
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printListFromTailToHead(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printListFromTailToHead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseList(t *testing.T) {
	type args struct {
		pHead *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test",
			args: args{
				pHead: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseList(tt.args.pHead); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args struct {
		pHead1 *ListNode
		pHead2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				pHead1: NewSingleList([]int{1, 2, 3}),
				pHead2: NewSingleList([]int{4, 5, 6}),
			},
			want: NewSingleList([]int{1, 2, 3, 4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.pHead1, tt.args.pHead2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findFirstCommonNode(t *testing.T) {
	a := NewSingleList([]int{1, 2, 3})
	b := NewSingleList([]int{4, 5})
	c := NewSingleList([]int{6, 7})

	a.Next.Next.Next = c
	b.Next.Next = c

	m := NewSingleList([]int{1, 2})
	n := NewSingleList([]int{3, 4})

	type args struct {
		pHead1 *ListNode
		pHead2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test",
			args: args{
				pHead1: a,
				pHead2: b,
			},
			want: c,
		},
		{
			name: "test2",
			args: args{
				pHead1: m,
				pHead2: n,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFirstCommonNode(tt.args.pHead1, tt.args.pHead2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFirstCommonNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntryNodeOfLoop(t *testing.T) {
	a := NewSingleList([]int{1, 2, 3})
	b := NewSingleList([]int{4, 5})

	a.Next.Next.Next = b
	b.Next.Next = a.Next.Next

	type args struct {
		pHead *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test",
			args: args{
				pHead: a,
			},
			want: a.Next.Next,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EntryNodeOfLoop(tt.args.pHead); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EntryNodeOfLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindKthToTail(t *testing.T) {
	type args struct {
		pHead *ListNode
		k     int
	}

	testLink := NewSingleList([]int{1, 2, 3, 4})
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				pHead: NewSingleList([]int{1, 2}),
				k:     3,
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				pHead: testLink,
				k:     3,
			},
			want: testLink.Next,
		},
		{
			name: "test2",
			args: args{
				pHead: testLink,
				k:     4,
			},
			want: testLink,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthToTail(tt.args.pHead, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindKthToTail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteNode(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test-nil",
			args: args{
				head: nil,
				val:  0,
			},
			want: nil,
		},
		{
			name: "test-head",
			args: args{
				head: NewSingleList([]int{1, 2, 3}),
				val:  1,
			},
			want: NewSingleList([]int{2, 3}),
		},
		{
			name: "test-first",
			args: args{
				head: NewSingleList([]int{1}),
				val:  1,
			},
			want: nil,
		},
		{
			name: "test-middle",
			args: args{
				head: NewSingleList([]int{1, 2, 3}),
				val:  2,
			},
			want: NewSingleList([]int{1, 3}),
		},
		{
			name: "test-last",
			args: args{
				head: NewSingleList([]int{1, 2, 3}),
				val:  3,
			},
			want: NewSingleList([]int{1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteDuplication(t *testing.T) {
	type args struct {
		pHead *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				pHead: nil,
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				pHead: NewSingleList([]int{1, 1}),
			},
			want: nil,
		},
		{
			name: "test3",
			args: args{
				pHead: NewSingleList([]int{1, 2, 2}),
			},
			want: NewSingleList([]int{1}),
		},
		{
			name: "test4",
			args: args{
				pHead: NewSingleList([]int{1, 2, 2, 3}),
			},
			want: NewSingleList([]int{1, 3}),
		},
		{
			name: "test5",
			args: args{
				pHead: NewSingleList([]int{1, 2, 3, 3, 4, 4, 5}),
			},
			want: NewSingleList([]int{1, 2, 5}),
		},
		{
			name: "test6",
			args: args{
				pHead: NewSingleList([]int{1, 1, 1, 1, 1, 1, 1}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplication(tt.args.pHead); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplication() = %v, want %v", got, tt.want)
			}
		})
	}
}
