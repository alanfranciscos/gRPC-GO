package main

import (
	"database/sql"
	"net"

	database "github.com/alanfranciscos/gRPC-GO/internal/databases"
	"github.com/alanfranciscos/gRPC-GO/internal/pb"
	"github.com/alanfranciscos/gRPC-GO/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Abre a conexão com db
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// cria o banco de dados
	categoryDb := database.NewCategory(db)
	// cria o serviço de categoria passando o banco de dados como parametro para que consiga fazer as operações no banco
	categoryService := service.NewCategoryService(*categoryDb)

	// cria o servidor grpc
	grpcServer := grpc.NewServer()
	// registra o serviço no servidor grpc
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	// registra o servidor grpc para que consiga fazer a comunicação com o evans (cliente grpc)
	reflection.Register(grpcServer)

	// Abre a conexão tcp na porta 50051 para que consiga ter a comunicação com o grpc
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	// inicia o servidor grpc
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
