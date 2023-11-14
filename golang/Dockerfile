# 開発用ステージ
FROM golang:1.20 as development

## コンテナの最初の作業ディレクトリ
WORKDIR /app

## ホストのファイルをコンテナの作業ディレクトリにコピー
COPY go.mod ./
COPY go.sum ./

## go.mod ファイルに記載されている依存関係を解決し、ダウンロードしてコンテナ内に配置
RUN go mod download
COPY . .
RUN go install github.com/cosmtrek/air@latest

## コンテナが起動した時に実行されるコマンド
CMD ["air"]

# 本番用ステージ
FROM alpine:latest as production

## コンテナの最初の作業ディレクトリ
WORKDIR /root/

## ホストのファイルをコンテナの作業ディレクトリにコピー
COPY --from=development /app/main .

## コンテナが起動した時に実行されるコマンド
CMD ["./main"]
