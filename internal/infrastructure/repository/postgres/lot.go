package postgres

import (
	"auction/internal/domain/entity"
	"auction/internal/domain/interfaces"
	derr "auction/internal/infrastructure/errors"
	"encoding/json"
	"errors"
	"fmt"
)

type LotBd struct{
    logger  interfaces.LoggerRepo
}

func NewLotDb(logger interfaces.LoggerRepo) *LotBd {
    return &LotBd{
        logger: logger,
    }
}

// Get получает лот по его идентификатору
func (l *LotBd) Get(db interfaces.DatabaseRepo, id int) (*entity.Lot, error) {
    l.logger.Debug("getting start")
    rows, err := db.Query(`SELECT lots."name", lots."parameters", users."id", users."email", users."password", users."count", users."freeze_count" FROM lots JOIN users ON lots."owner_id" = users."id" WHERE lots."id" = $1`, id)
    defer rows.Close()

    if err != nil {
        l.logger.Error("database query error:" + err.Error())
        return nil, err
    }
    
    var name, params, email, pass string
    var userId int
    var count, freeze float64
    if rows.Next() {
        rows.Scan(&name, &params, &userId, &email, &pass, &count, &freeze)
    } else {
        return nil, &derr.NoDataInBase{
            Err:  fmt.Sprintf("Lot with id: %d not exist", id),
            Code: 400,
        }
    }

    user, err := entity.CreateUser(userId, count, freeze, email, pass)

    if err != nil {
        l.logger.Error("database query error:" + err.Error())
        return nil, err
    }

    var parameters entity.LotParams
    err = json.Unmarshal([]byte(params), &parameters)

    if err != nil {
        l.logger.Error("Json decode error")
        return nil, &derr.JsonCodeError{
            Err:  "Json decode error",
            Code: 500,
        }
    }

    lot, err := entity.CreateLot(id, name, parameters.Description, user, parameters.MinPrice)

    if err != nil {
        l.logger.Error("database query error:" + err.Error())
        return nil, err
    }

    l.logger.Debug("getting successful")
    return lot, nil
}

// Create создает новый лот в базе данных и возвращает его идентификатор
func (l *LotBd) Create(db interfaces.DatabaseRepo, lot entity.Lot) (int, error) {
    l.logger.Debug("creating start")
    rows, err := db.Query(`INSERT INTO lots ("name", "owner_id", "parameters") VALUES ($1, $2, $3) RETURNING id`, lot.Name(), lot.Owner().Id(), lot.Params())
    
    if err != nil {
        l.logger.Error("database query error:" + err.Error())
        return -1, err
    }

    var id int
    if rows.Next() {
        rows.Scan(&id)
    }

    l.logger.Debug("creating successful")
    return id, nil
}

// Put обновляет существующий лот или создает новый, если лот не существует
func (l *LotBd) Put(db interfaces.DatabaseRepo, lot entity.Lot) error {
    l.logger.Debug("putting start")
    _, err := l.Get(db, lot.Id())

    if errors.Is(err, &derr.NoDataInBase{}) {
        _, err := l.Create(db, lot)
        return err
    }

    _, err = db.Query(`UPDATE lots SET "name" = $2, "owner_id" = $3, "parameters" = $4 WHERE "id" = %1`, lot.Id(), lot.Name(), lot.Owner().Id(), lot.Params())

    if err != nil {
        l.logger.Error("database query error:" + err.Error())
        return err
    }

    l.logger.Debug("putting successful")
    return nil
}

// Delete удаляет лот из базы данных по его идентификатору
func (l *LotBd) Delete(db interfaces.DatabaseRepo, id int) error {
    l.logger.Debug("deleting start")
    
    rows, err := db.Query(`DELETE FROM lots WHERE "id" = $1`, id)
    defer rows.Close()

    if err != nil {
        l.logger.Error("database query error:" + err.Error())
        return err
    }
    
    if !rows.Next() {
        return &derr.NoDataInBase{
            Err:  fmt.Sprintf("Lot with id: %d not exist", id),
            Code: 400,
        }
    }

    l.logger.Debug("deleting successful")
    return nil
}