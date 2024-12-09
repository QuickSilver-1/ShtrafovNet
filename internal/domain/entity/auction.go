package entity

import (
	derr "auction/internal/domain/errors"
	"time"
)

type Auction struct {
    id       int       // идентификатор аукциона
    lot      *Lot      // лот на аукционе
    min_step int       // минимальный шаг, на который можно повысить ставку
    expires  time.Time // дата и время истечения аукциона
}

func CreateAuction(id, min_step int, lot *Lot, time time.Time) (*Auction, error) {
	// Проверки валидности
	if lot.Id() == 0 {
        return nil, &derr.NonValidData{
            Err:  "Invalid lot's id",
            Code: 400,
        }
    }

    if min_step == 0 {
        return nil, &derr.NonValidData{
            Err:  "Invalid minimal step",
            Code: 400,
        }
    }

    t := Auction{}
    if time == t.expires {
        return nil, &derr.NonValidData{
            Err:  "Invalid expires value",
            Code: 400,
        }
    }

    return &Auction{
        id:       id,
        lot:      lot,
        min_step: min_step,
        expires:  time,
    }, nil
}

// Геттеры для полей класса
func (a *Auction) Id() int { return a.id }

func (a *Auction) Lot() *Lot { return a.lot }

func (a *Auction) MinStep() int { return a.min_step }

func (a *Auction) Expires() time.Time { return a.expires }
