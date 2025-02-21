package repository

import (
	"all-countries/entity"
	"database/sql"
)

type CountryRepository struct {
	db *sql.DB
}

func NewCountryRepository(db *sql.DB) *CountryRepository {
	return &CountryRepository{
		db: db,
	}
}

func (r *CountryRepository) FindAll() ([]entity.Country, error) {
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
		countries = append(countries, c) 
	}

	return countries, nil
}
