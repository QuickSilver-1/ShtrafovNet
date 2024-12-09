package workers

import (
	app "auction/internal/application/auction"
	"auction/internal/application/database"
	"auction/internal/application/logger"
	psql "auction/internal/infrastructure/repository/postgres"
	"fmt"
	"time"
)

// KillAuction - запускает тикер, который заканчивает все акционы с истекшим expires раз в минуту
func KillAuction() {
	logger.Log.Info("starting killer")
    ticker := time.NewTicker(time.Minute)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
			logger.Log.Debug("killing start")
			rows, err := database.DB.Query(` SELECT "id" FROM auctions WHERE "expires" < $1 `, time.Now())

            if err != nil {
				logger.Log.Fatal(fmt.Sprintf("KillAuction not work: %v", err))
			}

			if !rows.Next() {
				logger.Log.Debug("killing successful")
				return
			}

			for rows.Next() {
				var id int
				rows.Scan(&id)

				repoDb := psql.NewAuctionDb(logger.Log)
				auction, _ := repoDb.Get(database.DB, id)

				app.NewAuctionRepo().FinishAuction(*auction)
			}

			logger.Log.Debug("killing successful")
        }
    }
}