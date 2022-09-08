package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service"
	"strconv"
	"strings"
)

var commands = []string{
	"/help - help",
	"/list - list of entities",
	"/get - entity by id",
	"/create - create an entity",
	"/delete - delete an entity",
}

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

func (c *Commander) getIdFromArgs(update *tgbotapi.Update) (int, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"")

	text := update.Message.Text

	args := strings.Split(text, " ")

	if len(args) == 1 {
		msg.Text = fmt.Sprintf("Please, enter entity id.")
		c.bot.Send(msg)
		return 0, fmt.Errorf("error")
	}

	idx, err := strconv.Atoi(args[1])

	if err != nil {
		msg.Text = "Enter integer!"
		c.bot.Send(msg)
		return 0, fmt.Errorf("error")
	}
	return idx, nil
}

func (c *Commander) getStringFromArgs(update *tgbotapi.Update) (string, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"")

	text := update.Message.Text

	args := strings.Split(text, " ")

	if len(args) == 1 {
		msg.Text = fmt.Sprintf("Please, enter entity title.")
		c.bot.Send(msg)
		return "", fmt.Errorf("error")
	}

	return args[1], nil
}
