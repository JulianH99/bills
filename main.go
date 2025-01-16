package main

import (
	"fmt"

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
	cmd.ExecuteRootCommand(database.Instance())
}
