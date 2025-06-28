FROM golang:1.24.4-bookworm

WORKDIR /app
COPY . .

ENTRYPOINT ["go", "run", "cmd/music_share_api/main.go"]