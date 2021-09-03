package main

import (
	// "net/http"

	app "github.com/rizgust/go-finance/src/config"
	db "github.com/rizgust/go-finance/src/providers/database"
)

func main() {
	// get configuration stucts via .env file
	configuration, err := app.NewConfig()
	if err != nil {
		panic(err)
	}
	// establish DB connection
	db, err := db.Connect(configuration.Database)
	if err != nil {
		panic(err)
	} 

	// err = http.ListenAndServe(":"+configuration.Port, httpRouter)
	// if err != nil {
	// 	panic(err)
	// }

	defer db.Close()
}