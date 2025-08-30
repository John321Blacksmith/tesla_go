package usecase

import (
	"tesla_go/internal/domain"
)

type InputManagement interface {
	SplitIntoParticles(data string, sym string) []string
	RefineSentences(rawSentences []string) [][]string
	AggregateSentences(refinedSentences [][]string) ([]domain.Sentence, error)
}

type InputUseCase struct {
	behaviour InputManagement
}

func NewInputUseCase(actor InputManagement) *InputUseCase {
	return &InputUseCase{
		behaviour: actor,
	}
}

func (uc *InputUseCase) SplitToParticles(data string, sym string) []string {
	return uc.behaviour.SplitIntoParticles(data, sym)
}

func (uc *InputUseCase) RefineSentences(rawSentences []string) []string {
	return uc.behaviour.RefineSentences(rawSentences)
}

func (uc *InputUseCase) AggregateSentences(refinedSentences []string) ([]domain.Sentence, error) {
	return uc.behaviour.AggregateSentences(refinedSentences)
}
