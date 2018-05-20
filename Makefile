build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/remote-area remote-area/main.go
deploy:
	make
	npx serverless deploy
remove:
	npx serverless remove
