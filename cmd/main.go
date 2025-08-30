package main

import (
	"fmt"
	"strings"
	"tesla_go/pkg/classifier"
)

func main() {
	input := `Little they know, the ice cream was invented in Russia.
			  This kind of food was liked by almost everyone in the world.
			  because of it's taste. and unique nutrient features`

	inputManager := classifier.NewInputManager()
	rawSentences := inputManager.SplitIntoSentences(input)

	for i := range len(rawSentences) {
		fmt.Println("////")
		trimmedSentence := strings.TrimSpace(rawSentences[i])
		words := strings.Split(trimmedSentence, " ")
		refinedWords := []string{}
		for j := range len(words) {
			var word string
			var ind int
			for l := range len(words[j]) {
				letter := words[j][l]
				if !classifier.isAlpha(letter) {
					ind = strings.Index(words[j], string(letter))
				}
			}
			if ind == (len(words[j]) - 1) {
				word = words[j][:ind]
			} else {
				word = words[j]
			}
			refinedWords = append(refinedWords, word)
		}
		fmt.Println(refinedWords)
	}
}
