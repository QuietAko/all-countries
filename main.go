package main

import (
	"all-countries/cache"
	"all-countries/controller"
	"all-countries/db"
	"all-countries/repository"
	"all-countries/service"
	"all-countries/metrics"
	"log"
	"net/http"
)

func main() {
	metrics.Init()

	db, err := db.Connect()
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer db.Close()

	countryRepo := repository.NewCountryRepository(db)
	countryService := service.NewCountryService(countryRepo, cache.GetRedisClient())
	countryController := controller.NewCountryController(countryService)

	http.Handle("/metrics", metrics.Handler())
	http.HandleFunc("/api/country", countryController.GetAllCountries)

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
