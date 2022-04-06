package matchingMethod

func KMPMatch(text, pattern string) int {

	runeText := []rune(text)
	runePattern := []rune(pattern)

	var tLength int = len(runeText)
	var pLength int = len(runePattern)

	var fail = computeFail(runePattern, pLength)

	var i int = 0
	var j int = 0

	for i < tLength {
		if string(runePattern[j]) == string(runeText[i]) {

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

func computeFail(pattern []rune, length int) []int {
	var fail = make([]int, length)
	fail[0] = 0

	var j int = 0
	var i int = 1

	for i < length {
		if string(pattern[j]) == string(pattern[i]) {
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

	for i := 0; i < length; i++ {
		print(fail[i])
	}

	return fail
}
