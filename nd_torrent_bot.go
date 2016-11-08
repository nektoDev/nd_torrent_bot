package main

import (
	"flag"
	"log"
	"os"

	"github.com/coreos/pkg/flagutil"
	"gopkg.in/telegram-bot-api.v4"
	"strings"
	"strconv"
	"gopkg.in/resty.v0"
	"encoding/json"
)

type torrent struct {
	Magnet      string `json:"magnet"`
	Url         string `json:"url"`
	DownloadDir string `json:"downloadDir"`
}

func (f *torrent) SetDownloadDir(downloadDir string) {
	f.DownloadDir = downloadDir
}

func main() {
	flags := flag.NewFlagSet("tg-auth", flag.ExitOnError)
	tgtoken := flags.String("tg-token", "", "Telegram token")
	attt_url := flags.String("attt-url", "", "ATTT URL")
	movies_path := "/home/nektodev/Media/Movies/"
	series_path := "/home/nektodev/Media/Series/"

	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "BOT")

	if *tgtoken == "" {
		log.Fatal("Telegram token required")
	}

	tgBot, err := tgbotapi.NewBotAPI(*tgtoken)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on telegram account %s", tgBot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := tgBot.GetUpdatesChan(u)

	var torrents_type_map = make(map[int]torrent)
	i := 0

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		replyMsg := ""

		if (strings.Contains(update.Message.Text, "magnet:")) {
			i++;
			torrents_type_map[i] = torrent{update.Message.Text, "", ""};
			index := strconv.Itoa(i)
			replyMsg = "Got magnet URL. Choose if it is \n/movie_" + index + "\nor\n/series_" + index
		}

		if (strings.Contains(update.Message.Text, "rutracker.org")) {
			i++;
			torrents_type_map[i] = torrent{"", update.Message.Text, ""};
			index := strconv.Itoa(i)
			replyMsg = "Got rutracker URL. Choose if it is \n/movie_" + index + "\nor\n/series_" + index
		}

		if (update.Message.IsCommand()) {
			index := strings.Split(update.Message.Text, "_")[1]
			j, _ := strconv.Atoi(index)
			if tor, ok := torrents_type_map[j]; ok {
				if (strings.Contains(update.Message.Text, "/movie_")) {
					tor.SetDownloadDir(movies_path)
				}
				if (strings.Contains(update.Message.Text, "/series_")) {
					tor.SetDownloadDir(series_path)
				}

				var request [1]torrent
				request[0] = tor
				req, _ := json.Marshal(request)

				resty.R().
					SetHeader("Content-Type", "application/json").
					SetBody(req).
					Post(*attt_url)
				delete(torrents_type_map, j)
			}
		}

		reply := tgbotapi.NewMessage(update.Message.Chat.ID, replyMsg)
		tgBot.Send(reply)

	}

}