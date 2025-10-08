.PHONY: init preview deploy destroy refresh outputs help

init: ## Initialize Pulumi stack
	pulumi stack init dev || pulumi stack select dev

preview: ## Preview infrastructure changes
	pulumi preview

deploy: ## Deploy infrastructure
	pulumi up

destroy: ## Destroy all infrastructure
	pulumi destroy

refresh: ## Refresh infrastructure state
	pulumi refresh

outputs: ## Show stack outputs
	pulumi stack output

help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
