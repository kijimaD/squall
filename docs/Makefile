.DEFAULT_GOAL := help

.PHONY: merger
merger: ## mergerを実行する
	./scripts/merger.sh

.PHONY: tbls
tbls: ## tblsを実行する
	./scripts/tbls.sh

.PHONY: gengo
gengo: ## OpenAPI定義からGoコードを生成する
	./scripts/gen-go-api.sh

.PHONY: gents
gents: ## OpenAPI定義からTypeScriptコードを生成する
	./scripts/gen-ts-front.sh

.PHONY: help
help: ## ヘルプ表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
