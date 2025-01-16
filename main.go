package main

import (
	"fmt"
	"os"

	"github.com/JulianH99/bills/cmd"
	"github.com/JulianH99/bills/internal/data"
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
		fmt.Printf("An error ocurred. Error: %v", err)
		os.Exit(1)
	}
}
