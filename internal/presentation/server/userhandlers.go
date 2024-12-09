package server

import (
	"auction/internal/application/config"
	"auction/internal/application/users"
	"auction/internal/domain/odt"
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Имплементация сервера grpc
type SessionManager struct {
    UnimplementedAuctionServiceServer
}

func (h SessionManager) Register(ctx context.Context, user *UserData) (*JWT, error) {
    data := odt.UserOdt{
        Email: user.Email,
        Pass:  user.Password,
    }

    repo := users.NewUserRepo()
    token, err := repo.Registration(data)

    if err != nil {
        return nil, status.Errorf(400, err.Error())
    }

    return &JWT{
        Token: token,
    }, nil
}

func (h SessionManager) Login(ctx context.Context, user *UserData) (*JWT, error) {
    data := odt.UserOdt{
        Email: user.Email,
        Pass:  user.Password,
    }

    repo := users.NewUserRepo()
    token, err := repo.Login(data)

    if err != nil {
        return nil, status.Errorf(400, err.Error())
    }

    return &JWT{
        Token: token,
    }, nil
}

func (h SessionManager) Pay(ctx context.Context, money *Money) (*Empty, error) {
    md, _ := metadata.FromIncomingContext(ctx)
    token, _ := users.DecodeJWT(md.Get("Auntification")[0], string(config.AppConfig.SecretKey))

    app := users.NewUserRepo()
    err := app.Pay(token.Id, money.Money)

    if err != nil {
        return nil, err
    }

    return nil, nil
}