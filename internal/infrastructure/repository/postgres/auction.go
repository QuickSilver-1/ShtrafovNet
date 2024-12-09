package postgres

import (
	"auction/internal/domain/entity"
	"auction/internal/domain/interfaces"
	"auction/internal/domain/odt"
	derr "auction/internal/infrastructure/errors"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type AuctionDb struct{
    logger  interfaces.LoggerRepo
}

func NewAuctionDb(logger interfaces.LoggerRepo) *AuctionDb { 
    return &AuctionDb{
        logger: logger,
    } 
}

// Get получает аукцион по его идентификатору
func (a *AuctionDb) Get(db interfaces.DatabaseRepo, id int) (*entity.Auction, error) {
    a.logger.Debug("getting auction")
    rows, err := db.Query(` SELECT auctions."min_step", auctions."expires", users."id", users."email", users."password", users."count", users."freeze_count", lots."id", lots."name", lots."parameters" FROM auctions JOIN lots ON auctions."lot_id" = lots."id" JOIN users ON lots."owner_id" = users."id" WHERE auctions."id" = $1 `, id)
    defer rows.Close()

    if err != nil {
        a.logger.Error("database query error:" + err.Error())
        return nil, err
    }

    var expires time.Time
    var email, pass, name, paramsJson string
    var minStep, userId, lotId int
    var count, freeze float64
    if rows.Next() {
        rows.Scan(&minStep, &expires, &userId, &email, &pass, &count, &freeze, &lotId, &name, &paramsJson)
    } else {
        return nil, &derr.NoDataInBase{
            Err:  fmt.Sprintf("Auction with id: %d not exist", id),
            Code: 400,
        }
    }

    var parameters entity.LotParams
    err = json.Unmarshal([]byte(paramsJson), &parameters)

    if err != nil {
        a.logger.Error("Json decode error:" + err.Error())
        return nil, &derr.JsonCodeError{
            Err:  "Json decode error",
            Code: 500,
        }
    }

    a.logger.Debug("creating user")
    user, err := entity.CreateUser(userId, count, freeze, email, pass)

    if err != nil {
        a.logger.Error("database query error:" + err.Error())
        return nil, err
    }

    a.logger.Debug("creating lot")
    lot, err := entity.CreateLot(lotId, name, parameters.Description, user, parameters.MinPrice)
    
    if err != nil {
        a.logger.Error("database query error:" + err.Error())
        return nil, err
    }

    a.logger.Debug("creating auction")
    auction, err := entity.CreateAuction(id, minStep, lot, expires)

    if err != nil {
        a.logger.Error("database query error:" + err.Error())
        return nil, err
    }

    a.logger.Debug("getting successful")
    return auction, nil
}


// Start начинает новый аукцион и возвращает его идентификатор
func (a *AuctionDb) Start(db interfaces.DatabaseRepo, auction entity.Auction) (int, error) {
    a.logger.Debug("starting auction")
    rows, err := db.Query(` SELECT id FROM auctions WHERE "lot_id" = $1 `, auction.Lot().Id())

    if err != nil {
        return -1, err
    }

    var lotId int
    if rows.Next() {
        rows.Scan(&lotId)
    }

    if lotId != 0 {
        return -1, &derr.DublicateError{
            Err: "This auction already exist",
            Code: 400,
        }
    }

    rows, err = db.Query(` INSERT INTO auctions ("lot_id", "min_step", "expires") VALUES ($1, $2, $3) RETURNING id `, auction.Lot().Id(), auction.MinStep(), auction.Expires())

    if err != nil {
        return -1, err
    }

    var id int
    if rows.Next() {
        rows.Scan(&id)
    }

    a.logger.Debug("starting successful")
    return id, nil
}

// Stop завершает аукцион и возвращает результаты аукциона
func (a *AuctionDb) Stop(db interfaces.DatabaseRepo, auction entity.Auction) (*odt.FinalAuction, error) {
    a.logger.Debug("stopping auction")
    _, err := db.Query(`DELETE FROM auctions WHERE "id" = $1`, auction.Id())

    if err != nil {
        a.logger.Error("database query error:" + err.Error())
        return nil, err
    }

    _, err = db.Query(`DELETE FROM lots WHERE "id" = $1`, auction.Lot().Id())

    if err != nil {
        a.logger.Error("database query error:" + err.Error())
        return nil, err
    }

    rows, err := db.Query(`SELECT "email", "count", "freeze_count", "bid" FROM bids JOIN users ON bids."user_id" = users."id" WHERE "auction_id" = $1 ORDER BY auctions."id" DESC`, auction.Id())
    defer rows.Close()

    if err != nil {
        a.logger.Error("database query error:" + err.Error())
        return nil, err
    }

    var winner string
    var count, freeze, bit int
    if rows.Next() {
        rows.Scan(&winner, &count, &freeze, &bit)
    } else {
        return nil, &derr.NotSingleBid{
            Err:  "no bits were made",
            Code: 400,
        }
    }
    
    other := []string{}

    var email string
    for rows.Next() {
        rows.Scan(&email, &count, &freeze, &bit)
        other = append(other, email)
    }

    winnerInt, _ := strconv.Atoi(winner)

    _, err = db.Query(` UPDATE users SET "count" = $2, "freeze_count" = $3 WHERE "id" = $1 `, winnerInt, count-bit, freeze-bit)

    if err != nil {
        a.logger.Error("database query error:" + err.Error())
        return nil, err
    }

    _, err = db.Query(` UPDATE users JOIN bids ON user."id" = bids."user_id" SET "freeze_count" = $2 WHERE bids."auction_id" = $1 `, winnerInt, freeze-bit)

    if err != nil {
        a.logger.Error("database query error:" + err.Error())
        return nil, err
    }

    a.logger.Debug("stopping successful")
    return &odt.FinalAuction{
        Winner: winner,
        Other:  other,
    }, nil
}