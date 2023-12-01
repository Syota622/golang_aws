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

# docker compose
## コンテナ、ネットワーク作成
docker compose build
docker compose up

## コンテナ、ネットワーク削除
docker-compose down

## コンテナ、ネットワーク、ボリューム削除
docker-compose down -v

# ECR Push
## ECRリポジトリへログイン
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 235484765172.dkr.ecr.ap-northeast-1.amazonaws.com

## build：ECRリポジトリへイメージをpush
docker build -t ecr-test .

## イメージをタグ付け
docker tag ecr-test:latest 235484765172.dkr.ecr.ap-northeast-1.amazonaws.com/ecr-test:latest

## イメージをpush
docker push 235484765172.dkr.ecr.ap-northeast-1.amazonaws.com/ecr-test:latest

## ECRのリポジトリを確認
aws ecr list-images --repository-name ecr-test

## ECRのリポジトリの更新
docker build -t ecr-test .
docker tag ecr-test:latest 235484765172.dkr.ecr.ap-northeast-1.amazonaws.com/ecr-test:latest
docker push 235484765172.dkr.ecr.ap-northeast-1.amazonaws.com/ecr-test:latest
