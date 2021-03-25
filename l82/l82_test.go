package l82

import (
	"probestar.com/leet-go/list_node"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *list_node.ListNode
	}
	tests := []struct {
		name string
		args args
		want *list_node.ListNode
	}{
		{"1", args{list_node.GenerateList([]int{1, 2, 3, 3, 4, 4, 5})}, list_node.GenerateList([]int{1, 2, 5})},
		{"2", args{list_node.GenerateList([]int{1, 1, 1, 2, 3})}, list_node.GenerateList([]int{2, 3})},
		{"3", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
