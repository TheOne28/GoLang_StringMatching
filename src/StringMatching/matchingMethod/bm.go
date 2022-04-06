package matchingMethod

// BM Main function
func BMMatch(text, pattern string) int {
	var tLength int = len(text)
	var pLength int = len(pattern)

	var last = buildLast(pattern, pLength)

	var i int = pLength - 1

	if i > tLength-1 {
		return -1
	}

	var j int = pLength - 1

	for loop := true; loop; loop = i <= tLength-1 {
		if pattern[j] == text[i] {
			if j == 0 {
				return i
			} else {
				i--
				j--
			}
		} else {
			var lo int = last[text[i]]
			i = i + pLength - min(j, 1+lo)
			j = pLength - 1
		}
	}
	return -1
}

func buildLast(pattern string, length int) []int {
	var last = make([]int, 128)

	for i := 0; i < 128; i++ {
		last[i] = -1
	}

	for i := 0; i < length; i++ {
		last[pattern[i]] = i
	}

	return last
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
