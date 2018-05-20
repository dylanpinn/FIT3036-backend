build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/remote/area remote/area.go
deploy:
	make
	yarn serverless deploy
remove:
	yarn serverless remove
