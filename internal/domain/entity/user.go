package entity

import (
	derr "auction/internal/domain/errors"
	"fmt"
	"regexp"
)

type User struct {
    id         int     		// идентификатор пользователя
    email      string  		// email пользователя
    password   string  		// пароль пользователя
    count      float64 		// количество доступных средств
    freezeCount float64 	// количество замороженных средств
}

func CreateUser(id int, count, freeze float64, email, password string) (*User, error) {
    // Проверки валидности
    if !validEmail(email) {
        return nil, &derr.NonValidData{
            Err:  fmt.Sprintf("Invalid email: %s", email),
            Code: 400,
        }
    }

    if len(password) != 64 {
        return nil, &derr.NonValidData{
            Err:  "Invalid password",
            Code: 400,
        }
    }

    return &User{
        id:         id,
        email:      email,
        password:   password,
        count:      count,
        freezeCount: freeze,
    }, nil
}

// Валидация почты
func validEmail(email string) bool {
    var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return emailRegex.MatchString(email)
}

// Геттеры для полей класса
func (u *User) Id() int { return u.id }

func (u *User) Email() string { return u.email }

func (u *User) Password() string { return u.password }

func (u *User) Count() float64 { return u.count }

func (u *User) Freeze() float64 { return u.freezeCount }