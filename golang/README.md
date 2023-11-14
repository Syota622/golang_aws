# web-service-gin
ginとmysqlを使った簡単なwebアプリケーションです。

## mysql
ホストでport 3306で起動してください。

## 環境変数
mysqlを使用するために以下の環境変数を設定してください
- DBUSER
- DBPASS

## database
- myappでdatabaseを作成してください。
- sqlディレクトリ以下にテーブル作成とseedデータ作成のsqlファイルを作成しているので実行してください

## docker
docker化はしていないため、dockerを使用する場合は適宜コードを書き換えてください。

# 環境構築方法

# docker compose
docker compose build
docker compose up

# curl
## アルバムのリストを取得
curl http://localhost:8080/albums

```sh
(base) suzukishouta@shota golang-docker % curl http://localhost:8080/albums          
[
    {
        "id": 1,
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
    },
    {
        "id": 2,
        "title": "Giant Steps",
        "artist": "John Coltrane",
        "price": 63.99
    },
    {
        "id": 3,
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
    },
    {
        "id": 4,
        "title": "Sarah Vaughan",
        "artist": "Sarah Vaughan",
        "price": 34.98
    },
    {
        "id": 5,
        "title": "新しいアルバム",
        "artist": "アーティスト名",
        "price": 99.99
    }
]
```

## 新しいアルバムを追加
curl -X POST http://localhost:8080/albums \
     -H "Content-Type: application/json" \
     -d '{"title": "新しいアルバム", "artist": "アーティスト名", "price": 99.99}'

```sh
(base) suzukishouta@shota golang-docker % curl -X POST http://localhost:8080/albums \
     -H "Content-Type: application/json" \
     -d '{"title": "新しいアルバム", "artist": "アーティスト名", "price": 99.99}'
{
    "message": {
        "Number": 1146,
        "SQLState": [
            52,
            50,
            83,
            48,
            50
        ],
        "Message": "Table 'myapp.album' doesn't exist"
    }
}
```

## アルバムの詳細を取得
curl http://localhost:8080/albums/5

```sh
(base) suzukishouta@shota golang-docker % curl http://localhost:8080/albums/5
{
    "id": 5,
    "title": "新しいアルバム",
    "artist": "アーティスト名",
    "price": 99.99
}
```

## アルバムの更新
curl -X PATCH http://localhost:8080/albums/5 \
     -H "Content-Type: application/json" \
     -d '{"title": "更新されたタイトル", "artist": "更新されたアーティスト名", "price": 199.99}'

```sh
(base) suzukishouta@shota golang-docker % curl -X PATCH http://localhost:8080/albums/5 \
     -H "Content-Type: application/json" \
     -d '{"title": "更新されたタイトル", "artist": "更新されたアーティスト名", "price": 199.99}'
(base) suzukishouta@shota golang-docker % curl http://localhost:8080/albums/5
{
    "id": 5,
    "title": "更新されたタイトル",
    "artist": "更新されたアーティスト名",
    "price": 199.99
}
```

## アルバムの削除
curl -X DELETE http://localhost:8080/albums/5
```
(base) suzukishouta@shota golang-docker % curl -X DELETE http://localhost:8080/albums/5
(base) suzukishouta@shota golang-docker % curl http://localhost:8080/albums/5
{
    "message": {}
}
```
