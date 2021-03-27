package l61

import (
	"probestar.com/leet-go/list_node"
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *list_node.ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *list_node.ListNode
	}{
		{"1", args{head: list_node.GenerateList([]int{1, 2, 3, 4, 5}), k: 2}, list_node.GenerateList([]int{4, 5, 1, 2, 3})},
		{"2", args{head: list_node.GenerateList([]int{0, 1, 2}), k: 4}, list_node.GenerateList([]int{2, 0, 1})},
		{"3", args{head: nil, k: 0}, nil},
		{"4", args{head: list_node.GenerateList([]int{1, 2}), k: 0}, list_node.GenerateList([]int{1, 2})},
		{"4", args{head: list_node.GenerateList([]int{1, 2}), k: 1}, list_node.GenerateList([]int{2, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
