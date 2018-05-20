build:
	rm -rf bin/*
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/remote/area remote/area/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/remote/road remote/road/main.go
deploy:
	make
	yarn serverless deploy
remove:
	yarn serverless remove
