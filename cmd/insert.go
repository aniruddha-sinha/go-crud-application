/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aniruddha-sinha/go-crud-app-cobra/internal/config"
	"github.com/aniruddha-sinha/go-crud-app-cobra/internal/database"
	"github.com/aniruddha-sinha/go-crud-app-cobra/internal/store"
	"github.com/spf13/cobra"
)

var name, email string

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert user by name and email",
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

		id, err := userStore.CreateUser(ctx, name, email)
		if err != nil {
			log.Fatalf("Error creating User %v\n", err)
		}

		fmt.Printf("Successfully created user with ID: %d\n", id)
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)

	insertCmd.Flags().StringVarP(&name, "user-name", "", "", "user\\'s name")
	insertCmd.Flags().StringVarP(&email, "email-id", "", "", "user\\'s email ID")

	if err := insertCmd.MarkFlagRequired("user-name"); err != nil {
		log.Fatal(err)
	}

	if err := insertCmd.MarkFlagRequired("email-id"); err != nil {
		log.Fatal(err)
	}

}
