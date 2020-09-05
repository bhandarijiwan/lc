// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

package main

import (
	"fmt"
)

var letterForDigits = [][]byte{{}, {},
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	outputLength := 1
	if len(digits) < 1 {
		outputLength = 0
	}
	for i := 0; i < len(digits); i++ {
		outputLength *= len(letterForDigits[digits[i]-'0'])
	}
	outputBytes := make([][]byte, outputLength)
	output := make([]string, outputLength)
	for i := 0; i < len(outputBytes); i++ {
		outputBytes[i] = make([]byte, len(digits))
	}
	letterRepeatCount := outputLength
	for i := 0; i < len(digits); i++ {
		letters := letterForDigits[digits[i]-'0']
		l := len(letters)
		letterRepeatCount = letterRepeatCount / l
		for k := 0; k < outputLength; {
			for j := 0; j < l; j++ {
				for m := 0; m < letterRepeatCount; m++ {
					outputBytes[k][i] = letters[j]
					k++
				}
			}
		}
	}
	for i := 0; i < outputLength; i++ {
		output[i] = string(outputBytes[i])
	}
	return output
}

func main() {
	fmt.Println(len(letterCombinations("23")))
}
