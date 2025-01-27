package main

import (
	"fmt"
	"os"

	_ "embed"

	"github.com/JulianH99/bills/cmd"
	"github.com/JulianH99/bills/internal/config"
	"github.com/JulianH99/bills/internal/data"
	"github.com/JulianH99/bills/internal/ui"
)

//go:embed icon/icon.svg
var iconContent string

func main() {
	err := config.InitializeConfig()
	if err != nil {
		fmt.Println(ui.Message(err.Error(), true))
		os.Exit(1)
	}

	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println(ui.Message(err.Error(), true))
		os.Exit(1)
	}

	database := data.NewDatabase("bills.db")

	err = database.Open()
	if err != nil {
		fmt.Printf("Error opening database. Error %v\n", err)
	}

	if !cfg.DatabaseCreated {
		err = database.Initialize()
		cfg.DatabaseCreated = true
		_ = config.Set(*cfg)
		if err != nil {
			fmt.Printf("Error intializing database. Error: %v\n", err)
		}
	}

	defer database.Close()

	if cfg.IconPath == "" {
		iconPath, err := config.RegisterIcon(iconContent)
		if err != nil {
			fmt.Println(ui.Message(err.Error(), true))
		}

		cfg.IconPath = iconPath
		_ = config.Set(*cfg)
	}

	rootCmd := cmd.CreateRootCommand(database.Instance())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(ui.Message(err.Error(), true))
		os.Exit(1)
	}
}
