package entity

import (
	derr "auction/internal/domain/errors"
)

type Bid struct {
    id       int       // идентификатор ставки
    user     *User     // пользователь, сделавший ставку
    auction  *Auction  // аукцион, на котором сделана ставка
    bid      int       // Цена ставки
}

func CreateBid(id, bid int, user *User, auction *Auction) (*Bid, error) {
    // Проверки валидности
    if user.Id() == 0 {
        return nil, &derr.NonValidData{
            Err:  "Invalid user's id",
            Code: 400,
        }
    }

    if auction.Id() == 0 {
        return nil, &derr.NonValidData{
            Err:  "Invalid auction's id",
            Code: 400,
        }
    }

    if bid == 0 {
        return nil, &derr.NonValidData{
            Err:  "Invalid bid's value",
            Code: 400,
        }
    }

    return &Bid{
        id:      id,
        user:    user,
        auction: auction,
        bid:     bid,
    }, nil
}

// Геттеры для полей класса
func (b *Bid) User() *User { return b.user }

func (b *Bid) Auction() *Auction { return b.auction }

func (b *Bid) Bid() int { return b.bid }