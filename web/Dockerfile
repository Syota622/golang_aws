# ビルド用ステージ
FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

## 全てのGoファイルをコンパイルする
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

## 開発用ステージ
FROM golang:1.20 as development

WORKDIR /app

## 実行ファイルを/app/にコピーする
COPY --from=builder /app/main .
RUN go install github.com/cosmtrek/air@latest

CMD ["air"]

# 本番用ステージ
FROM alpine:latest as production

WORKDIR /root/

## 実行ファイルを/root/にコピーする
COPY --from=builder /app/main .

## viewディレクトリを/app/viewにコピーする
COPY --from=builder /app/view /app/view

CMD ["./main"]