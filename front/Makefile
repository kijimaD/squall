.DEFAULT_GOAL := help

.PHONY: run
run: ## 実行する
	@npm run start

.PHONY: fmt
fmt: ## format
	@npx --yes prettier@^3 --write "**/*.ts"
	@npx --yes prettier@^3 --write "**/*.tsx"

.PHONY: help
help: ## ヘルプ表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
