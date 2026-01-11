package main

import (
	"reflect"
	"testing"
)

func TestRollingSum(t *testing.T) {
	tests := []struct {
		name     string
		data     []bool
		window   int
		expected []int
	}{
		{
			name:     "simple case with window 3",
			data:     []bool{true, false, true, true, false},
			window:   3,
			expected: []int{1, 2, 2, 2, 1},
		},
		{
			name:     "all true values",
			data:     []bool{true, true, true, true},
			window:   3,
			expected: []int{2, 3, 3, 2},
		},
		{
			name:     "all false values",
			data:     []bool{false, false, false, false},
			window:   3,
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "window size 1",
			data:     []bool{true, false, true, false},
			window:   1,
			expected: []int{1, 0, 1, 0},
		},
		{
			name:     "window size equals data length",
			data:     []bool{true, true, false, true},
			window:   4,
			expected: []int{2, 2, 3, 2},
		},
		{
			name:     "window size 2",
			data:     []bool{true, false, true, true, false},
			window:   2,
			expected: []int{1, 1, 1, 2, 1},
		},
		{
			name:     "single element",
			data:     []bool{true},
			window:   1,
			expected: []int{1},
		},
		{
			name:     "alternating pattern",
			data:     []bool{true, false, true, false, true, false},
			window:   3,
			expected: []int{1, 2, 1, 2, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RollingSumBool(tt.data, tt.window)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RollingSum(%v, %d) = %v, want %v", tt.data, tt.window, result, tt.expected)
			}
		})
	}
}

func TestRollingSumEdgeCases(t *testing.T) {
	t.Run("empty data", func(t *testing.T) {
		result := RollingSumBool([]bool{}, 1)
		expected := []int{}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("RollingSum([], 1) = %v, want %v", result, expected)
		}
	})
}
