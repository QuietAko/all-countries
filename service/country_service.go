package service

import (
	"all-countries/entity"
	"all-countries/repository"
	"context"
	"encoding/json"
	"log"
	"time"
	"github.com/go-redis/redis/v8"
)

type CountryService struct {
	repo  *repository.CountryRepository
	redis *redis.Client
}

// NewCountryService теперь принимает redis.Client
func NewCountryService(repo *repository.CountryRepository, redis *redis.Client) *CountryService {
	return &CountryService{
		repo:  repo,
		redis: redis, // Инициализируем redis
	}
}

func (s *CountryService) GetAllCountries() ([]entity.Country, error) {
	ctx := context.Background()

	cached, err := s.redis.Get(ctx, "countries").Result()
	if err == nil {
		var countries []entity.Country
		if err := json.Unmarshal([]byte(cached), &countries); err == nil {
			log.Println("Cache hit")
			return countries, nil
		}
	} else if err != redis.Nil { 
		log.Println("Ошибка при получении данных из Redis:", err)
	}

	countries, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	// Сохраняем данные в кэш
	data, err := json.Marshal(countries)
	if err != nil {
		return nil, err
	}

	if err := s.redis.Set(ctx, "countries", data, 5*time.Minute).Err(); err != nil {
		log.Println("Ошибка при сохранении данных в Redis:", err)
	}

	return countries, nil
}