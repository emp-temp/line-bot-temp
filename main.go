package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	bot, err := linebot.New(
		"0f1262d7af126fabd662b56092f21f4f",
		"47A0zmbI64hjGIWbDvg4aAmpxewr3cMwXo9g0pe/FKOYkXlBM/INxLOj0vjojjxcXx8hL6pvTW4EyKjt9VkcQgwyiwHSwgtw0Ks1DIROzBhlkyvr2heIeMuRuJScqhuVfWlRSd+bc/OIuCxP1kXF4QdB04t89/1O/w1cDnyilFU=",
	)

	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.POST("/callback", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Print(err)
			}
			return
		}

		var replyText string
		replyText = "可愛い"

		var response string
		response = "ありがとう!!"

		var replySticker string
		replySticker = "おはよう"

		responseSticker := linebot.NewStickerMessage("11537", "52002757")

		var replyImage string
		replyImage = "猫"

		responseImage := linebot.NewImageMessage("https://i.gyazo.com/2db8f85c496dd8f21a91eccc62ceee05.jpg", "https://i.gyazo.com/2db8f85c496dd8f21a91eccc62ceee05.jpg")

		var replyLocation string
		replyLocation = "ディズニー"

		responseLocation := linebot.NewLocationMessage("東京ディズニーランド", "千葉県浦安市舞浜", 35.632896, 139.880394)

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					replyMessage := message.Text
					if strings.Contains(replyMessage, replyText) {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(response)).Do()
					} else if strings.Contains(replyMessage, replySticker) {
						bot.ReplyMessage(event.ReplyToken, responseSticker)
					} else if strings.Contains(replyMessage, replyImage) {
						bot.ReplyMessage(event.ReplyToken, responseImage).Do()
					} else if strings.Contains(replyMessage, replyLocation) {
						bot.ReplyMessage(event.ReplyToken, responseLocation).Do()
					}
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
					if err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
}
