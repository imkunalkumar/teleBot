package telebot

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

//send reply to users
func ReplyMessege(chatId int64, msg string) error {
	sendMessageObj := SendMessageReqBody{
		ChatID: chatId,
		Text:   msg,
	}
	reqBytes, err := json.Marshal(sendMessageObj)
	if err != nil {
		return err
	}
	url, _ := GetUrl()
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

//this function returns messege text,username,and chatId
func ExtractDataFromUser(req UserReqBody) (string, string, int64) {
	return req.Message.Text, req.Message.Chat.NAME, req.Message.Chat.ID
}
