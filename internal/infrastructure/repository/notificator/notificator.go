package notificator

import (
	"auction/internal/domain/entity"
	"auction/internal/domain/interfaces"
	"auction/internal/domain/odt"
)

type NotificationService struct {
    db       interfaces.DatabaseRepo  // репозиторий базы данных
    smtpHost string
    smtpPort int 
    username string
    password string
    logger   interfaces.LoggerRepo
}

func NewNotificationService(db interfaces.DatabaseRepo, smtpHost string, smtpPort int, username, password string, logger interfaces.LoggerRepo) *NotificationService {
    return &NotificationService{
        db:       db,
        smtpHost: smtpHost,
        smtpPort: smtpPort,
        username: username,
        password: password,
        logger:   logger,
    }
}

// NoteStart отправляет уведомление о начале аукциона
func (n *NotificationService) NoteStart(auction entity.Auction) error {
    // n.logger.Debug("connecting to smtp")
    // msg := gomail.NewMessage()
    // msg.SetHeader("From", n.username)
    // msg.SetHeader("Subject", "AUCTION!")
    // msg.SetBody("text/plain", fmt.Sprintf("Auction has started! Lot: %s Starting price: %.2f", auction.Lot().Name(), auction.Lot().MinPrice()))

    // interf, err := n.db.Query(`SELECT "email" FROM users`)
    // if err != nil {
    //     n.logger.Error("database query error:" + err.Error())
    //     return err
    // }

    // if interf == nil {
    //     return nil
    // }

    // rows := interf.(*sql.Rows)

    // var addresses []string
    // for rows.Next() {
    //     var user string

    //     if err := rows.Scan(&user); err != nil {
    //         n.logger.Error("scan error: " + err.Error())
    //         return err
    //     }

    //     addresses = append(addresses, user)
    // }

    // msg.SetHeader("To", addresses...)

    // d := gomail.NewDialer(n.smtpHost, n.smtpPort, n.username, n.password)

    // n.logger.Debug("sending messages about the start of the auction")
    // if err := d.DialAndSend(msg); err != nil {
    //     n.logger.Error("sending email error: " + err.Error())
    //     return &derr.NotificatorError{
    //         Err:  "Start auction error",
    //         Code: 500,
    //     }
    // }

    // n.logger.Debug("messages were successfully delivered")
    return nil
}

// NoteEnd отправляет уведомление о завершении аукциона
func (n *NotificationService) NoteEnd(auction entity.Auction, res odt.FinalAuction) error {
    // n.logger.Debug("connecting to smtp")
    // msg := gomail.NewMessage()
    // msg.SetHeader("From", n.username)
    // msg.SetHeader("To", res.Winner)
    // msg.SetHeader("Subject", "AUCTION IS OVER!")
    // msg.SetBody("text/plain", fmt.Sprintf("YOU ARE THE WINNER! Lot: %s. CONGRATULATIONS!", auction.Lot().Name()))

    // d := gomail.NewDialer(n.smtpHost, n.smtpPort, n.username, n.password)

    // n.logger.Debug("sending messages for winner")
    // if err := d.DialAndSend(msg); err != nil {
    //     n.logger.Error("sending email error: " + err.Error())
    //     return &derr.NotificatorError{
    //         Err:  "End auction error",
    //         Code: 500,
    //     }
    // }

    // msg.SetBody("text/plain", fmt.Sprintf("AUCTION IS OVER!\nLot: %s.", auction.Lot().Name()))
    // msg.SetHeader("To", res.Other...)

    // n.logger.Debug("sending messages for others")
    // if err := d.DialAndSend(msg); err != nil {
    //     n.logger.Error("sending email error: " + err.Error())
    //     return &derr.NotificatorError{
    //         Err:  "End auction error",
    //         Code: 500,
    //     }
    // }

    // n.logger.Debug("messages were successfully delivered")
    return nil
}