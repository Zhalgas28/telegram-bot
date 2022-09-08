package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service"
)

func (c *Commander) Create(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Entity created!")

	title, err := c.getStringFromArgs(update)

	if err != nil {
		return
	}
	
	entity := service.Entity{Title: title}

	c.service.Create(entity)
	c.bot.Send(msg)
}
