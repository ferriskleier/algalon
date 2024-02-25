package bot

import "GoGPT/models"

var Bots = map[string]models.Bot{
	"General": {
		Name:      "YOUR BOT NAME, e.g. 'My Chatbot'",
		ShortName: "YOUR BOT SHORTNAME, e.g. 'my_bot'",
		ID:        "RANDOM ID TO SEPERATE MULTIPLE BOTS",
		Role:      "SPECIFY AS INSTRUCTION FOR GPT, e.g. 'You are a helpful assistant with profound knowledge on capybaras'",
		Key:       "YOUR TELEGRAM BOT API KEY",
	},
}
