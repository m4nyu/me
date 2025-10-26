.PHONY: help install preview deploy update destroy clean config ssh logs

help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

install: ## Install Go dependencies
	go mod download

config: ## Show current stack configuration
	@pulumi config
	@echo "\nStack info:"
	@pulumi stack

preview: ## Preview infrastructure changes
	pulumi preview

deploy: ## Deploy infrastructure
	pulumi up

update: ## Update deployment (refresh and deploy)
	pulumi up --refresh

destroy: ## Destroy infrastructure
	@echo "Warning: This will remove Cloudflare DNS records"
	@read -p "Are you sure? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		pulumi destroy; \
	fi

ssh: ## SSH into VPS
	@HOST=$$(pulumi config get vpsHost 2>/dev/null || echo "158.69.218.225"); \
	USER=$$(pulumi config get vpsUser 2>/dev/null || echo "debian"); \
	KEY=$$(pulumi config get sshKeyPath 2>/dev/null || echo "~/.ssh/id_rsa"); \
	ssh -i $$KEY $$USER@$$HOST

logs: ## View application logs on VPS
	@HOST=$$(pulumi config get vpsHost 2>/dev/null || echo "158.69.218.225"); \
	USER=$$(pulumi config get vpsUser 2>/dev/null || echo "debian"); \
	KEY=$$(pulumi config get sshKeyPath 2>/dev/null || echo "~/.ssh/id_rsa"); \
	ssh -i $$KEY $$USER@$$HOST "docker logs -f me-app"

status: ## Check application status
	@HOST=$$(pulumi config get vpsHost 2>/dev/null || echo "158.69.218.225"); \
	USER=$$(pulumi config get vpsUser 2>/dev/null || echo "debian"); \
	KEY=$$(pulumi config get sshKeyPath 2>/dev/null || echo "~/.ssh/id_rsa"); \
	ssh -i $$KEY $$USER@$$HOST "docker ps && echo '\n--- Nginx Status ---' && sudo systemctl status nginx --no-pager"

outputs: ## Show deployment outputs
	pulumi stack output

clean: ## Clean up local build artifacts
	@rm -rf bin/
	@echo "Cleaned build artifacts"

.DEFAULT_GOAL := help
