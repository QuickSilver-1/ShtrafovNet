package main

import (
	"auction/internal/application/config"
	"auction/internal/application/database"
	"auction/internal/application/workers"
	s "auction/internal/presentation/server"
)

func main() {
	server := s.NewServer(config.AppConfig.GrpcPort, config.AppConfig.HttpPort)

	go workers.KillAuction()
	err := server.StartServer()

	if err != nil {
		return 
	}

	database.DB.Close()
}