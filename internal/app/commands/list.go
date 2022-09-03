package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) List(update *tgbotapi.Update) {
	text := ""
	list := c.service.List()

	for _, elem := range list {
		text += elem.Title
		text += "\n"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	c.bot.Send(msg)
}
