package anagrams

import (
	"sort"
)

// GroupAnagrams will group all anagram words together
func GroupAnagrams(strs []string) [][]string {

	anagramsMp := make(map[string][]string)

	for _, str := range strs {
		chars := make([]byte, len(str))
		copy(chars, []byte(str))

		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		inStr := string(chars)

		anagrams, in := anagramsMp[inStr]

		if !in {
			anagrams = make([]string, 0, len(strs))
		}

		anagrams = append(anagrams, str)
		anagramsMp[inStr] = anagrams
	}

	res := make([][]string, 0, len(strs))

	for _, anagrams := range anagramsMp {
		res = append(res, anagrams)
	}

	return res
}
