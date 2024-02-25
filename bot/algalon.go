package bot

import (
	"GoGPT/errorHandler"
	"GoGPT/gpt"
	"GoGPT/logger"
	"GoGPT/models"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Algalon(config models.Bot) {

	fmt.Println("Started " + config.Name)

	bot, err := tgbotapi.NewBotAPI(config.Key)
	errorHandler.Handle(err)
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if _, authorized := AuthorizedUsers[update.Message.From.ID]; !authorized {
			logger.Log(
				config.ShortName,
				"Unauthorized access attempt by "+update.Message.From.UserName,
				update.Message.Text)
			continue
		}

		if len(update.Message.Text) > 0 {
			response := ""
			if strings.ToLower(update.Message.Text) == "/reset" {
				response = "New Chat Started for " + config.Name
				gpt.ClearHistory(config.ID)
			} else {
				response = gpt.GetResponse(config, update.Message.Text)
				logger.Log(config.ShortName, update.Message.From.UserName, update.Message.Text)
				logger.Log(config.ShortName, config.Name, response)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			msg.ParseMode = tgbotapi.ModeMarkdown

			_, err := bot.Send(msg)
			errorHandler.Handle(err)
		}
	}
}
