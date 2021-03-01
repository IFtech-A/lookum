.PHONY = all build clean

.DEFAULT_GOAL = all

clean:
	rm -f bin/lookum

build:
	go build -o lookum cmd/main.go

all: build