package service

import (
	"app-destination/model"
	"app-destination/repository"
)

type DestinationService struct {
	Repo repository.DestinationRepository
}

func NewDestinationService(repo repository.DestinationRepository) *DestinationService {
	return &DestinationService{Repo: repo}
}

func (s *DestinationService) GetAllDestinations() ([]model.Destination, error) {
	return s.Repo.GetAllDestinations()
}

func (s *DestinationService) CreateDestination(d model.Destination) error {
	return s.Repo.InsertDestination(d)
}