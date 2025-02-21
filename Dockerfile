FROM golang:1.22-alpine AS builder

# Устанавливаем migrate через go install
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Устанавливаем команду по умолчанию для запуска Go приложения
CMD ["go", "run", "/app/main.go"]
