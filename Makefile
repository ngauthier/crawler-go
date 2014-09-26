default: build

build:
	go build -o scrape

run: build
	./scrape ${ARGS}
