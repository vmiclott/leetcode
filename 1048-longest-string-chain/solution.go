package longeststringchain

import (
	"sort"
)

func longestStrChain(words []string) int {
	longestChain := 1

	dpMap := make(map[string]int)

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	for _, word := range words {
		dpMap[word] = 1
		for i := range word {
			predecessor := word[:i] + word[i+1:]
			predecessorChainLength := dpMap[predecessor]
			if predecessorChainLength > 0 && dpMap[word] < predecessorChainLength+1 {
				dpMap[word] = predecessorChainLength + 1
				if longestChain < predecessorChainLength+1 {
					longestChain = predecessorChainLength + 1
				}
			}
		}
	}

	return longestChain
}
