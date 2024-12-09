package service

import (
	"auction/internal/domain/entity"
	"auction/internal/domain/interfaces"
)

type AuctionService struct {
    db          interfaces.DatabaseRepo       // репозиторий базы данных
    auction     interfaces.AuctionRepo        // репозиторий аукционов
    notificator interfaces.NotificatiorRepo   // репозиторий уведомлений
    lot         interfaces.LotRepo            // репозиторий лотов
}

func NewAuctionService(db interfaces.DatabaseRepo, auction interfaces.AuctionRepo, note interfaces.NotificatiorRepo, lot interfaces.LotRepo) *AuctionService {
    return &AuctionService{
        db:          db,
        auction:     auction,
        notificator: note,
        lot:         lot,
    }
}

// CreateLot создает новый лот и возвращает его идентификатор
func (s *AuctionService) CreateLot(lot entity.Lot) (int, error) {
    id, err := s.lot.Create(s.db, lot)

    if err != nil {
        return -1, err
    }

    return id, nil
}

// StartAuction запускает новый аукцион и отправляет уведомление
func (s *AuctionService) StartAuction(auction entity.Auction) (int, error) {
    id, err := s.auction.Start(s.db, auction)

    if err != nil {
        return -1, err
    }

    err = s.notificator.NoteStart(auction)

    if err != nil {
        return -1, err
    }

    return id, nil
}

// FindWinner завершает аукцион, определяет победителя и отправляет уведомление
func (s *AuctionService) FindWinner(id entity.Auction) (string, error) {
    result, err := s.auction.Stop(s.db, id)

    if err != nil {
        return "", err
    }

    err = s.notificator.NoteEnd(id, *result)
    
    if err != nil {
        return "", err
    }

    return result.Winner, nil
}
