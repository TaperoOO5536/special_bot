package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/TaperoOO5536/special_bot/internal/config"
	"github.com/TaperoOO5536/special_bot/internal/handler"
	"github.com/TaperoOO5536/special_bot/internal/kafka"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config.LoadEnv()
	token := config.GetToken()
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	orderHandler := handler.NewOrderHandler("orders", bot)
	c, err := kafka.NewConsumer([]string{"localhost:5215"}, []string{orderHandler.Topic, "userevents"}, orderHandler)
	if err != nil {
		log.Fatal(err)
	}
	go c.Start()

	go func () {
		  for update := range updates {
		  if update.Message != nil {
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				var text string
				switch update.Message.Text {
				case "/start":
					text = "Welcome"
				default:
					text = "Hello, " + update.Message.From.UserName + "! Your id: " + strconv.FormatInt(update.Message.From.ID, 10)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				bot.Send(msg)
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	c.Stop()
}