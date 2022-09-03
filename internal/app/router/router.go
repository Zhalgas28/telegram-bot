package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands"
)

type Router struct {
	bot       *tgbotapi.BotAPI
	commander *commands.Commander
}

func NewRouter(bot *tgbotapi.BotAPI) *Router {
	return &Router{
		bot:       bot,
		commander: commands.NewCommander(bot),
	}
}

func (r *Router) HandleUpdate(update *tgbotapi.Update) {
	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "help":
			r.commander.List(update)
		case "list":
			r.commander.List(update)
		case "get":
			r.commander.List(update)
		}
	}
}
