package main

import (
	"os"

	"github.com/Peikkin/discord_bot/bot"
	"github.com/Peikkin/discord_bot/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	err := config.ReadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка чтения конфигурации")
	}

	bot.Start()

	<-make(chan struct{})
}
