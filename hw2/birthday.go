package hw2

// HasRepeat returns true if any of the elements in the slice are the same
func HasRepeat(is []int) bool {
	for i := 0; i < len(is)-1; i++ {
		jc := is[i]
		for j := i + 1; j < len(is); j++ {
			if jc == is[j] {
				return true
			}
		}

	}
	return false
}

