package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"

	"auction/internal/application/config"
	"auction/internal/application/logger"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpc "google.golang.org/grpc"
)

// Server представляет сервер с конфигурацией для gRPC и HTTP
type Server struct {
    grpcPort int         // Порт для gRPC сервера
    httpPort int         // Порт для HTTP сервера
    wg       *sync.WaitGroup // Wait группа дляя для синхронизации горутин
}

func NewServer(grpc, http int) *Server {
    return &Server{
        grpcPort: grpc,
        httpPort: http,
        wg:       &sync.WaitGroup{},
    }
}

// StartServer запускает gRPC и HTTP серверы
func (s *Server) StartServer() error {
    // Настройка gRPC сервера
    logger.Log.Info("setupping server")
    lis, err := net.Listen("tcp", ":"+strconv.Itoa(s.grpcPort))

    if err != nil {
        logger.Log.Fatal(fmt.Sprintf("start server error: %v", err))
    }

    // Создание мидлвара
    middleware := NewAuthInterceptor(config.AppConfig.SecretKey)
    server := grpc.NewServer(
        grpc.UnaryInterceptor(middleware.Unary()),
    )

    // Регистрация сервисов на gRPC сервере
    RegisterAuctionServiceServer(server, SessionManager{})

    // Настройка HTTP сервера с использованием gRPC Gateway
    mux := runtime.NewServeMux(
        runtime.WithIncomingHeaderMatcher(customHeaderMatcher),
    )
    opts := []grpc.DialOption{grpc.WithInsecure()}

    err = RegisterAuctionServiceHandlerFromEndpoint(context.Background(), mux, "localhost:"+strconv.Itoa(s.grpcPort), opts)

    if err != nil {
        logger.Log.Fatal(fmt.Sprintf("start server error: %v", err))
    }

    httpAddr := ":" + strconv.Itoa(s.httpPort)

    // Запуск gRPC и HTTP серверов
    s.wg.Add(2)
    go func() {
        defer s.wg.Done()

        err = server.Serve(lis)

        if err != nil {
            logger.Log.Fatal(fmt.Sprintf("grpc server error: %v", err))    
        }
    }()

    go func() {
        defer s.wg.Done()

        logger.Log.Info("startting grpc server successful")

        err = http.ListenAndServe(httpAddr, mux)

        if err != nil {
            logger.Log.Fatal(fmt.Sprintf("http server error: %v", err))
        }
    }()

    logger.Log.Info("startting http server successful")
    s.wg.Wait()
    return nil
}

// Функция для настройки заголовков
func customHeaderMatcher(key string) (string, bool) {
    return key, true
}