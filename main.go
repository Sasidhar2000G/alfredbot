package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//chat photo struct

// Error Codes
const (
	NO_ERROR int = iota
	TOKEN_FILE_MISSING
)

// APIKey for alfredbot from @botFather
var token string = func() string {
	token, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Println(err)
		return ""
	}
	log.Println(string(token))
	return string(token)
}()

// right now the bot just echoes messages sent to it.
func main() {

	if len(token) == 0 {
		log.Println("Token file not present.")
		os.Exit(TOKEN_FILE_MISSING)
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true // debug turned on for now for testing purposes.

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			var msg tgbotapi.MessageConfig
			switch strings.ToLower(update.Message.Text) {
			case "hi":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Alfred at your service!")
			case "what can you do ?":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Right now nothing, but worry not soon I'll be ready.")
			case "haren shylak":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Snake")
			default:
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			}
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
