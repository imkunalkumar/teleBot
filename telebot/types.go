package telebot

//user req body
type UserReqBody struct {
	Message Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	ID   int64  `json:"id"`
	NAME string `json:"first_name"`
}

// messege body
type SendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}
