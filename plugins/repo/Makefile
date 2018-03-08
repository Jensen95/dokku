include ../../common.mk

GO_ARGS ?= -a

SUBCOMMANDS = subcommands/gc subcommands/purge-cache subcommands/report subcommands/set
TRIGGERS = triggers/install triggers/post-delete triggers/post-extract
build-in-docker: clean
	docker run --rm \
		-v $$PWD/../..:$(GO_REPO_ROOT) \
		-w $(GO_REPO_ROOT)/plugins/repo \
		$(BUILD_IMAGE) \
		bash -c "GO_ARGS='$(GO_ARGS)' make build" || exit $$?

build: commands subcommands triggers triggers-copy

commands: **/**/commands.go
	go build $(GO_ARGS) -o commands src/commands/commands.go

subcommands: $(SUBCOMMANDS)

subcommands/%: src/subcommands/*/%.go
	go build $(GO_ARGS) -o $@ $<

clean:
	rm -rf commands subcommands triggers install post-*

src-clean:
	rm -rf .gitignore src triggers vendor Makefile *.go

triggers: $(TRIGGERS)

triggers/%: src/triggers/*/%.go
	go build $(GO_ARGS) -o $@ $<

triggers-copy:
	cp triggers/* .
