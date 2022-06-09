package main

import (
	"com-line-bot/server"
	"log"
	"net/http"
	"os"
)

func main() {
	bot := new(server.Bot)
	bot.New()

	http.HandleFunc("/callback", bot.Callback)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
