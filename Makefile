all: clip

clip: *.go
	go tool 6g -o clip.6 *.go
	go tool 6l -o clip