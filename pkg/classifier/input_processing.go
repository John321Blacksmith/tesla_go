package classifier

import (
	"fmt"
	"tesla_go/internal/domain"
)

// check the symbol
// if it's alphabetic
func is_alpha(sym byte) bool {
	return (sym >= 65 && sym <= 90) ||
		(sym >= 97 && sym <= 122)
}

// check each word for
// having a non alphabetic
// symbol
func RefineSentences(particles []string) []string {
	var refined_sentences []string
	if len(particles) != 0 {
		for i := range len(particles) {
			if len(particles[i]) != 0 {
				signs_count := 0
				for j := range len(particles[i]) {
					if !is_alpha(particles[i][j]) {
						signs_count += 1
					}
				}
				if signs_count == 0 {
					refined_sentences = append(refined_sentences, particles[i])
				}
			}
		}
	}
	return refined_sentences
}

// receive the data
// from the external client
// and split to sentences
func SplitToParticles(data string, sym string) []string {
	partiles := []string{}
	if data != "" {
		flag := 0
		for i := range len(data) {
			if string(data[i]) == sym {
				partiles = append(partiles, data[flag:i])
				flag = (i + 1)
			}
		}
		partiles = append(partiles, data[flag:])
	}
	return partiles
}

// format every string
// into the sentence
// object
func AggregateSentences(refinedSentences []string) ([]domain.Sentence, error) {
	var sentenceCollection []domain.Sentence
	if len(refinedSentences) != 0 {
		for i := range len(refinedSentences) {
			sentenceCollection = append(
				sentenceCollection,
				domain.Sentence{
					Category:    "unknown",
					LiteralData: SplitToParticles(refinedSentences[i], ""),
				},
			)
		}
		return sentenceCollection, nil
	}
	return nil, fmt.Errorf("no sentences received")
}
