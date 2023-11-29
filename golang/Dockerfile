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
# CGO_ENABLED=0 を追加して静的にリンクされたバイナリを生成
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
## コンテナが起動した時に実行されるコマンド（開発用）
CMD ["air"]

# 本番用ステージ
FROM alpine:latest as production
## コンテナの最初の作業ディレクトリ
WORKDIR /root/
## ホストのファイルをコンテナの作業ディレクトリにコピー
COPY --from=development /app/main .
## コンテナが起動した時に実行されるコマンド（本番用）
CMD ["./main"]
