// Пакет для описания интерфейсов
package interfaces

import (
	"auction/internal/domain/entity"
	"auction/internal/domain/odt"
)

type ServerRepo interface {
	Start() error
	Stop() error
}

type UserRepo interface {
	Registration(*odt.UserOdt) (string, error) //Возвращают JWT для аунтификации
	Login(*odt.UserOdt)	(string, error)
}

type UserDbRepo interface {
	Get(DatabaseRepo, string) (*entity.User, error)
	Create(DatabaseRepo, entity.User) (int, error) //Возвращает id созданного пользователя
	Put(DatabaseRepo, entity.User) error
}

type BidRepo interface {
	Place(DatabaseRepo, entity.Bid) (int, error) //Возвращает id ставки
}

type AuctionRepo interface {
	Start(DatabaseRepo, entity.Auction) (int, error) //Возвращает id запущенного аукциона
	Stop(DatabaseRepo, entity.Auction) (*odt.FinalAuction, error)
}

type LotRepo interface {
	Get(DatabaseRepo, int) (*entity.Lot, error)
	Create(DatabaseRepo, entity.Lot) (int, error) //Возвращает id созданного лота
	Put(DatabaseRepo, entity.Lot) error
	Delete(DatabaseRepo, int) error
}