include .env

# Имя контейнера
CONTAINER_NAME=all-countries-app-1

# Команда для входа в контейнер
entry:
	docker exec -it $(CONTAINER_NAME) /bin/sh

# Команда для применения миграций внутри контейнера
migrate-up:
	@echo "Применение миграций..."
	@docker-compose exec -T app migrate -path ./db/migrations -database $(DATABASE_URL) up

# Команда для отката миграций внутри контейнера
migrate-down:
	@echo "Откат миграций..."
	@docker-compose exec -T app migrate -path ./db/migrations -database $(DATABASE_URL) down
