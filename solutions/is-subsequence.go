package main

func isSubsequence(s string, t string) bool {
	hitCounter := 0
	if s == "" {
		return true
	}

	for _, c := range t {
		if string(s[hitCounter]) == string(c) {
			hitCounter++
		}
		if hitCounter == len(s) {
			return true
		}

	}

	return false

}
