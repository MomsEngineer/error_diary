// Даны два монотонно возрастающих массива, нужно найти и вернуть новый массив
// с элементами первого массива, которых нет во втором.

package main

import (
	"reflect"
	"testing"
)

func TestFindDifference(t *testing.T) {
	tests := []struct {
		name     string
		arr1     []int
		arr2     []int
		expected []int
	}{
		{
			name:     "Non-overlapping arrays",
			arr1:     []int{1, 2, 3},
			arr2:     []int{4, 5, 6},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Fully overlapping arrays",
			arr1:     []int{1, 2, 3},
			arr2:     []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "Partial overlap",
			arr1:     []int{1, 2, 3, 4, 5},
			arr2:     []int{2, 4, 6},
			expected: []int{1, 3, 5},
		},
		{
			name:     "Empty first array",
			arr1:     []int{},
			arr2:     []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "Empty second array",
			arr1:     []int{1, 2, 3},
			arr2:     []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Both arrays empty",
			arr1:     []int{},
			arr2:     []int{},
			expected: []int{},
		},
		{
			name:     "Duplicate elements in first array",
			arr1:     []int{1, 1, 2, 2, 3, 3},
			arr2:     []int{2, 3},
			expected: []int{1, 1},
		},
		{
			name:     "Duplicate elements in second array",
			arr1:     []int{1, 2, 3},
			arr2:     []int{2, 2, 3, 3},
			expected: []int{1},
		},
		{
			name:     "No common elements, different ranges",
			arr1:     []int{1, 2, 3},
			arr2:     []int{10, 20, 30},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findDifference(tt.arr1, tt.arr2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func findDifference(nums1, nums2 []int) []int {
	// Решение
	return []int{}
}
