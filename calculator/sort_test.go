package calculator

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Should sort numbers",
			args: args{
				numbers: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "Should sort mixed order numbers",
			args: args{
				numbers: []int{1, 11, 2, 9, 8},
			},
			want: []int{11, 9, 8, 2, 1},
		},
		{
			name: "Should handler duplicate numbers",
			args: args{
				numbers: []int{11, 11, 20, 9, 20, 8},
			},
			want: []int{20, 20, 11, 11, 9, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortDesc(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
