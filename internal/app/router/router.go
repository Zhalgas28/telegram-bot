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
	if update.CallbackQuery != nil {

		r.bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data))
	} else {
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "help":
				r.commander.Help(update)
			case "list":
				r.commander.List(update)
			case "get":
				r.commander.Get(update)
			case "create":
				r.commander.Create(update)
			case "delete":
				r.commander.Delete(update)
			default:
				r.commander.Default(update)
			}
		}
	}
}
