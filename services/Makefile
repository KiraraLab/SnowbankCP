

build:
	go build -o ../build/main .

build-platform:
  CGO_ENABLED=0 && GOOS=darwin && GOARCH=arm64 &&	make build 

run:
	# apply .env to the environment
	source ../.env.example && \
	go run .

clean:
	rm -rf ../build/



.PHONY: clean build build-platform run