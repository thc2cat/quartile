package main

import (
	"reflect"
	"testing"
)

func Test_quartileCalcf32(t *testing.T) {
	tests := []struct {
		name string
		args []float32
		want [3]float32
	}{
		{"basic", []float32{10, 25, 30, 40, 41, 42, 50, 55, 70, 101, 110, 111}, [3]float32{46, 40, 70}},
		{"empty", []float32{}, [3]float32{-1, 0, 0}},
		{"not enought value", []float32{5}, [3]float32{-1, 0, 0}},
		{"medianeven", []float32{1, 4, 6, 7}, [3]float32{5, 4, 6}},
		{"medianodd1", []float32{1., 4, 6, 7, 9}, [3]float32{6, 4, 7}},
		{"medianodd2", []float32{1, 4, 6, 7, 9, 10, 211}, [3]float32{7, 4, 10}},
		{"khan", []float32{25, 28, 29, 39, 30, 34, 35, 35, 37, 38}, [3]float32{34.5, 29, 37}},
		{"khan1", []float32{5, 7, 10, 15, 19, 21, 21, 22, 22, 23, 23, 23, 23, 23, 24, 24, 24, 24, 25}, [3]float32{23, 19, 24}},
		{"test2", []float32{1, 11, 15, 19, 20, 24, 28, 34, 37, 47, 50, 61, 68}, [3]float32{28, 19, 47}},
		{"some values", []float32{3, 13, 28, 31, 37, 50, 57, 62, 78, 79, 81, 83}, [3]float32{53.5, 31, 78}},
	}

	// Exemple :
	// set :  3, 13, 28, 31, 37, 50, 57, 62, 78, 79, 81, 83 => 12 valeurs
	// even size set => 12/2 = 6 : Sum value(6,7)/2 => (50+57)/2 => 107/2 = 53.5 == Median
	// Q1 = 12/4 = 4 => 4 eme valeur => 31 = Q1
	// Q2 = 12*3/4   => 9 eme valeur => 78 = Q2
	// Expect [3]float32{53.5, 31, 78} as result
	// Online Verifications :
	// https://www.calculatorsoup.com/calculators/statistics/quartile-calculator.php

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quartileCalcf32(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quartileCalc() %v\n= %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
