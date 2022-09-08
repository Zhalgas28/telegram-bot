package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Help(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"")

	for _, c := range commands {
		msg.Text += c + "\n"
	}
	c.bot.Send(msg)
}
