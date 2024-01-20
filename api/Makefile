.DEFAULT_GOAL := help

.PHONY: test
test: ## テストを実行する
	SQUALL_SILENT_MODE=true \
	SQUALL_APP_ENV=testing \
	SQUALL_LOG_LEVEL=error \
	go test -race -shuffle=on -v ./...

.PHONY: lint
lint: ## lintを実行する
	golangci-lint run --fix --timeout=5m

.PHONY: help
help: ## ヘルプ表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
