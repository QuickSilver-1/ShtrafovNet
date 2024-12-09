package postgres

import (
	"auction/internal/domain/entity"
	"auction/internal/domain/interfaces"
	derr "auction/internal/infrastructure/errors"
	"errors"
	"fmt"
)

type UserDb struct{
    logger interfaces.LoggerRepo
}

func NewUserDb(logger interfaces.LoggerRepo) *UserDb {
    return &UserDb{
        logger: logger,
    }
}

// Get получает пользователя по его идентификатору
func (u *UserDb) Get(db interfaces.DatabaseRepo, id int) (*entity.User, error) {
    u.logger.Debug("getting start")
    rows, err := db.Query(`SELECT "email", "password", "count", "freeze_count" FROM users WHERE "id" = $1`, id)
    
    if err != nil {
        u.logger.Error("database query error:" + err.Error())
        return nil, err
    }

    if !rows.Next() {
        return nil, &derr.NoDataInBase{
            Err:  fmt.Sprintf("User with id: %d not exist", id),
            Code: 400,
        }
    }

    var email, pass string
    var count, freeze float64
    rows.Scan(&email, &pass, &count, &freeze)

    user, err := entity.CreateUser(id, count, freeze, email, pass)

    if err != nil {
        u.logger.Error("database query error:" + err.Error())
        return nil, err
    }

    u.logger.Debug("getting successful")
    return user, nil
}

// GetEmail получает пользователя по его email
func (u *UserDb) GetEmail(db interfaces.DatabaseRepo, email string) (*entity.User, error) {
    u.logger.Debug("getting start")
    rows, err := db.Query(`SELECT "id", "password", "count", "freeze_count" FROM users WHERE "email" = $1`, email)
    
    if err != nil {
        if err.Error() == "pq: duplicate key value violates unique constraint" {
            return nil, &derr.DublicateError{
                Err: "User with this email already exist",
                Code: 400,
            }
        }
        u.logger.Error("database query error: " + err.Error())
        return nil, err
    }

    if !rows.Next() {
        return nil, &derr.NoDataInBase{
            Err:  fmt.Sprintf("User with email: %s not exist", email),
            Code: 400,
        }
    }

    var pass string
    var id int
    var count, freeze float64
    rows.Scan(&id, &pass, &count, &freeze)

    user, err := entity.CreateUser(id, count, freeze, email, pass)

    if err != nil {
        u.logger.Error("database query error:" + err.Error())
        return nil, err
    }

    u.logger.Debug("getting successful")
    return user, nil
}

// Create создает нового пользователя в базе данных и возвращает его идентификатор
func (u *UserDb) Create(db interfaces.DatabaseRepo, user entity.User) (int, error) {
    u.logger.Debug("creating start")
    _, err := db.Query(`INSERT INTO users (email, password, count, freeze_count) VALUES ($1, $2, $3, $4) RETURNING id, email`, user.Email(), user.Password(), user.Count(), user.Freeze())

    if err != nil {
        if err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"` {
            return -1, &derr.DublicateError{
                Err:  "A user with this email already exists",
                Code: 400,
            }
        }

        u.logger.Error("database query error:" + err.Error())
        return -1, err
    }

    userNew, err := u.GetEmail(db, user.Email())
    
    if err != nil {
        u.logger.Error("database query error:" + err.Error())
        return -1, err
    }

    u.logger.Debug("creating successful")
    return userNew.Id(), nil
}

// Put обновляет существующего пользователя или создает нового, если пользователь не существует
func (u *UserDb) Put(db interfaces.DatabaseRepo, user entity.User) error {
    u.logger.Debug("putting start")
    _, err := u.Get(db, user.Id())

    if errors.Is(err, &derr.NoDataInBase{}) {
        _, err := u.Create(db, user)
        return err
    }

    _, err = db.Query(`UPDATE users SET "email" = $2, "password" = $3, "count" = $4, "freeze_count" = $5 WHERE "id" = $1`, user.Id(), user.Email(), user.Password(), user.Count(), user.Freeze())

    if err != nil {
        u.logger.Error("database query error:" + err.Error())
        return err
    }

    u.logger.Debug("putting successful")
    return nil
}