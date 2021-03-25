.PHONY = all api clean

.DEFAULT_GOAL = all

clean:
	rm -f bin/lookum

api:
	go build -o bin/lookum cmd/lookum/main.go

all: api