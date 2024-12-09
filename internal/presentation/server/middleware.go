package server

import (
	"auction/internal/application/config"
	"auction/internal/application/users"
	derr "auction/internal/presentation/errors"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AuthInterceptor - мидлвар для аутентификации с использованием JWT
type AuthInterceptor struct {
    jwtKey []byte // ключ для подписи JWT
}

func NewAuthInterceptor(jwtKey []byte) *AuthInterceptor { return &AuthInterceptor{jwtKey: jwtKey} }

// Unary возвращает новый UnaryServerInterceptor для аутентификации запросов
func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
    return func(
        ctx context.Context,
        req interface{},
        info *grpc.UnaryServerInfo,
        handler grpc.UnaryHandler,
    ) (interface{}, error) {
        // Пропуск аутентификации для регистрации и входа
        if info.FullMethod == "/server.AuctionService/Register" || info.FullMethod == "/server.AuctionService/Login" {
            return handler(ctx, req)
        }

        // Получение метаданных из входящего контекста
        md, ok := metadata.FromIncomingContext(ctx)

        if !ok {
            return nil, &derr.NeedAuntification{
                Err:  "Not provided JWT",
                Code: 401,
            }
        }

        // Декодирование JWT из метаданных
        auth := md.Get("Auntification")

        if len(auth) == 0 {
            return nil, &derr.NeedAuntification{
                Err:  "Not JWT",
                Code: 401,
            }
        }

        _, err := users.DecodeJWT(auth[0], string(config.AppConfig.SecretKey))

        if err != nil {
            return nil, err
        }

        // Продолжение обработки запроса
        return handler(ctx, req)
    }
}