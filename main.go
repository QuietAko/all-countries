package main

import (
	"all-countries/cache"
	"all-countries/controller"
	"all-countries/db"
	"all-countries/repository"
	"all-countries/service"
	"log"
	"net/http"
)

func main() {
	// Подключаемся к базе данных
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer db.Close()

	countryRepo := repository.NewCountryRepository(db)
	countryService := service.NewCountryService(countryRepo, cache.GetRedisClient())
	countryController := controller.NewCountryController(countryService)

	http.HandleFunc("/api/country", countryController.GetAllCountries)

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
