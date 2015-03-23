FLAGS=GOARM=5 GOARCH=arm GOOS=linux
TARGET=gophotobox

gophotobox:
	$(FLAGS) go build -o $(TARGET)

desktop:
	go build -o $(TARGET)

clean:
	rm $(TARGET)

test:
	go test ./...
