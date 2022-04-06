package matchingMethod

func KMPMatch(text, pattern string) int {

	var tLength int = len(text)
	var pLength int = len(pattern)

	var fail = computeFail(pattern, pLength)

	var j int = 0
	var i int = 0

	for i < tLength {
		if pattern[j] == text[i] {
			if j == pLength-1 {
				return i - pLength + 1
			}
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			i++
		}
	}
	return -1
}

func computeFail(pattern string, length int) []int {
	var fail = make([]int, length)
	fail[0] = 0

	var j int = 0
	var i int = 1

	for i < length {
		if pattern[j] == pattern[i] {
			fail[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			fail[i] = 0
			i++
		}
	}
	return fail
}
