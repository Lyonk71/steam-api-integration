// Package main sets up and runs the service
package main

import (
	"SteamPurchaseService/pkg/api"
	"SteamPurchaseService/pkg/app"
	"SteamPurchaseService/util"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	itemDef, err := util.LoadItemDef("./itemdef.json")
	if err != nil {
		log.Fatal("cannot load items:", err)
	}

	// create router dependency
	router := gin.Default()
	router.Use(cors.Default())

	// create steam service
	steamService := api.NewSteamService(&config, itemDef)

	server := app.NewServer(router, steamService)

	// start the server
	err = server.Run(&config)
	if err != nil {
		return err
	}

	return nil
}
