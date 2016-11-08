FROM golang

ADD nd_torrent_bot.go  /go/src/github.com/nektodev/nd_torrent_bot/

RUN go get github.com/coreos/pkg/flagutil && go get gopkg.in/telegram-bot-api.v4 && go get gopkg.in/resty.v0 && go install github.com/nektodev/nd_torrent_bot
ENTRYPOINT /go/bin/nd_torrent_bot