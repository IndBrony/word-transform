package word

import (
	"math"
	"strings"
)

//CountTransformationStepWrapper is a wrapper for CountTransformationStep
//to place unnecesary input processing that has noting to do with the real algorithm
//param : (string) two words separated by a space, ex. "Bayam ayam"
func CountTransformationStepWrapper(input string) int {
	//to insure only uppercase letter
	input = strings.ToUpper(input)
	splitted := strings.Split(input, " ")
	//if it only one word then return the step required to remove the letter 1 by 1
	//else use the algorithm
	if len(splitted) == 1 {
		return len(input)
	}
	return CountTransformationStep(splitted[0], splitted[1])
}

//CountTransformationStep counts the step required to transform a word into another
//by removing/adding letter (swapping doesn't count)
//params:
//    from: (string) the origin string
//    to  : (string) the string after transformation
func CountTransformationStep(from, to string) int {
	//creates a map for each input to simplify the diff-ing process
	origin := ConvStringToLetterCountMap(from)
	goal := ConvStringToLetterCountMap(to)

	//diff with the "goal" map being the frame of reference
	//count the letters on "goal" that need tobe added/removed to/from the origin
	//and do it the other way around too for the second diff
	firstDiff, secondDiff := diffCount(goal, origin), diffCount(origin, goal)

	return int(firstDiff + math.Abs(firstDiff-secondDiff))
}

//ConvStringToLetterCountMap converts the string into a map of letters
func ConvStringToLetterCountMap(s string) (mapped map[string]float64) {
	mapped = make(map[string]float64)
	for _, letter := range strings.Split(s, "") {
		mapped[letter]++
	}
	return
}

//diffCount counts the difference between base and the target
//using base as the frame of reference
//ex. base{a=1,b=2,c=2}, target{a=2,c=1,d=2}
//    start
//        remove 1 "a" from target
//        add 2 "b" to target
//        add 1 "c" to target
//    end
//base doesn't have any knowledge of "d"
func diffCount(base, target map[string]float64) (count float64) {
	for letter, letterCount := range base {
		if target[letter] != letterCount {
			count += math.Abs(letterCount - target[letter])
			target[letter] = 0
		}
	}
	return
}
