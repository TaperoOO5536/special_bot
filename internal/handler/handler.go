package handler

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/TaperoOO5536/special_bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot     *tgbotapi.BotAPI
	adminID int64
}

func NewHandler(bot *tgbotapi.BotAPI, adminID int64) Handler {
	return Handler{bot: bot, adminID: adminID}
}

func (h *Handler) HandleMessage(eventType string, topic string, msg []byte) {
	if topic == "orders" {
		var order models.Order
		json.Unmarshal(msg, &order)
		userId, err := strconv.Atoi(order.UserID)
		if err != nil {
			fmt.Println("failed to parse user id", err)
		}
		h.handleOrderMessage(eventType, userId, order)
	} else if topic == "userevents" {
		var userEvent models.UserEvent
		json.Unmarshal(msg, &userEvent)
		h.handleUserEventMessage(eventType, userEvent)
	}
}

func (h *Handler) handleOrderMessage(eventType string, userId int, order models.Order) {
	var text string
	switch eventType {
	case "order.create":
		text = "ваш заказ создан, номер вашего заказа: " + order.Number
	case "order.update":
		text = "ваш заказ номер " + order.Number +
						"изменён, новый статус вашего заказа: " + order.Status
	case "order.delete":
		text = "ваш заказ номер " + order.Number + " удалён"
	default:
		fmt.Println("unknown event type")
	}
	h.sendMessage(userId, text)
}

func (h *Handler) handleUserEventMessage(eventType string, userEvent models.UserEvent) {
	var text string
	switch eventType {
	case "userevent.create":
		text = "Новая запись на мероприятие " + userEvent.EventTitle +
	" записался " + userEvent.UserNickName + " мест свободно " + 
	strconv.FormatInt(userEvent.EventOccupiedSeats, 10) + "/" + strconv.FormatInt(userEvent.EventTotalSeats, 10)
	case "userevent.update":
		text = "Изменена запись на мероприятие " + userEvent.EventTitle +
	" пользователя " + userEvent.UserNickName + " мест свободно " + 
	strconv.FormatInt(userEvent.EventOccupiedSeats, 10) + "/" + strconv.FormatInt(userEvent.EventTotalSeats, 10)
	case "userevent.delete":
		text = "Отменена запись на мероприятие " + userEvent.EventTitle +
	" пользователя " + userEvent.UserNickName + " мест свободно " + 
	strconv.FormatInt(userEvent.EventOccupiedSeats, 10) + "/" + strconv.FormatInt(userEvent.EventTotalSeats, 10)
	default:
		fmt.Println("unknown event type")
	}
	h.sendMessage(int(h.adminID), text)
}

func (h *Handler) sendMessage(id int, text string) {
	tgMsg := tgbotapi.NewMessage(int64(id), text)
	h.bot.Send(tgMsg)
}
