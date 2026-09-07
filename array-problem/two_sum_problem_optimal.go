package arrayproblem

import (
	"github.com/munnaMia/Data-Structure-Algorithms/sorting"
)

// using hash map to solve the two sum problem
// time complexity is O(n)
func TwoSumBest(array []int, target int) (bool, int, int) {
	sorting.Quicksort(0, len(array)-1, array)

	L, R := 0, len(array)-1

	for L < R {
		sum := array[L] + array[R]
		if sum == target {
			return true, L, R
		} else if sum > target {
			R--
		} else {
			L++
		}
	}
	return false, -1, -1

}
