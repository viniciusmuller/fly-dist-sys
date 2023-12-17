GO=go
MAELSTROM=maelstrom

build_echo: echo/main.go
	cd echo; $(GO) build .;

test_echo: build_echo
	$(MAELSTROM) test -w echo --bin ./echo/echo --node-count 1 --time-limit 10

build_unique_ids: unique_ids/main.go
	cd unique_ids; $(GO) build .;

test_unique_ids: build_unique_ids
	$(MAELSTROM) test -w unique-ids --bin ./unique_ids/unique_ids --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition
