package collectionsutils

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		list []string
		f    func(string) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "bidule",
			args: args{
				list: []string{
					"a", "aa", "aaa",
				},
				f: func(s string) int {
					return len(s)
				},
			},
			want: []int{1,2,3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.list, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestReduce(t *testing.T) {
// 	type args struct {
// 		list []A
// 		init B
// 		f    func(B, A) B
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want B
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Reduce(tt.args.list, tt.args.init, tt.args.f); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Reduce() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestFilter(t *testing.T) {
// 	type args struct {
// 		list []A
// 		f    func(A) bool
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []A
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Filter(tt.args.list, tt.args.f); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Filter() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestIntersection(t *testing.T) {
// 	type args struct {
// 		a []A
// 		b []A
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []A
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Intersection(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Intersection() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestContains(t *testing.T) {
// 	type args struct {
// 		list []A
// 		elem A
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Contains(tt.args.list, tt.args.elem); got != tt.want {
// 				t.Errorf("Contains() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
