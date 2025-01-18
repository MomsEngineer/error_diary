package main

import "math"

func findMinProduct(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	min1, min2, max1, max2 := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32

	for _, num := range nums {
		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}

		if max1 < num {
			max2 = max1
			max1 = num
		} else if max2 < num {
			max2 = num
		}
	}

	if min1 >= 0 {
		return min1 * min2
	} else if max1 < 0 {
		return max1 * max2
	}

	return min1 * max1
}
