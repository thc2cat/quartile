package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_sortints(t *testing.T) {
	tests := []struct {
		name string
		args []int32
		want []int32
	}{
		{"empty", []int32{}, []int32{}},
		{"one value", []int32{0}, []int32{0}},
		{"some values", []int32{1, 4, 7, 2, 6}, []int32{1, 2, 4, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortints(tt.args); !reflect.DeepEqual(got, tt.want) {
				// if got := sortints(tt.args); got != tt.want {
				t.Errorf("sortints() = %v, want %v", got, tt.want)
			}

			if len(tt.args) > 2 && reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("sortints() destroy original data %v %v", &tt.args, &tt.want)
			}
		})
	}
}

func Test_quartileDeviantPrint(t *testing.T) {
	if os.Getenv("DO_COLORTEST") != "YES" {
		t.Skip("Skipping testing colors")
	}
	quartileDeviantPrint([]int32{-10, 25, 30, 40, 41, 42, 50, 55, 70, 101, 500, 1110})
}
func Test_quartileCalc(t *testing.T) {
	tests := []struct {
		name string
		args []int32
		want [3]float32
	}{
		// {"empty", []int32{}, [3]float32{-1, 0, 0}},
		// {"not enought value", []int32{5}, [3]float32{-1, 0, 0}},
		// {"medianeven", []int32{1, 4, 6, 7}, [3]float32{5, 0, 0}},
		// {"medianodd", []int32{1, 4, 6, 7, 9}, [3]float32{6, 0, 0}},
		// {"medianodd", []int32{1, 4, 6, 7, 9, 10, 211}, [3]float32{7, 4, 10}},
		{"khan", []int32{25, 28, 29, 39, 30, 34, 35, 35, 37, 38}, [3]float32{32, 29, 35}},
		{"khan1", []int32{5, 7, 10, 15, 19, 21, 21, 22, 22, 23, 23, 23, 23, 23, 24, 24, 24, 24, 25}, [3]float32{23, 19, 24}},
		{"test1", []int32{10, 25, 30, 40, 41, 42, 50, 55, 70, 101, 110, 111}, [3]float32{46, 40, 70}},
		{"test2", []int32{1, 11, 15, 19, 20, 24, 28, 34, 37, 47, 50, 61, 68}, [3]float32{28, 19, 47}},
		// {"some values", []int32{3, 13, 28, 31, 37, 50, 57, 62, 78, 79, 81, 83}, [3]float32{53.5, 28.75, 78.75}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quartileCalc(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quartileCalc() %v\n= %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
