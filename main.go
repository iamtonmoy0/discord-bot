package main

import (
	"fmt"

	"github.com/iamtonmoy0/discord-bot/bot"
	"github.com/iamtonmoy0/discord-bot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()

	<-make(chan struct{})
	return
}
