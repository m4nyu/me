#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$SCRIPT_DIR/../.."
INFRA_DIR="$PROJECT_ROOT/src/infra"

if [ -f "$PROJECT_ROOT/.env" ]; then
    source "$PROJECT_ROOT/.env"
fi

echo "======================================"
echo "  Production Deployment Script"
echo "======================================"
echo ""
echo "⚠️  WARNING: This deploys to PRODUCTION"
echo "This will deploy to your live domain with"
echo "full Cloudflare CDN, security, and firewall."
echo ""

check_command() {
    if ! command -v "$1" &> /dev/null; then
        echo "Error: $1 is not installed"
        exit 1
    fi
}

check_command "pulumi"
check_command "ssh"

if [ ! -f ~/.ssh/id_ed25519 ] && [ ! -f ~/.ssh/id_rsa ]; then
    echo "No SSH key found."
    read -p "Generate SSH key? (y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        ssh-keygen -t ed25519 -f ~/.ssh/id_ed25519 -N ""
        echo "✓ SSH key generated"
    else
        echo "Error: SSH key required for deployment"
        exit 1
    fi
fi

cd "$INFRA_DIR"

pulumi stack select prod 2>/dev/null || {
    echo "Creating new production stack..."
    pulumi stack init prod
}

VPS_HOST=${VPS_HOST:-158.69.218.225}
VPS_USER=${VPS_USER:-debian}
DOMAIN=${DOMAIN:-m4nuel.net}
GIT_REPO=${GIT_REPO:-m4nyu/me}
GIT_BRANCH=${GIT_BRANCH:-main}
CF_ZONE=${CLOUDFLARE_ZONE_ID}

if [ -z "$CLOUDFLARE_API_TOKEN" ]; then
    echo ""
    read -p "Cloudflare API Token: " CLOUDFLARE_API_TOKEN
    if [ -z "$CLOUDFLARE_API_TOKEN" ]; then
        echo "Error: Cloudflare API Token is required for production"
        exit 1
    fi
fi

if [ -z "$CF_ZONE" ]; then
    echo ""
    read -p "Cloudflare Zone ID: " CF_ZONE
    if [ -z "$CF_ZONE" ]; then
        echo "Error: Cloudflare Zone ID is required for production"
        exit 1
    fi
fi

echo ""
echo "Configuration:"
echo "  VPS Host: $VPS_HOST"
echo "  VPS User: $VPS_USER"
echo "  Domain: $DOMAIN"
echo "  Git Repo: $GIT_REPO"
echo "  Git Branch: $GIT_BRANCH"
echo "  Cloudflare Zone: $CF_ZONE"

pulumi config set vpsHost "$VPS_HOST"
pulumi config set vpsUser "$VPS_USER"
pulumi config set domain "$DOMAIN"
pulumi config set --secret cloudflareApiToken "$CLOUDFLARE_API_TOKEN"
pulumi config set cloudflareZoneId "$CF_ZONE"
pulumi config set --secret cloudflare:apiToken "$CLOUDFLARE_API_TOKEN"
pulumi config set gitRepo "$GIT_REPO"
pulumi config set gitBranch "$GIT_BRANCH"

echo ""
echo "✓ Configuration applied"
echo ""
echo "Preview changes..."
echo ""

pulumi preview

echo ""
echo "⚠️  WARNING: You are about to deploy to PRODUCTION"
echo "Domain: $DOMAIN"
echo "VPS: $VPS_HOST"
echo ""
read -p "Deploy production? (y/n): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Deployment cancelled"
    exit 0
fi

echo ""
echo "Deploying production..."
pulumi up -y

echo ""
echo "======================================"
echo "  Production Deployment Complete!"
echo "======================================"
echo ""
pulumi stack output
echo ""
echo "Your site should be live at:"
echo "  https://$DOMAIN"
echo "  https://www.$DOMAIN"
echo ""
echo "Cloudflare CDN, security, and firewall configured."
echo "Container running on port 80."
echo ""
