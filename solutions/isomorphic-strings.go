package main

import (
	"bytes"
)

func isIsomorphic(s string, t string) bool {

	var charMap = make(map[string]string)
	//fill the character hashMap
	for i, c := range s {
		character := string(c)

		values := mapKeysValues(charMap)
		if !contains(values, charMap[character]) && !contains(values, string(t[i])) {
			charMap[character] = string(t[i])
		}

	}

	var stringBuilder bytes.Buffer
	//recreate the second string
	for i, _ := range s {
		stringBuilder.WriteString(charMap[string(s[i])])
	}

	return stringBuilder.String() == t
}

func contains(s []string, e string) bool {

	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func mapKeysValues(m map[string]string) []string {
	values := make([]string, 0, 1)

	for _, value := range m {
		values = append(values, value)
	}

	return values
}
