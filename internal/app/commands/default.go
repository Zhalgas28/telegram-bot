package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) Default(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"This command doesn't exists!\nClick here /help for more information.")
	c.bot.Send(msg)
}
