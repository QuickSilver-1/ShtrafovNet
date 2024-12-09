package postgres

import (
	"auction/internal/domain/entity"
	"auction/internal/domain/interfaces"
	derr "auction/internal/infrastructure/errors"
)

type BidDb struct{
    logger interfaces.LoggerRepo
}

func NewBidDb(logger interfaces.LoggerRepo) *BidDb {
    return &BidDb{
        logger: logger,
    }
}

// Place размещает новую ставку в базе данных и возвращает её идентификатор
func (b *BidDb) Place(db interfaces.DatabaseRepo, bid entity.Bid) (int, error) {
    // Получаем последнюю ставку пользователя
    b.logger.Debug("placcing bid")
    rows, err := db.Query(`SELECT "bid" FROM bids WHERE "user_id" = $1 ORDER BY "id" DESC LIMIT 1`, bid.User().Id())
    defer rows.Close()

    if err != nil {
        b.logger.Error("database query error:" + err.Error())
        return -1, err
    }

    var lastBid float64
    if rows.Next() {
        rows.Scan(&lastBid)
    } else {
        lastBid = 0
    }

    // Проверяем, достаточно ли средств у пользователя
    if bid.User().Count() - (bid.User().Freeze() - lastBid) < 0 {
        return -1, &derr.NoMoney{
            Err:  "Not enough funds",
            Code: 400,
        }
    }

    // Получаем последнюю ставку на аукционе
    rows, err = db.Query(`SELECT "bid" FROM bids WHERE "auction_id" = $1 ORDER BY "id" DESC LIMIT 1`, bid.Auction().Id())

    if err != nil {
        b.logger.Error("database query error:" + err.Error())
        return -1, err
    }

    var maxBid float64
    if rows.Next() {
        rows.Scan(&maxBid) 
    } else {
        maxBid = bid.Auction().Lot().MinPrice()
    }

    // Проверяем, превышает ли новая ставка предыдущую на минимальный шаг
    if maxBid + float64(bid.Auction().MinStep()) > float64(bid.Bid()) {
        return -1, &derr.NeedMoreBid{
            Err:  "This bid cannot beat the previous one",
            Code: 400,
        }
    }

    rows, err = db.Query(` INSERT INTO bids (user_id, auction_id, bid) VALUES ($1, $2, $3) RETURNING id `, bid.User().Id(), bid.Auction().Id(), bid.Bid())

    if err != nil {
        return -1, err
    }

    var id int
    if rows.Next() {
        rows.Scan(&id)
    }

    b.logger.Debug("placcing successful")
    return id, nil
}