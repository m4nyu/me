#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
INFRA_DIR="$SCRIPT_DIR/../app/infra"

echo "1) deploy"
echo "2) update"
echo "3) destroy"
echo "4) exit"
echo ""
read -p "choice: " OPERATION

case $OPERATION in
    1) ACTION="deploy" ;;
    2) ACTION="update" ;;
    3) ACTION="destroy" ;;
    4) exit 0 ;;
    *) echo "invalid"; exit 1 ;;
esac

check_command() {
    if ! command -v "$1" &> /dev/null; then
        echo "$1 not installed"
        exit 1
    fi
}

check_command "pulumi"
check_command "ssh"

if [ "$ACTION" = "deploy" ]; then
    if [ ! -f ~/.ssh/id_ed25519 ] && [ ! -f ~/.ssh/id_rsa ]; then
        read -p "generate ssh key? (y/n): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            ssh-keygen -t ed25519 -f ~/.ssh/id_ed25519 -N ""
        else
            exit 1
        fi
    fi
fi

read -p "stack (prod/staging): " STACK
if [ "$STACK" != "prod" ] && [ "$STACK" != "staging" ]; then
    echo "invalid stack"
    exit 1
fi

cd "$INFRA_DIR"

if [ "$ACTION" = "deploy" ]; then
    pulumi stack select $STACK 2>/dev/null || pulumi stack init $STACK

    read -p "vps host: " VPS_HOST
    read -p "vps user: " VPS_USER
    VPS_USER=${VPS_USER:-debian}

    read -p "cloudflare token: " CF_TOKEN
    read -p "cloudflare zone: " CF_ZONE
    read -p "domain: " DOMAIN

    read -p "git repo: " GIT_REPO
    GIT_REPO=${GIT_REPO:-m4nyu/me}

    read -p "git branch: " GIT_BRANCH
    GIT_BRANCH=${GIT_BRANCH:-main}

    pulumi config set vpsHost "$VPS_HOST"
    pulumi config set vpsUser "$VPS_USER"
    pulumi config set --secret cloudflareApiToken "$CF_TOKEN"
    pulumi config set cloudflareZoneId "$CF_ZONE"
    pulumi config set domain "$DOMAIN"
    pulumi config set gitRepo "$GIT_REPO"
    pulumi config set gitBranch "$GIT_BRANCH"
    pulumi config set --secret cloudflare:apiToken "$CF_TOKEN"

    echo ""
    pulumi preview
    echo ""

    read -p "deploy? (y/n): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 0
    fi

    pulumi up -y
    echo ""
    pulumi stack output

elif [ "$ACTION" = "update" ]; then
    pulumi stack select $STACK

    echo ""
    pulumi preview
    echo ""

    read -p "update? (y/n): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 0
    fi

    pulumi up -y
    echo ""
    pulumi stack output

elif [ "$ACTION" = "destroy" ]; then
    pulumi stack select $STACK

    echo ""
    pulumi preview --destroy
    echo ""

    read -p "destroy? (y/n): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 0
    fi

    pulumi destroy -y
fi
