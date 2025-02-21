package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// Структура для отображения страны
type Country struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// Подключение к базе данных
func getDBConnection() (*sql.DB, error) {
	// Указываем параметры подключения к PostgreSQL
	dsn := "host=db user=user password=password dbname=mydb port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Функция для получения списка стран
func getCountries(w http.ResponseWriter, r *http.Request) {
	// Подключаемся к базе данных
	db, err := getDBConnection()
	if err != nil {
		http.Error(w, "Не удалось подключиться к базе данных", http.StatusInternalServerError)
		log.Println("Ошибка подключения к базе данных: ", err)
		return
	}
	defer db.Close()

	// Запрашиваем все страны из базы данных
	rows, err := db.Query("SELECT id, name FROM country")
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса", http.StatusInternalServerError)
		log.Println("Ошибка запроса: ", err)
		return
	}
	defer rows.Close()

	// Массив для хранения стран
	var countries []Country

	// Чтение данных из результата запроса
	for rows.Next() {
		var country Country
		if err := rows.Scan(&country.ID, &country.Name); err != nil {
			http.Error(w, "Ошибка при чтении данных из базы данных", http.StatusInternalServerError)
			log.Println("Ошибка при чтении данных: ", err)
			return
		}
		countries = append(countries, country)
	}

	// Проверка на ошибки при итерации по строкам
	if err := rows.Err(); err != nil {
		http.Error(w, "Ошибка при обработке данных", http.StatusInternalServerError)
		log.Println("Ошибка обработки данных: ", err)
		return
	}

	// Преобразуем список стран в JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(countries); err != nil {
		http.Error(w, "Ошибка при сериализации данных в JSON", http.StatusInternalServerError)
		log.Println("Ошибка сериализации JSON: ", err)
	}
}

func main() {
	// Определение маршрута для получения списка стран
	http.HandleFunc("/api/countries", getCountries)

	// Запуск HTTP сервера на порту 8080
	fmt.Println("Запуск сервера на порту 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
}
