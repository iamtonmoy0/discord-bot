package main

import (
	"fmt"

	"github.com/iamtonmoy0/discord-bot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.start()

	<-make(chan struct{})
	return
}
