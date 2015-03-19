// Package scanner checks whether a game configuration is legal or not
package scanner

// Returns zero-based index of the tile having value 0
// returns -1 if not found, but that won't be the case
func zeroIndex(values []int) int {
	for i, v := range values {
		if v == 0 {
			return i
		}
	}

	return -1
}

// IsLegal returns whether the configuration is legal or not
//
// Implements a O(n*n) complexity check as described in the paper
// "Notes on the 15 puzzle" by Wm. Woolsey Johnson
func IsLegal(size int, values []int) (bool, int) {
	var inversions int
	n := len(values)

	for i := 0; i < (n - 1); i++ {
		if (values[i] != 0) && (values[i] != 1) {
			for j := i + 1; j < n; j++ {
				if (values[j] != 0) && (values[i] > values[j]) {
					inversions++
				}
			}
		}
	}

	if (size%2 == 1) && (inversions%2 == 0) {
		return true, zeroIndex(values)
	}
	return false, -1
}
