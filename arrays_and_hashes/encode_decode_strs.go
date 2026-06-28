package arraysandhashes

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {

	if len(strs) == 1 && strs[0] == "" {
		fmt.Println("True ")
		return "101109112116121"
	}

	strAscii := []string{}

	for i := range strs {

		str := strs[i]

		for j := range str {

			byteOfStr := str[j]

			asciiInt := int(byteOfStr)

			strAscii = append(strAscii, strconv.Itoa(asciiInt))

			if j != len(str)-1 {
				strAscii = append(strAscii, "/")
			}

		}

		if i != len(strs)-1 {
			strAscii = append(strAscii, "+")
		}

	}

	return strings.Join(strAscii, "")
}

func (s *Solution) Decode(encoded string) []string {

	if encoded == "101109112116121" {
		return []string{""}
	} else if len(encoded) == 0 {
		return []string{}
	}

	asciis := strings.Split(encoded, "+")

	decoded := []string{}
	for i := range asciis {

		eachAscii := asciis[i]
		splitted := strings.Split(eachAscii, "/")

		strB := strings.Builder{}

		for j := range splitted {

			eachSplit := splitted[j]

			intOfStr, _ := strconv.Atoi(eachSplit)

			if intOfStr != 0 {

				strB.WriteByte(byte(intOfStr))
			}

		}

		if len(strB.String()) != 0 {
			decoded = append(decoded, strB.String())
		} else {
			decoded = append(decoded, "")
		}

	}

	return decoded
}
