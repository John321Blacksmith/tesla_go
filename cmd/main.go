package main

import (
	"fmt"
	"tesla_go/pkg/classifier"
)

func main() {
	input := `Little/ they/ know. the ice cream. was invented. in Russia.
			  This kind/ of. food. was. liked. by almost everyone/ in the world.
			  because of it's taste/. and unique, nutrient/ features/`

	inputManager := classifier.NewInputManager()
	rawSentences := inputManager.SplitToParticles(input)
	refinedSentences := inputManager.RefineSentences(rawSentences)
	fmt.Println(refinedSentences)
	for ind, v := range refinedSentences {
		fmt.Println(ind, v)
	}

	preparedSentences := inputManager.AggregateSentences(refinedSentences)
	fmt.Println(preparedSentences)
	for ind, v := range preparedSentences {
		fmt.Println(ind, v)
	}
}
