package usecase

import (
	"biophilia/internal/domain/entities"
	"biophilia/internal/domain/interfaces"
)

// ProteinUsecase содержит логику работы с белками
type ProteinUsecase struct {
	repository interfaces.ProteinRepository
}

func NewProteinUsecase(repo interfaces.ProteinRepository) *ProteinUsecase {
	return &ProteinUsecase{repository: repo}
}

func (u *ProteinUsecase) AddProtein(protein entities.Protein) error {
	return u.repository.Add(protein)
}

func (u *ProteinUsecase) GetProteins() ([]entities.Protein, error) {
	return u.repository.GetAll()
}

func (u *ProteinUsecase) GetProteinByID(id string) (*entities.Protein, error) {
	return u.repository.GetByID(id)
}

func (u *ProteinUsecase) UpdateProtein(protein entities.Protein) error {
	return u.repository.Update(protein)
}

func (u *ProteinUsecase) DeleteProtein(id string) error {
	return u.repository.Delete(id)
}
