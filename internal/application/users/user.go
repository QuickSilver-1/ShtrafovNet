package users

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"unicode"

	"auction/internal/application/config"
	db "auction/internal/application/database"
	derr "auction/internal/application/errors"
	"auction/internal/application/logger"
	"auction/internal/domain/entity"
	"auction/internal/domain/odt"
	pay "auction/internal/infrastructure/repository/payment"
	repo "auction/internal/infrastructure/repository/postgres"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo { return &UserRepo{} }

// Registration регистрирует нового пользователя и возвращает JWT
func (u *UserRepo) Registration(user odt.UserOdt) (string, error) {
    // Проверка пароля
    logger.Log.Debug("registration start")
    err := validPass(user.Pass)

    if err != nil {
        return "", err
    }

    // Хэширование пароля
    hash := genHash(user.Pass)

    userRepo := repo.NewUserDb(logger.Log)
    userModel, err := entity.CreateUser(0, 0, 0, user.Email, hash)

    if err != nil {
        return "", err
    }

    id, err := userRepo.Create(db.DB, *userModel)

    if err != nil {
        return "", err
    }

    // Создание JWT
    token := NewToken(id, user.Email)
    JWT, err := token.MakeToken(string(config.AppConfig.SecretKey))

    if err != nil {
        return "", err
    }

    logger.Log.Debug("registration successful")
    return JWT, nil
}

// Login выполняет вход пользователя и возвращает JWT
func (u *UserRepo) Login(user odt.UserOdt) (string, error) {
    // Хэширование пароля
    logger.Log.Debug("logging start")
    hash := genHash(user.Pass)

    userRepo := repo.NewUserDb(logger.Log)
    userModel, err := userRepo.GetEmail(db.DB, user.Email)

    if err != nil {
        return "", err
    }

    if hash != userModel.Password() {
        return "", &derr.InvalidPassword{
            Err:  "Incorrect password",
            Code: 401,
        }
    }

    // Создание JWT
    token := NewToken(userModel.Id(), user.Email)
    JWT, err := token.MakeToken(string(config.AppConfig.SecretKey))

    if err != nil {
        return "", err
    }

    logger.Log.Debug("logging successful")
    return JWT, nil
}

// Pay выполняет платеж и обновляет баланс пользователя
func (u *UserRepo) Pay(id int, money float64) error {
    logger.Log.Debug("paying start")
    userRepo := repo.NewUserDb(logger.Log)
    user, err := userRepo.Get(db.DB, id)

    if err != nil {
        return err
    }

    payService := pay.MockPaymentService{}
    err = payService.Payment(user.Id(), money)

    if err != nil {
        return err
    }

    updated, err := entity.CreateUser(user.Id(), user.Count()+money, user.Freeze(), user.Email(), user.Password())
    
    if err != nil {
        return err
    }

    err = userRepo.Put(db.DB, *updated)

    if err != nil {
        return err
    }

    logger.Log.Debug("paying successful")
    return nil
}

// validPass проверяет валидность пароля
func validPass(pass string) error {
    if len(pass) < 3 || len(pass) > 30 {
        return &derr.InvalidPassword{
            Err:  "Password must have from 3 to 30 symbols",
            Code: 400,
        }
    }

    hasUpper, hasDigit := false, false

    // Проверяем каждый символ пароля
    for _, char := range pass {
        switch {
        case unicode.IsUpper(char):
            hasUpper = true // Проверяем наличие заглавной буквы
        case unicode.IsDigit(char):
            hasDigit = true // Проверяем наличие цифры
        case !unicode.IsLetter(char) && !unicode.IsDigit(char) && !strings.Contains("_!@#&*-", string(char)):
            return &derr.InvalidPassword{
                Err:  "The password must consist of letters of the Latin alphabet, numbers and symbols _!@#&*-",
                Code: 400,
            }
        }
    }

    if !hasDigit || !hasUpper {
        return &derr.InvalidPassword{
            Err:  "The password must have at least 1 capital letter and 1 number",
            Code: 400,
        }
    }

    return nil
}

// genHash создает хэш пароля
func genHash(str string) string {
    hasher := sha256.New()
    hasher.Write([]byte(str)) // Преобразуем строку в хэш
    hash := hasher.Sum(nil)

    return hex.EncodeToString(hash) // Возвращаем хэш в виде шестнадцатеричной строки
}