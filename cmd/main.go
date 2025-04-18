package main

import (
	"fmt"
	"tesla_go/internal/engine"
)

func main() {
	input := `Little they know, the ice cream was invented in Russia
			  This kind of food was liked by almost everyone in the world
			  because of it's taste and unique nutrient features`

	// var dataset []engine.Category = engine.GetDataset("/home/ivankuzmichev/exp/tesla_go/categories.json")

	sentences := engine.Split_to_chunks(input, ".")

	for i := range len(sentences) {
		fmt.Printf("%v\n", sentences[i])
	}
}
