FROM golang:1.23-alpine AS builder

WORKDIR /song-app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/song-app .



FROM alpine AS runner

COPY --from=builder /usr/local/bin/song-app /
COPY .env .env
COPY ./migrations ./migrations

CMD ["/song-app"]
