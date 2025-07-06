

build:
	npm run build --prefix ./ui/webpages
	go mod tidy
	go build -o ./build/main ./cmd/cli/*

build-platform:
  CGO_ENABLED=0 && GOOS=darwin && GOARCH=arm64 &&	make build 

run:
# npm run dev --prefix ./internal/services/webpages && \
# apply .env to the environment 
#	source ../.env.example &&
	go run ./cmd/cli/*

clean:
	rm -rf ../build/



.PHONY: clean build build-platform run