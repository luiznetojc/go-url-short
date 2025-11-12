package main

import (
	"fmt"
	"strings"
	"github.com/luizn/go-url-short/internal/config"
	"github.com/luizn/go-url-short/internal/repository/db"
)

func teste( s string)(string) {
	return strings.ToUpper(s)
}

func main(){
	
	cfg := config.Load()
	session := db.ConnectCassandra(cfg)
	defer session.Close()
}