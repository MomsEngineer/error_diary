// Найти в массиве наименьшее произведение двух чисел.

package main

import (
	"testing"
)

func TestFindMinProduct(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "Positive numbers only",
			arr:      []int{1, 2, 3, 4, 5},
			expected: 2, // 1 * 2 (наименьшее произведение)
		},
		{
			name:     "Negative numbers only",
			arr:      []int{-5, -4, -3, -2, -1},
			expected: 2, // -1 * -2 (наименьшее произведение)
		},
		{
			name:     "Mixed positive and negative numbers",
			arr:      []int{-10, -3, 1, 2, 3},
			expected: -30, // -10 * 3 (наименьшее произведение)
		},
		{
			name:     "Array with zero",
			arr:      []int{-5, 0, 3, 4},
			expected: -20, // -5 * 4 (наименьшее произведение)
		},
		{
			name:     "Array with two elements",
			arr:      []int{2, 3},
			expected: 6, // 2 * 3
		},
		{
			name:     "Array with one element",
			arr:      []int{42},
			expected: 0, // Недостаточно элементов для произведения
		},
		{
			name:     "Empty array",
			arr:      []int{},
			expected: 0, // Нет элементов
		},
		{
			name:     "All elements are zero",
			arr:      []int{0, 0, 0, 0},
			expected: 0, // 0 * 0
		},
		{
			name:     "Large positive numbers",
			arr:      []int{1000000, 2000000, 3000000},
			expected: 2000000000000, // 1000000 * 2000000
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinProduct(tt.arr)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
