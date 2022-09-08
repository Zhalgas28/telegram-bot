package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Delete(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	idx, err := c.getIdFromArgs(update)

	if err != nil {
		return
	}

	if err = c.service.Delete(idx); err != nil {
		msg.Text = fmt.Sprintf("Entity with id %d doesn't exists.", idx)
		c.bot.Send(msg)
		return
	}
	msg.Text = "Entity deleted!"
	c.bot.Send(msg)
}
