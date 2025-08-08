package usecase

import (
	"tesla_go/internal/domain"
)

type InputManagement interface {
	SplitToParticles(data string, sym string) []string
	RefineSentences(rawSentences []string) []string
	AggregateSentences(refinedSentences []string) ([]domain.Sentence, error)
}

type InputUseCase struct {
	behaviour InputManagement
}

func NewInputUseCase(behaviour InputManagement) InputUseCase {
	return InputUseCase{
		behaviour: behaviour,
	}
}

func (uc *InputUseCase) SplitToParticles(data string, sym string) []string {
	return uc.behaviour.SplitToParticles(data, sym)
}

func (uc *InputUseCase) RefineSentences(rawSentences []string) []string {
	return uc.behaviour.RefineSentences(rawSentences)
}

func (uc *InputUseCase) AggregateSentences(refinedSentences []string) ([]domain.Sentence, error) {
	return uc.behaviour.AggregateSentences(refinedSentences)
}
