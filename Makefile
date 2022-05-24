build:
	go build -o bin/main src/*.go

run: build
	./bin/main

clean:
	rm -rf bin/*