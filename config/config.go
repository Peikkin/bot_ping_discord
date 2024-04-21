package config

import (
	"encoding/json"
	"os"

	"github.com/rs/zerolog/log"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botPrefix"`
}

func ReadConfig() error {
	log.Info().Msg("Чтение файла конфигурации")

	file, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка чтения файла конфигурации")
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка декодирования конфигурации")
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
