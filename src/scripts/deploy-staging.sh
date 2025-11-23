#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
INFRA_DIR="$SCRIPT_DIR/../infra"

echo "======================================"
echo "  Staging Deployment Script"
echo "======================================"
echo ""
echo "This deploys a private staging environment"
echo "accessible only via SSH tunnel on port 8080."
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

pulumi stack select staging 2>/dev/null || {
    echo "Creating new staging stack..."
    pulumi stack init staging
}

echo ""
echo "Configuration:"
echo "-------------"

read -p "VPS Host [158.69.218.225]: " VPS_HOST
VPS_HOST=${VPS_HOST:-158.69.218.225}

read -p "VPS User [debian]: " VPS_USER
VPS_USER=${VPS_USER:-debian}

read -p "Git Repo [m4nyu/me]: " GIT_REPO
GIT_REPO=${GIT_REPO:-m4nyu/me}

read -p "Git Branch [main]: " GIT_BRANCH
GIT_BRANCH=${GIT_BRANCH:-main}

pulumi config set vpsHost "$VPS_HOST"
pulumi config set vpsUser "$VPS_USER"
pulumi config set gitRepo "$GIT_REPO"
pulumi config set gitBranch "$GIT_BRANCH"

echo ""
echo "✓ Configuration saved"
echo ""
echo "Preview changes..."
echo ""

pulumi preview

echo ""
read -p "Deploy staging? (y/n): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Deployment cancelled"
    exit 0
fi

echo ""
echo "Deploying staging..."
pulumi up -y

echo ""
echo "======================================"
echo "  Staging Deployment Complete!"
echo "======================================"
echo ""
pulumi stack output
echo ""
echo "Access staging via SSH tunnel:"
echo "  ssh -L 8080:localhost:8080 $VPS_USER@$VPS_HOST"
echo ""
echo "Then visit: http://localhost:8080"
echo ""
echo "Note: Staging is NOT accessible from the internet"
echo ""
