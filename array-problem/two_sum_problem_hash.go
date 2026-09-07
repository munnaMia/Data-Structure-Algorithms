package arrayproblem

// using hash map to solve the two sum problem
// time complexity is O(n) + map time complexity for order O(logn) for underOrder O(n)
func TwoSumHash(array []int, target int) (bool, int, int) {
	hash := make(map[int]int) // elem -> idx

	for i, v := range array {
		if _, ok := hash[target-v]; ok {
			return true, hash[target-v], i
		}
		hash[v] = i
	}
	return false, -1, -1
}
