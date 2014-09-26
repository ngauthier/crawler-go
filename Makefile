default: build

build:
	go build -o crawler

run: build
	./crawler ${ARGS}
