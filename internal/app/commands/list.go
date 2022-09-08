package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) List(update *tgbotapi.Update) {
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next Page", "some data"),
		),
	)

	text := ""
	list := c.service.List()

	for i := 0; i < len(list); i += 2 {

	}

	for _, elem := range list {
		text += elem.Title
		text += "\n"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = numericKeyboard

	c.bot.Send(msg)
}
