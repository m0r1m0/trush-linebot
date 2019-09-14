package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	garbage "github.com/yuki-wd/kawasaki-garbage"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/job", func(w http.ResponseWriter, req *http.Request) {
		if req.Header.Get("X-Appengine-Cron") == "true" {
			tomorrowWeekday := time.Now().AddDate(0, 0, 1).Weekday()

			garbageInfos := garbage.GetGarbageType(tomorrowWeekday)
			messages := []linebot.SendingMessage{}
			if len(garbageInfos) > 0 {
				for _, info := range garbageInfos {
					messages = append(
						messages,
						linebot.NewTextMessage(fmt.Sprintf("明日は%vの日だよ！", info)),
					)
				}
			} else {
				messages = append(messages, linebot.NewTextMessage("明日のゴミは無いよ！"))
			}
			_, err := bot.PushMessage(os.Getenv("GROUP_ID"), messages...).Do()
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}

}
