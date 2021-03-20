package l92

import (
	"probestar.com/leet-go/list_node"
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *list_node.ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *list_node.ListNode
	}{
		//{"1", args{generateList([]int{1, 2, 3, 4, 5}), 2, 4}, generateList([]int{1, 4, 3, 2, 5})},
		//{"2", args{generateList([]int{5}), 1, 1}, generateList([]int{5})},
		{"3", args{list_node.GenerateList([]int{1, 2, 3, 4, 5}), 1, 5}, list_node.GenerateList([]int{5, 4, 3, 2, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
