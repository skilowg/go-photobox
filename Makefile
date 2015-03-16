FLAGS=GOARM=5 GOARCH=arm GOOS=linux

gophotobox:
	$(FLAGS) go build
