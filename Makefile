.PHONY: dev
dev: | build runbindev

.PHONY: start
start: | build runbin

.PHONY: build
build:
	go build .

.PHONY: runbin
runbin:
	sudo ./dimon

.PHONY: runbindev
runbindev:
	./dimon --socketpath dimon.sock -d dimon.db

.PHONY: install
install: 
	./install.sh