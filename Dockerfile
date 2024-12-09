# Используем официальный образ Golang как базовый
FROM golang:1.22.7-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum, и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код приложения
COPY . .

# Устанавливаем рабочую директорию
WORKDIR /app/cmd/ShtrafovNet

# Собираем приложение
RUN go build -o app/main .

# Запускаем приложение
CMD ["app/main"]
