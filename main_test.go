package main

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		args    [][]int
		name    string
		wantErr bool
		want    [][]int
	}{
		{
			args:    [][]int{{3, 4}, {1, 2}},
			name:    "Non overlapping merge",
			wantErr: false,
			want:    [][]int{{1, 2}, {3, 4}},
		},

		{
			args:    [][]int{{1, 2}, {2, 3}},
			name:    "Non overlapping successive",
			wantErr: false,
			want:    [][]int{{1, 2}, {2, 3}},
		},

		{
			args:    [][]int{{2, 5}, {1, 3}},
			name:    "Overlapping merge",
			wantErr: false,
			want:    [][]int{{1, 5}},
		},

		{
			args:    [][]int{{1, 5}, {1, 3}},
			name:    "Encasing merge",
			wantErr: false,
			want:    [][]int{{1, 5}},
		},

		{
			args:    [][]int{{1, 5}, {1, 5}},
			name:    "Duplicate merge",
			wantErr: false,
			want:    [][]int{{1, 5}},
		},

		{
			args:    [][]int{{1, 2}, {2, 2}},
			name:    "Duplicate merge",
			wantErr: false,
			want:    [][]int{{1, 2}, {2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mergeIntervals(tt.args[0], tt.args[1])

			if (err != nil) != tt.wantErr {
				t.Errorf("unexpected error: %s", err.Error())
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got != want: got = %+v, want = %+v", got, tt.want)
			}
		})
	}

}

func TestMERGE(t *testing.T) {
	tests := []struct {
		name    string
		args    [][]int
		want    [][]int
		wantErr bool
	}{
		{
			name:    "assignment example",
			args:    [][]int{{25, 30}, {2, 19}, {14, 23}, {4, 8}},
			want:    [][]int{{2, 23}, {25, 30}},
			wantErr: false,
		},
		{
			name:    "ascending order",
			args:    [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {6, 6}, {6, 7}},
			want:    [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {6, 7}},
			wantErr: false,
		},
		{
			name:    "descending order",
			args:    [][]int{{6, 6}, {4, 5}, {3, 4}, {2, 3}, {1, 2}},
			want:    [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {6, 6}},
			wantErr: false,
		},
		{
			name:    "duplicates",
			args:    [][]int{{1, 2}, {1, 2}, {3, 4}, {3, 4}},
			want:    [][]int{{1, 2}, {3, 4}},
			wantErr: false,
		},
		{
			name:    "overlapping duplicates",
			args:    [][]int{{1, 4}, {1, 4}, {3, 5}, {3, 5}},
			want:    [][]int{{1, 5}},
			wantErr: false,
		},
		{
			name:    "Successive merge",
			args:    [][]int{{1, 4}, {6, 8}, {3, 7}},
			want:    [][]int{{1, 8}},
			wantErr: false,
		},
		{
			name:    "Everything together",
			args:    [][]int{{1, 4}, {1, 4}, {6, 8}, {7, 10}, {12, 14}, {15, 18}, {19, 22}, {23, 25}, {3, 7}, {11, 26}, {1, 3}, {3, 4}},
			want:    [][]int{{1, 10}, {11, 26}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MERGE(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("MERGE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MERGE() got = %v, want %v", got, tt.want)
			}
		})
	}
}
