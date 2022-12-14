package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Get(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	idx, err := c.getIdFromArgs(update)

	if err != nil {
		return
	}

	entity, err := c.service.Get(idx)
	if err != nil {
		msg.Text = fmt.Sprintf("Entity with id %d doesn't exists.", idx)
		c.bot.Send(msg)
		return
	}
	msg.Text = entity.Title
	c.bot.Send(msg)
}
