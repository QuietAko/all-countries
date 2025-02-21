package controller

import (
	"encoding/json"
	"net/http"
	"all-countries/service"
)

// CountryController предоставляет методы для обработки HTTP-запросов, связанных с странами
type CountryController struct {
	service *service.CountryService
}

// NewCountryController создает новый экземпляр CountryController
func NewCountryController(service *service.CountryService) *CountryController {
	return &CountryController{service: service}
}

// GetAllCountries обрабатывает запрос на получение всех стран
func (c *CountryController) GetAllCountries(w http.ResponseWriter, r *http.Request) {
	countries, err := c.service.GetAllCountries()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}
