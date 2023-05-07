build:
	go build -o ./bin/main .

start: build
	./bin/main

package: build
	fyne package -os darwin -icon icon/icon.jpeg