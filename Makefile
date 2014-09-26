default: build

build:
	go build

run: build
	./scrape ${ARGS}
