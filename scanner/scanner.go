// Package scanner checks whether a game configuration is legal or not
package scanner

// Returns whether the configuration is legal or not
//
// Implements a O(n*n) complexity check as described in the paper
// "Notes on the 15 puzzle" by Wm. Woolsey Johnson
func IsLegal(size int, values []int) bool {
  var inversions int
  n := len(values)

  for i := 0; i < (n - 1); i++ {
    if ((values[i] != 0) && (values[i] != 1)) {
      for j := i+1; j < n; j++ {
        if ((values[j] != 0) && (values[i] > values[j])) {
          inversions += 1
        }
      }
    }
  }

  if ((size%2 == 1) && (inversions%2 == 0)) {
    return true
  } else {
    return false
  }
}
