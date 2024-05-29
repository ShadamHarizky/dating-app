package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func Explode(s, separator string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, separator)
}

func ExplodeToInt(s, seperator string) []int {
	strSlice := Explode(s, seperator)
	intSlice := make([]int, len(strSlice))

	for i, str := range strSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Error converting element %s: %v\n", str, err)
			continue
		}
		intSlice[i] = num
	}

	return intSlice
}

func StrSliceToInt(strSlice []string) []int {
	intSlice := make([]int, len(strSlice))

	for i, str := range strSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Error converting element %s: %v\n", str, err)
			continue
		}
		intSlice[i] = num
	}

	return intSlice
}
