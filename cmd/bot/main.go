package main

import (
	"fmt"
	"github.com/Sergei3232/bot_v2/internal/postgres"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	os.LookupEnv("TELEGRAM_ID")

	dbBot, err := postgres.NewDBConnect()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(dbBot)

	token, found := os.LookupEnv("TELEGRAM_ID")
	if !found {
		log.Panic("environment variable TELEGRAM_ID not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
