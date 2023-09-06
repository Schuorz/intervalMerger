package intervalparser

import (
	"reflect"
	"testing"
)

func TestParseIntervalsFromString(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    [][]int
		wantErr bool
	}{
		{
			name:    "assignment conversion",
			args:    "[25,30], [2,19], [14, 23], [4,8]",
			want:    [][]int{{25, 30}, {2, 19}, {14, 23}, {4, 8}},
			wantErr: false,
		},
		{
			name:    "assignment conversion no commas",
			args:    "[25,30] [2,19] [14, 23] [4,8]",
			want:    [][]int{{25, 30}, {2, 19}, {14, 23}, {4, 8}},
			wantErr: false,
		},
		{
			name:    "assignment conversion no seperators between intervals",
			args:    "[25,30][2,19][14, 23][4,8]",
			want:    [][]int{{25, 30}, {2, 19}, {14, 23}, {4, 8}},
			wantErr: false,
		},
		{
			name:    "assignment conversion no whitespaces",
			args:    "[25,30],[2,19],[14,23],[4,8]",
			want:    [][]int{{25, 30}, {2, 19}, {14, 23}, {4, 8}},
			wantErr: false,
		},
		{
			name:    "assignment conversion more whitespaces",
			args:    "[2 5, 3 0] ,[ 2 , 1 9 ],[ 1 4,2 3],[4 , 8 ]",
			want:    [][]int{{25, 30}, {2, 19}, {14, 23}, {4, 8}},
			wantErr: false,
		},
		{
			name:    "broken interval string missing number",
			args:    "[25,],[2,19],[14,23],[4,8]",
			wantErr: true,
		},
		{
			name:    "broken interval string missing bracket",
			args:    "[25,30],[2,19],[14,23,[4,8]",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseIntervalsFromString(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseIntervalsFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseIntervalsFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
