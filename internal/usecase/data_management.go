package usecase

import (
	"tesla_go/internal/domain"
)

type DataManagement interface {
	LoadDataSet(source string) []domain.Category
	UpdateDataSet(object map[string]domain.Category) error
}

type DataUseCase struct{}

func NewDataUseCase() DataUseCase {
	return DataUseCase{}
}

func LoadDataset(source string) []domain.Category {
	return nil
}

func UpdateDataset(object map[string]domain.Category) error {
	return nil
}
