.PHONY: build clean

build:
	go build -o bin/subwords

clean:
	rm -f bin/subwords
