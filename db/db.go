package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Connect() (*sql.DB, error) {
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		return nil, fmt.Errorf("переменная окружения DATABASE_URL не установлена")
	}

	db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к базе данных: %v", err)
	}

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка при проверке подключения: %v", err)
	}

	log.Println("Успешное подключение к базе данных")
	return db, nil
}
