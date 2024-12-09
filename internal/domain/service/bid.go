package service

import (
	"auction/internal/domain/entity"
	"auction/internal/domain/interfaces"
)

type BidService struct {
    db  interfaces.DatabaseRepo // репозиторий базы данных
    bid interfaces.BidRepo      // репозиторий ставки
}

func NewBidService(db interfaces.DatabaseRepo, bid interfaces.BidRepo) *BidService {
    return &BidService{
        db:  db,
        bid: bid,
    }
}

// PlaceBid размещает новую ставку и возвращает ее идентификатор
func (b *BidService) PlaceBid(bid entity.Bid) (int, error) {
    id, err := b.bid.Place(b.db, bid)

    if err != nil {
        return -1, err
    }

    return id, nil
}
