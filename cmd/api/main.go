package main

import (
	"fmt"
	delivery "github.com/luizn/go-url-short/internal/delivery/http"

	"github.com/luizn/go-url-short/internal/config"
	"github.com/luizn/go-url-short/internal/repository"
	"github.com/luizn/go-url-short/internal/repository/db"
	"github.com/luizn/go-url-short/internal/usercase"
)


func main(){
	
	cfg := config.Load()
	session := db.ConnectCassandra(cfg)
	defer session.Close()

	repo := repository.NewUrlRepositoryCassandra(session)
	createUseCase := usercase.NewCreateUrlUseCase(repo)

	handler := delivery.NewUrlHandler(createUseCase)
	router := delivery.NewRouter(handler)

	port := ":8080"
	fmt.Println("🚀 Server running on port", port)
	router.Run(port)
}