package main

import (
	"teleBot/telebot"

	"github.com/gin-gonic/gin"
)

func main() {
	telebot.SetToken("")
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		rq := telebot.UserReqBody{}
		c.BindJSON(&rq)
		msg, uname, chatId := telebot.ExtractDataFromUser(rq)
		messege := msg + " from " + uname
		telebot.ReplyMessege(chatId, messege)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
