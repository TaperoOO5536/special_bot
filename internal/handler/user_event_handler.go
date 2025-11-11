package handler

// import (
// 	"encoding/json"
// 	"strconv"

// 	"github.com/TaperoOO5536/special_bot/internal/models"
// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// type UserEventHandler struct {
// 	Topic   string
// 	bot     *tgbotapi.BotAPI
// 	AdminID int64
// }

// func NewUserEventHandler(topic string, bot *tgbotapi.BotAPI) UserEventHandler {
// 	return UserEventHandler{Topic: topic, bot: bot}
// }

// func (h *UserEventHandler) CreateUserEvent(msg []byte) {
// 	var userEvent models.UserEvent
// 	json.Unmarshal(msg, &userEvent)
// 	text := "Новая запись на мероприятие " + userEvent.EventTitle +
// 	" записался " + userEvent.UserNickName + " мест свободно " +
// 	strconv.FormatInt(userEvent.EventOccupiedSeats, 10) + "/" + strconv.FormatInt(userEvent.EventTotalSeats, 10)
// 	tgMsg := tgbotapi.NewMessage(h.AdminID, text)
// 	h.bot.Send(tgMsg)
// }

// func (h *UserEventHandler) UpdateUserEvent(msg []byte) {
// 	var userEvent models.UserEvent
// 	json.Unmarshal(msg, &userEvent)
// 	text := "Изменена запись на мероприятие " + userEvent.EventTitle +
// 	" пользователя " + userEvent.UserNickName + " мест свободно " +
// 	strconv.FormatInt(userEvent.EventOccupiedSeats, 10) + "/" + strconv.FormatInt(userEvent.EventTotalSeats, 10)
// 	tgMsg := tgbotapi.NewMessage(h.AdminID, text)
// 	h.bot.Send(tgMsg)
// }

// func (h *UserEventHandler) DeleteUserEvent(msg []byte) {
// 	var userEvent models.UserEvent
// 	json.Unmarshal(msg, &userEvent)
// 	text := "Изменена запись на мероприятие " + userEvent.EventTitle +
// 	" пользователя " + userEvent.UserNickName + " мест свободно " +
// 	strconv.FormatInt(userEvent.EventOccupiedSeats, 10) + "/" + strconv.FormatInt(userEvent.EventTotalSeats, 10)
// 	tgMsg := tgbotapi.NewMessage(h.AdminID, text)
// 	h.bot.Send(tgMsg)
// }