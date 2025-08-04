package usecase

import (
	"tesla_go/internal/domain"
)

type InputManagement interface {
	ProcessInput(data []byte) string
	SplitToParticles(data string, sym string) []string
	RefineSentences(rawSentences []string) []string
	AggregateSentences(refinedSentences []string) []domain.Sentence
}

type InputUseCase struct {
	InputManagement
}

func NewInputUseCase(behaviour InputManagement) InputUseCase {
	return InputUseCase{}
}

func ProcessInput(data []byte) string {
	return ""
}

func SplitToParticles(data string, sym string) []string {
	return nil
}

func RefineSentences(rawSentences []string) []string {
	return nil
}

func AggregateSentences(refinedSentences []string) []domain.Sentence {
	return nil
}
