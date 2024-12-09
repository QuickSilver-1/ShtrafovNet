package config

import (
	"auction/internal/application/logger"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
    AppConfig = NewConfig()
)

// Config представляет конфигурацию приложения
type Config struct {
    HttpPort int        // Порт для запуска приложения на HTTP
    GrpcPort int        // Порт для запуска приложения на gRPC
    PgHost   string     // Хост для базы данных PostgreSQL
    PgPort   int        // Порт для базы данных PostgreSQL
    PgName   string     // Имя базы данных PostgreSQL
    PgUser   string     // Имя пользователя для базы данных PostgreSQL
    PgPass   string     // Пароль для базы данных PostgreSQL
    SecretKey []byte    // Секретный код для JWT
    SmtpHost string     // Хост для отправки сообщений через почтовый сервис
    SmtpPort int        // Порт для отправки сообщений через почтовый сервис
    MailUser string     // Корпоративная почта
    MailPass string     // Пароль к корпоративной почте
}

func NewConfig() *Config {
    logger.Log.Info("collect config")
    // Загрузка переменных окружения из .env файла
    err := godotenv.Load("config.env")

    if err != nil {
        panic(fmt.Sprintf(`Error loading configuration file: %v`, err))
    }

    port, err := strconv.Atoi(os.Getenv("DB_PORT"))

    if err != nil {
        panic(`Invalid "DB_PORT"`)
    }

    smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))

    if err != nil {
        panic(`Invalid "SMTP_PORT"`)
    }

    httpPort, err := strconv.Atoi(os.Getenv("HTTP_PORT"))

    if err != nil {
        panic(`Invalid "HTTP_PORT"`)
    }

    grpcPort, err := strconv.Atoi(os.Getenv("GRPC_PORT"))

    if err != nil {
        panic(`Invalid "GRPC_PORT"`)
    }

    config := &Config{
        HttpPort: httpPort,                     
        GrpcPort: grpcPort,                      
        PgHost:   os.Getenv("DB_HOST"),          
        PgPort:   port,                          
        PgName:   os.Getenv("DB_NAME"),          
        PgUser:   os.Getenv("DB_USER"),          
        PgPass:   os.Getenv("DB_PASSWORD"),      
        SecretKey: []byte(os.Getenv("SECRET_KEY")),
        SmtpHost: os.Getenv("SMTP_HOST"),        
        SmtpPort: smtpPort,                      
        MailUser: os.Getenv("MAIL_USER"),        
        MailPass: os.Getenv("MAIL_PASS"),        
    }

    return config
}