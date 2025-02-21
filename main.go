package main

import (
	"all-countries/controller"
	"all-countries/db"
	"all-countries/repository"
	"all-countries/service"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Подключаемся к базе данных
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer db.Close()

	// Подключаемся к Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DSN"), // Адрес Redis-сервера
		Password: "",                     // Пароль (если есть)
		DB:       0,                      // Номер базы данных
	})

	// Проверяем подключение к Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		log.Fatalf("Ошибка при подключении к Redis: %v", err)
	}

	countryRepo := repository.NewCountryRepository(db, redisClient)
	countryService := service.NewCountryService(countryRepo)
	countryController := controller.NewCountryController(countryService)

	http.HandleFunc("api/countries", countryController.GetAllCountries)

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
