# ND Torrent bot
Author Tsykin V.A. aka NektoDev
## Description

Simple telegram bot written on Golang that add torrent to transmission by magnet or rutracker link. 

## Quickstart

1. Build container:
```
docker build -t nektodev/nd_torrent_bot:latest .
```

2. _(optional)_ Push container:
docker push nektodev/nd_torrent_bo:latest

3. Run container:
```
sudo docker run --restart=always --name=nd_torrent_bot -d \
    -e BOT_TG_TOKEN= \
    -e BOT_ATTT_URL= \
    nektodev/nd_torrent_bot
```

