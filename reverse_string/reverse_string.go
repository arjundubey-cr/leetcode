package leetcode

func reverseString(s []byte) {
	initial := 0
	final := len(s) - 1
	for initial < final {
		temp := s[initial]
		s[initial] = s[final]
		s[final] = temp
		initial += 1
		final -= 1
	}
}
