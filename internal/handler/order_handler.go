package handler

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/TaperoOO5536/special_bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
type OrderHandler struct {
	Topic string
	bot   *tgbotapi.BotAPI
}

func NewOrderHandler(topic string, bot *tgbotapi.BotAPI) OrderHandler {
	return OrderHandler{Topic: topic, bot: bot}
}

func (h *OrderHandler) CreateOrder(msg []byte) {
	var order models.CreatedOrder
	json.Unmarshal(msg, &order)
	userId, err := strconv.Atoi(order.UserID)
	if err != nil {
		fmt.Println("failed to parse user id", err)
	}
	orderNumber := strconv.FormatInt(int64(order.Number), 10)
	text := "Номер вашего заказа:" + orderNumber
	tgMsg := tgbotapi.NewMessage(int64(userId), text)
	h.bot.Send(tgMsg)
}