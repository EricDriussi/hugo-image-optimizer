.DEFAULT_GOAL := help

MAKE_PARAMS := $(filter-out $@,$(MAKECMDGOALS))
FILE_PATH := $(word 2,$(MAKE_PARAMS))

.PHONY: help
help: ## This help menu
	@grep -E '^\S+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "%-30s %s\n", $$1, $$2}'

.PHONY: compile
compile: ## Compile binary
	@go build -o optimize

.PHONY: install
install: ## Local install
	@go env -w GOBIN=$$HOME/.local/bin
	@go install
	@mv $$HOME/.local/bin/hugo-image-optimizer $$HOME/.local/bin/optimize

.PHONY: setup
setup: ## Setup dev env
	@go install github.com/vektra/mockery/v2@latest
	@go get

.PHONY: mocks
mocks: ## Create mocks
	@mockery -r --case=snake --outpkg=mocks --output=test/mocks --name=PostRepository

.PHONY: test
test: ## Run tests
ifdef FILE_PATH
	@$(call COLOR_COMMAND, go test $(FILE_PATH))
else
	@$(call COLOR_COMMAND, go test ./...)
endif

.PHONY: test-watch
test-watch: ## Run tests in watch mode
ifdef FILE_PATH
	@$(call COLOR_COMMAND, go test $(FILE_PATH))
	@while true; do \
		inotifywait -qq -r -e create,modify,move,delete ./; \
		printf "\n[ . . . Re-running command . . . ]\n"; \
		$(call COLOR_COMMAND, go test $(FILE_PATH)); \
	done
else
	@$(call COLOR_COMMAND, go test ./...)
	@while true; do \
		inotifywait -qq -r -e create,modify,move,delete ./; \
		printf "\n[ . . . Re-running command . . . ]\n"; \
		$(call COLOR_COMMAND, go test ./...); \
	done
endif


# Color Command
PASS_COLOR=$(shell echo -e "\e[1;32m")
FAIL_COLOR=$(shell echo -e "\e[1;31m")
RESET_COLOR=$(shell echo -e "\e[0m")

COLORED_PASS_TERMS=✅ $(PASS_COLOR)&$(RESET_COLOR)
COLORED_FAIL_TERMS=❌ $(FAIL_COLOR)&$(RESET_COLOR)
COLOR_COMMAND = $(1) | sed -Ee "s/\<pass\>|\<ok\>/$(COLORED_PASS_TERMS)/I" -Ee "s/\<fail\>|\<failed\>/$(COLORED_FAIL_TERMS)/I"
