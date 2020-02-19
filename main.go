package main

import (
	//"database/sql"
	"github.com/ivansukach/profile-service/repositories"
	"net"

	//"github.com/jmoiron/sqlx"
	"github.com/ivansukach/profile-service/protocol"
	"github.com/ivansukach/profile-service/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	//"net"
)

func main() {
	rps, _ := repositories.OpenPostgreSQLRepository()

	log.Info("GRPC-server started")
	s := grpc.NewServer()                         //создали сервер
	srv := &server.Server{}                       //ссылка на пустую структуру
	protocol.RegisterProfileServiceServer(s, srv) //зарегистрировали сервер
	listener, err := net.Listen("tcp", ":1323")   //просто слушаем порт
	s.Serve(listener)
	if err != nil {
		log.Error(err)
	}

	defer rps.CloseDB()
}
