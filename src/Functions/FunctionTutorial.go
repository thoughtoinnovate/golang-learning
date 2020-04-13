package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "go lang module"
	author := "ThoughToInnovate"

	fmt.Println(convert(module, author))

	//Variadic functions having n arguments
	bestScore := findBestDcore(100, 90, 80, 70, 500, 110, 220, 10, 30)
	fmt.Println("The best score is :", bestScore)

}

func convert(first, second string) (s1, s2 string) {
	first = strings.Title(first)
	second = strings.ToUpper(second)

	return first, second
}

func findBestDcore(scores ...int) (highestScore int) {
	bestScore := scores[0]
	for _, score := range scores {
		if bestScore < score {
			bestScore = score
		}
	}
	return bestScore
}
