package server

import (
	"bytes"
	"com-line-bot/commands"
	"com-line-bot/utils"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Bot struct {
	bot *linebot.Client
}

type FoodList struct {
	Foods []string `json:"foods"`
}

func (b *Bot) New() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	b.bot = bot
}

func (b *Bot) Callback(w http.ResponseWriter, r *http.Request) {
	events, err := b.bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				content := message.Text

				rand.Seed(time.Now().UnixNano())

				if strings.Contains(content, " vs ") && strings.HasPrefix(content, "!") {
					slice := strings.Split(content[1:], " vs ")
					b.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(commands.Select(slice))).Do()
					break
				} else if strings.HasPrefix(content, "!선택") {
					b.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("TODO"))
				}

				switch content {
				case "!뭐먹지":
				case "ㅁㅁㅈ":
					resp, err := http.Get("https://raw.githubusercontent.com/BeLeap/com-line-bot/main/resources/foodlist.json")
					if err != nil {
						log.Print(err)
						b.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("리스트를 불러오지 못했어요.")).Do()
						break
					}
					defer resp.Body.Close()

					var foodListJson FoodList
					buf := new(bytes.Buffer)
					buf.ReadFrom(resp.Body)
					respByte := buf.Bytes()
					if err := json.Unmarshal(respByte, &foodListJson); err != nil {
						log.Print(err)
						b.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("리스트를 해석하지 못했어요.")).Do()
						break
					}

					foodList := foodListJson.Foods

					if _, err := b.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(foodList[utils.Random(0, len(foodList))])).Do(); err != nil {
						log.Print(err)
					}
					break
				}
			}
		}
	}
}
