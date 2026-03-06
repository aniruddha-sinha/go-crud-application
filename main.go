/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/aniruddha-sinha/go-crud-app-cobra/cmd"
	"github.com/aniruddha-sinha/go-crud-app-cobra/internal/config"
	"github.com/aniruddha-sinha/go-crud-app-cobra/internal/database"
	"github.com/aniruddha-sinha/go-crud-app-cobra/internal/store"
)

func main() {
	//load zshrc vars
	cfg := config.LoadConfig()

	//pass the db_url string to init db
	db, cleanup := database.InitDB(cfg.DBURL)
	defer cleanup()

	//init store
	userStore := &store.UserStore{DB: db}

	fmt.Println("db connection successful")
	defer userStore.DB.Close()

	cmd.Execute()
}
