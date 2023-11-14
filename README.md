# ECR

1. ECRリポジトリへログイン
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 235484765172.dkr.ecr.ap-northeast-1.amazonaws.com

2. build：ECRリポジトリへイメージをpush
docker build -t ecr-test .

