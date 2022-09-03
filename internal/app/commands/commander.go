package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service"
)

type Commander struct {
	bot     *tgbotapi.BotAPI
	service *service.Service
}

func NewCommander(bot *tgbotapi.BotAPI) *Commander {
	return &Commander{
		bot:     bot,
		service: service.NewService(),
	}
}
