# docker ビルド
docker-compose build
docker-compose run --rm web bin/setup
docker-compose run --rm web yarn install
docker-compose up -d
docker-compose run --rm web rails db:create
docker-compose run --rm web bundle exec rails db:migrate
docker-compose exec web bin/rails db:migrate RAILS_ENV=development
http://localhost:3100

# Dockerfile更新
docker-compose build
docker-compose run --rm web bundle install
