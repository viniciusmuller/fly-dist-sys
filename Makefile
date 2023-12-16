GO=go
MAELSTROM=maelstrom

build_echo: echo/main.go
	cd echo; $(GO) build .;

echo: build_echo
	$(MAELSTROM) test -w echo --bin ./echo/echo --node-count 1 --time-limit 10
