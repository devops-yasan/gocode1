.PHONY: test build docker-build deploy

test:
go test ./...

build:
go build ./...

docker-build:
docker build -t ghcr.io/devops-yasan/app:latest .

deploy:
ENVIRONMENT=staging bash scripts/deploy.sh
