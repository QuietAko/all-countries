package repository

import (
	"all-countries/entity"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"
	"github.com/go-redis/redis/v8"
)

type CountryRepository struct {
	db    *sql.DB
	redis *redis.Client
}

func NewCountryRepository(db *sql.DB, redis *redis.Client) *CountryRepository {
	return &CountryRepository{
		db:    db,
		redis: redis,
	}
}

func (r *CountryRepository) FindAll() ([]entity.Country, error) {
	ctx := context.Background()

	cached, err := r.redis.Get(ctx, "countries").Result()
	if err == nil {
		var countries []entity.Country
		if err := json.Unmarshal([]byte(cached), &countries); err == nil {
			log.Println("Cache hit")
			return countries, nil
		}
	}

	rows, err := r.db.Query("SELECT id, name FROM country")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var countries []entity.Country
	for rows.Next() {
		var c entity.Country
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		countries = append(countries, c) // Используем значение, а не указатель
	}

	// Сохраняем данные в Redis
	data, err := json.Marshal(countries)
	if err != nil {
		return nil, err
	}
	r.redis.Set(ctx, "countries", data, 5*time.Minute)

	return countries, nil
}
