package leet

import (
	"reflect"
	"testing"
)

func Test_generateList(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"0", args{[]int{}}, nil},
		{"1", args{[]int{1}}, generateLitNode(1)},
		{"2", args{[]int{1, 2}}, generateLitNodeWithNext(1, generateLitNode(2))},
		{"3", args{[]int{1, 2, 3}}, generateLitNodeWithNext(1, generateLitNodeWithNext(2, generateLitNode(3)))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateList(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
