.PHONY: install-maelstrom
install-maelstrom:
	@brew install openjdk graphviz gnuplot
	@curl -L https://github.com/jepsen-io/maelstrom/releases/download/v0.2.3/maelstrom.tar.bz2 -o maelstrom.tar.bz2
	@tar -xjf maelstrom.tar.bz2
	@rm maelstrom.tar.bz2
	@chmod +x maelstrom/maelstrom
	@echo "Maelstrom installed in ./maelstrom directory"

.PHONY: debug-maelstrom
debug-maelstrom:
	@cd maelstrom && ./maelstrom serve

.PHONY: run-echo
run-echo:
	@go build -o echo ./cmd/echo/main.go
	@mv echo maelstrom/
	@cd maelstrom && ./maelstrom test -w echo --bin echo --node-count 1 --time-limit 10
	@cd maelstrom && rm echo
