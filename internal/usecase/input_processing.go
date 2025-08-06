package usecase

import (
	"tesla_go/internal/domain"
	"tesla_go/pkg/classifier"
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

func NewInputUseCase() InputUseCase {
	return InputUseCase{}
}

func (uc *InputUseCase) SplitToParticles(data string, sym string) []string {
	return classifier.SplitToParticles(data, sym)
}

func (uc *InputUseCase) RefineSentences(rawSentences []string) []string {
	return classifier.RefineSentences(rawSentences)
}

func (uc *InputUseCase) AggregateSentences(refinedSentences []string) ([]domain.Sentence, error) {
	return classifier.AggregateSentences(refinedSentences)
}
