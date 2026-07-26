package hash

// basic hash array. return a hasharray
func Hasharray(maxElementValue int, arr []int) []int {
	hashArray := make([]int, maxElementValue)

	for _, v := range arr {
		hashArray[v]++
	}

	return hashArray
}

// checking on hash array
func CheckHashValue(target int, hashArr []int) bool {
	if target >= len(hashArr) && target < 0 {
		return false
	}
	return hashArr[target] > 0
}

// basic hash map. return a maping
func Hashmap(value any, arr []int) map[any]int {
	mapHash := make(map[any]int)

	for _, v := range arr {
		mapHash[v]++
	}

	return mapHash
}
