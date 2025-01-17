package main

import (
	"fmt"
	"os"

	"github.com/JulianH99/bills/cmd"
	"github.com/JulianH99/bills/internal/data"
	"github.com/JulianH99/bills/internal/ui"
)

func main() {
	database := data.NewDatabase("bills.db")
	err := database.Open()
	if err != nil {
		fmt.Printf("Error opening database. Error %v\n", err)
	}
	err = database.Initialize()
	if err != nil {
		fmt.Printf("Error intializing database. Error: %v\n", err)
	}
	defer database.Close()

	rootCmd := cmd.CreateRootCommand(database.Instance())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(ui.Message(err.Error(), true))
		os.Exit(1)
	}
}
