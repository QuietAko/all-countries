package service

import (
	"all-countries/entity"
	"all-countries/repository"
)

type CountryService struct {
	repo *repository.CountryRepository
}

func NewCountryService(repo *repository.CountryRepository) *CountryService {
	return &CountryService{repo: repo}
}

func (s *CountryService) GetAllCountries() ([]entity.Country, error) {
	return s.repo.FindAll()
}