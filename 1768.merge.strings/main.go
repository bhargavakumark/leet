package main

import "log"

func mergeAlternately(word1 string, word2 string) string {
	mergeString := ""
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			mergeString = mergeString + string(word1[i])
		}
		if i < len(word2) {
			mergeString = mergeString + string(word2[i])
		}
	}

	return mergeString
}

func main() {
	log.Printf("%s", mergeAlternately("abc", "pqr"))
	log.Printf("%s", mergeAlternately("ab", "pqrs"))
}
