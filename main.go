/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"gocli/cmd"
	"github.com/spf13/viper"
)

func main() {
	readEnv()

	cmd.Execute()
}

func readEnv()  {
	fmt.Println("main read Env")
	viper.SetConfigFile(".env")	
	 // Read the .env file
	 if err := viper.ReadInConfig(); err != nil {
        fmt.Printf("Error reading config file, %s", err)
    }

    // Set default values for environment variables
    viper.SetDefault("PORT", "8080")

    // Read environment variables
    viper.AutomaticEnv()

    // Get the values
    port := viper.GetString("PORT")
    dbHost := viper.GetString("DB_HOST")
    dbUser := viper.GetString("DB_USER")
    dbPass := viper.GetString("DB_PASS")

    // Print the values
    fmt.Printf("PORT: %s\n", port)
    fmt.Printf("DB_HOST: %s\n", dbHost)
    fmt.Printf("DB_USER: %s\n", dbUser)
    fmt.Printf("DB_PASS: %s\n", dbPass)
}