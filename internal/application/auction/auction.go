package auction

import (
	"auction/internal/application/config"
	db "auction/internal/application/database"
	"auction/internal/application/logger"
	"auction/internal/domain/entity"
	"auction/internal/domain/service"
	"auction/internal/infrastructure/repository/notificator"
	psql "auction/internal/infrastructure/repository/postgres"
)

var(
    auctionDb = psql.NewAuctionDb(logger.Log)
    bidDb = psql.NewBidDb(logger.Log)
    budService = service.NewBidService(db.DB, bidDb)
    lotDb = psql.NewLotDb(logger.Log)
    not = collectNotificator()
    auctionService = service.NewAuctionService(db.DB, auctionDb, not, lotDb)
)

type AuctionRepo struct{}

func NewAuctionRepo() *AuctionRepo { return &AuctionRepo{} }

func (a *AuctionRepo) CreateLot(lot entity.Lot) (int, error) {
    id, err := auctionService.CreateLot(lot)

    if err != nil {
        return -1, err
    }

    return id, nil
}

func (h *AuctionRepo) StartAuction(auction entity.Auction) (int, error) {
    id, err := auctionService.StartAuction(auction)
    
    if err != nil {
        return -1, err
    }

    return id, nil
}

func (h *AuctionRepo) FinishAuction(auction entity.Auction) (string, error) {
    winner, err := auctionService.FindWinner(auction)

    if err != nil {
        return "", err
    }

    return winner, nil
}

func (h *AuctionRepo) PlaceBid(bid entity.Bid) (int, error) {
    id, err := budService.PlaceBid(bid)

    if err != nil {
        return -1, err
    }

    return id,nil
}

// collectNotificator создает новый экземпляр сервиса уведомлений
func collectNotificator() *notificator.NotificationService {
    smtpHost := config.AppConfig.SmtpHost
    smtpPort := config.AppConfig.SmtpPort
    username := config.AppConfig.MailUser
    password := config.AppConfig.MailPass
    
    return notificator.NewNotificationService(db.DB, smtpHost, smtpPort, username, password, logger.Log)
}