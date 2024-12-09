package server

import (
	"context"
	"strconv"
	"strings"
	"time"

	app "auction/internal/application/auction"
	"auction/internal/application/config"
	"auction/internal/application/database"
	"auction/internal/application/logger"
	"auction/internal/application/users"
	"auction/internal/domain/entity"
	psql "auction/internal/infrastructure/repository/postgres"
	derr "auction/internal/presentation/errors"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (h SessionManager) CreateLot(ctx context.Context, lot *Lot) (*ID, error) {
    md, _ := metadata.FromIncomingContext(ctx)
    token, _ := users.DecodeJWT(md.Get("Auntification")[0], string(config.AppConfig.SecretKey))

    userDB := psql.NewUserDb(logger.Log)
    user, err := userDB.Get(database.DB, token.Id)

    if err != nil {
        return nil, status.Errorf(400, err.Error())
    }

    data, err := entity.CreateLot(-1, lot.Name, lot.Description, user, lot.MinPrice)

    if err != nil {
        return nil, status.Errorf(400, err.Error())
    }

    repo := app.NewAuctionRepo()
    id, err := repo.CreateLot(*data)

    if err != nil {
        return nil, err
    }

    return &ID{
        Id: int32(id),
    }, err
}

// StartAuction запускает аукцион
func (h SessionManager) StartAuction(ctx context.Context, auction *Auction) (*ID, error) {
    md, _ := metadata.FromIncomingContext(ctx)
    token, _ := users.DecodeJWT(md.Get("Auntification")[0], string(config.AppConfig.SecretKey))

    lotDB := psql.NewLotDb(logger.Log)
    lot, err := lotDB.Get(database.DB, int(auction.Lot))

    if err != nil {
        return nil, err
    }

    if lot.Owner().Id() != token.Id {
        return nil, status.Errorf(400, "No rights for this action")
    }

	// Валидация даты
    dateTime := strings.Split(auction.Expires, ".")
    day, err := strconv.Atoi(dateTime[0])
    if err != nil {
        return nil, &derr.InvalidDate{
            Err:  "Invalid format of date",
            Code: 400,
        }
    }
    month_id, err := strconv.Atoi(dateTime[1])
    if err != nil {
        return nil, &derr.InvalidDate{
            Err:  "Invalid format of date",
            Code: 400,
        }
    }
    month := time.Month(month_id)
    year, err := strconv.Atoi(dateTime[2])
    if err != nil {
        return nil, &derr.InvalidDate{
            Err:  "Invalid format of date",
            Code: 400,
        }
    }
    hour, err := strconv.Atoi(dateTime[3])
    if err != nil {
        return nil, &derr.InvalidDate{
            Err:  "Invalid format of date",
            Code: 400,
        }
    }
    minute, err := strconv.Atoi(dateTime[4])
    if err != nil {
        return nil, &derr.InvalidDate{
            Err:  "Invalid format of date",
            Code: 400,
        }
    }

    expires := time.Date(year, month, day, hour, minute, 0, 0, time.Local)
    data, err := entity.CreateAuction(-1, int(auction.MinStep), lot, expires)

    if err != nil {
        return nil, status.Errorf(400, err.Error())
    }

    repo := app.NewAuctionRepo()
    id, err := repo.StartAuction(*data)

    if err != nil {
        return nil, status.Errorf(500, err.Error())
    }

    return &ID{
        Id: int32(id),
    }, err
}

// PlaceBid размещает ставку на аукцион
func (h SessionManager) PlaceBid(ctx context.Context, bid *Bid) (*ID, error) {
    md, _ := metadata.FromIncomingContext(ctx)
    token, _ := users.DecodeJWT(md.Get("Auntification")[0], string(config.AppConfig.SecretKey))

    auctionDB := psql.NewAuctionDb(logger.Log)
    auction, err := auctionDB.Get(database.DB, int(bid.Auction))

    if err != nil {
        return nil, status.Errorf(400, err.Error())
    }

    userDB := psql.NewUserDb(logger.Log)
    user, err := userDB.Get(database.DB, token.Id)

    if err != nil {
        return nil, status.Errorf(400, err.Error())
    }

    bidEntity, err := entity.CreateBid(-1, int(bid.Bid), user, auction)

    if err != nil {
        return nil, status.Errorf(400, err.Error())
    }

    repo := app.NewAuctionRepo()
    id, err := repo.PlaceBid(*bidEntity)

    if err != nil {
        return nil, status.Errorf(400, err.Error())
    }

    return &ID{
        Id: int32(id),
    }, nil
}