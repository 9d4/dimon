.PHONY: start
start: | build runbin

.PHONY: build
build:
	go build .

.PHONY: runbin
runbin:
	sudo ./dimon