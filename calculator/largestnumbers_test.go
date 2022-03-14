package calculator

import (
	"reflect"
	"testing"
)

func TestCalculateLargestNumbers(t *testing.T) {
	type args struct {
		filePath string
		topX     int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "Should return top numbers",
			args: args{
				filePath: "../test1.txt",
				topX:     3,
			},
			want:    []int{5, 2, 0},
			wantErr: false,
		},
		{
			name: "Should throw errors",
			args: args{
				filePath: "../test4.txt",
				topX:     3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should handle duplicate numbers well",
			args: args{
				filePath: "../test3.txt",
				topX:     4,
			},
			want:    []int{38949, 23435, 11111, 11111},
			wantErr: false,
		},
		{
			name: "Should return all numbers when topX greater than length of array",
			args: args{
				filePath: "../test2.txt",
				topX:     5,
			},
			want:    []int{9, 3, 2, 1},
			wantErr: false,
		},
		{
			name: "Should return 0 when topX is less then 1",
			args: args{
				filePath: "../test2.txt",
				topX:     0,
			},
			want:    []int{},
			wantErr: false,
		},
		{
			name: "Should return 0 when topX is less then 1",
			args: args{
				filePath: "../test2.txt",
				topX:     -1,
			},
			want:    []int{},
			wantErr: false,
		},
		{
			name: "Should throw error when fail to parse file",
			args: args{
				filePath: "../not_exist.txt",
				topX:     3,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateLargestNumbers(tt.args.filePath, tt.args.topX)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateLargestNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("test CalculateLargestNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
