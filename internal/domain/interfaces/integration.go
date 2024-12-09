// Пакет для описания интерфейсов
package interfaces

import (
	"auction/internal/domain/entity"
	"auction/internal/domain/odt"
	"database/sql"
)

type DatabaseRepo interface {
	Connect() error
	Query(string, ...any) (*sql.Rows, error)
	Close()
}

type PaymentRepo interface {
	Payment(amount float64) error
}

type NotificatiorRepo interface {
	NoteEnd(entity.Auction, odt.FinalAuction) error
	NoteStart(entity.Auction) error
}

type LoggerRepo interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}