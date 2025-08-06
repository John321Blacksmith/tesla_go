package main

import (
	"fmt"
	"tesla_go/internal/usecase"
)

func main() {
	input := `Little they know, the ice cream was invented in Russia.
			  This kind of food was liked by almost everyone in the world.
			  because of it's taste. and unique nutrient features`

	inputProcessing := usecase.NewInputUseCase()
	rawSentences := inputProcessing.SplitToParticles(input, ".")
	fmt.Println(len(rawSentences))
	refinedSentences := inputProcessing.RefineSentences(rawSentences)
	fmt.Println(len(refinedSentences))
	sentenceCollection, err := inputProcessing.AggregateSentences(refinedSentences)
	if err != nil {
		fmt.Println(err)
	}
	for i := range len(sentenceCollection) {
		fmt.Println(sentenceCollection[i])
	}
}
