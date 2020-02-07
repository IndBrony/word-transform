package word

import (
	"math"
	"strings"
)

func CountTransformationStep(from, to string) int {
	origin := ConvStringToLetterCountMap(from)
	goal := ConvStringToLetterCountMap(to)
	var firstDiff, secondDiff float64
	for letter, letterCount := range goal {
		if origin[letter] != letterCount {
			firstDiff += math.Abs(letterCount - origin[letter])
			origin[letter] = 0
		}
	}
	for letter, letterCount := range origin {
		if goal[letter] != letterCount {
			secondDiff += math.Abs(letterCount - goal[letter])
			goal[letter] = 0
		}
	}
	return int(firstDiff + math.Abs(firstDiff-secondDiff))
}

func ConvStringToLetterCountMap(s string) map[string]float64 {
	mapped := make(map[string]float64)
	for _, letter := range strings.Split(s, "") {
		mapped[letter]++
	}
	return mapped
}
