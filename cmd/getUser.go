/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aniruddha-sinha/go-crud-app-cobra/internal/config"
	"github.com/aniruddha-sinha/go-crud-app-cobra/internal/database"
	"github.com/aniruddha-sinha/go-crud-app-cobra/internal/store"
	"github.com/spf13/cobra"
)

var id int

// getUserCmd represents the getUser command
var getUserCmd = &cobra.Command{
	Use:   "getUser",
	Short: "Get User by ID",
	Run: func(cmd *cobra.Command, args []string) {
		//load zshrc vars
		cfg := config.LoadConfig()

		//pass the db_url string to init db
		db, cleanup := database.InitDB(cfg.DBURL)
		defer cleanup()

		//init store
		userStore := &store.UserStore{DB: db}
		fmt.Println("db connection successful")

		//create a background context
		ctx := context.Background()

		//add context timeout
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel() //call cancel to release resources

		user, err := userStore.GetUser(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
		jsonData, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Json data to Json string", string(jsonData))
	},
}

func init() {
	rootCmd.AddCommand(getUserCmd)
	getUserCmd.Flags().IntVarP(&id, "user-id", "", 0, "user id of the user")

	//* mark flag required
	if err := getUserCmd.MarkFlagRequired("user-id"); err != nil {
		log.Fatal(err)
	}
}
