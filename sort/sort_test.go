package sort

import (
	"reflect"
	"testing"
)

func Test_bubble(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
		{"2", args{[]int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}}, []int{2, 3, 4, 5, 15, 19, 26, 27, 36, 38, 44, 46, 47, 48, 50}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bubble(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selection(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
		{"2", args{[]int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}}, []int{2, 3, 4, 5, 15, 19, 26, 27, 36, 38, 44, 46, 47, 48, 50}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selection(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertion(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
		{"2", args{[]int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}}, []int{2, 3, 4, 5, 15, 19, 26, 27, 36, 38, 44, 46, 47, 48, 50}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertion(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shell(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
		{"2", args{[]int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}}, []int{2, 3, 4, 5, 15, 19, 26, 27, 36, 38, 44, 46, 47, 48, 50}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shell(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shell() = %v, want %v", got, tt.want)
			}
		})
	}
}
