package bot

import (
	"github.com/rs/zerolog/log"

	"github.com/Peikkin/discord_bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка создания сессии")
	}
	user, err := goBot.User("@me")
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка получения информации о пользователе")
	}

	BotID = user.ID
	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка открытия сессии")
	}
	log.Info().Msg("Бот запущен")
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == BotID {
		return
	}

	if message.Content == "ping" {
		_, err := session.ChannelMessageSend(message.ChannelID, "pong")
		if err != nil {
			log.Error().Err(err).Msg("Ошибка отправки сообщения")
		}
	}
}
