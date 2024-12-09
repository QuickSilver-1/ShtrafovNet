package users

import (
	"fmt"
	"time"

	derr "auction/internal/application/errors"
	"auction/internal/application/logger"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
    Id    int    `json:"id"`    // идентификатор пользователя
    Email string `json:"email"` // email пользователя
    jwt.StandardClaims
}

func NewToken(id int, email string) *Token {
    return &Token{
        Id:    id,
        Email: email,
    }
}

// MakeToken создает JWT токен для данного email и id
func (t *Token) MakeToken(key string) (string, error) {
    // Устанавливаем время истечения токена
    logger.Log.Debug("creating token")
    expires := time.Now().Add(24 * time.Hour)
    claims := &Token{
        Email: t.Email,
        Id:    t.Id,

        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expires.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenS, err := token.SignedString([]byte(key))

    if err != nil {
        logger.Log.Error(fmt.Sprintf("generation token error: %v", err))
        return "", &derr.JWTError{
            Err:  "Generation token error",
            Code: 500,
        }
    }

    logger.Log.Debug("creating successful")
    return tokenS, nil
}

// DecodeJWT декодирует JWT токен и возвращает Token
func DecodeJWT(tokenStr string, key string) (*Token, error) {
    logger.Log.Debug("decoding start")
    claims := &Token{}
    // Парсинг токена с claims
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(key), nil
    })

    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            return nil, &derr.JWTError{
                Err:  "Invalid form token",
                Code: 401,
            }
        }

        logger.Log.Error(fmt.Sprintf("generation token error: %v", err))
        return nil, &derr.JWTError{
            Err:  fmt.Sprintf("Decode token error: %v", err),
            Code: 500,
        }
    }

    if !token.Valid {
        return nil, &derr.JWTError{
            Err:  "Invalid token",
            Code: 401,
        }
    }

    logger.Log.Debug("decoding success")
    return claims, nil
}