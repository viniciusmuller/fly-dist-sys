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

build_broadcast: broadcast/main.go
	cd broadcast; $(GO) build .;

test_broadcast: build_broadcast
	$(MAELSTROM) test -w broadcast --bin ./broadcast/broadcast --node-count 5 --time-limit 20 --rate 10
