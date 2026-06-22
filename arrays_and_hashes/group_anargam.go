package arraysandhashes

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {

	anargamsMap := map[string][]string{}
	anargamsSlices := [][]string{}

	for i := range strs {

		str := string(strs[i])

		splittedStrs := strings.Split(str, "")

		sort.Strings(splittedStrs)

		joinedStr := strings.Join(splittedStrs, "")

		insideSlices, _ := anargamsMap[joinedStr]

		insideSlices = append(insideSlices, str)

		anargamsMap[joinedStr] = insideSlices

	}

	for _, value := range anargamsMap {
		anargamsSlices = append(anargamsSlices, value)
	}

	return anargamsSlices
}
