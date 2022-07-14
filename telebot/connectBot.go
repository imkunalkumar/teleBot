package telebot

import (
	"errors"
)

var token string

func SetToken(tkn string) {
	token = tkn
}

func GetUrl() (string, error) {
	if token != "" {
		return "https://api.telegram.org/bot" + token + "/sendMessage", nil
	} else {
		return "", errors.New("token not found")
	}

}
