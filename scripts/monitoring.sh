#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
INFRA_DIR="$SCRIPT_DIR/../app/infra"

read -p "stack (prod/staging): " STACK
if [ "$STACK" != "prod" ] && [ "$STACK" != "staging" ]; then
    echo "invalid stack"
    exit 1
fi

cd "$INFRA_DIR"

pulumi stack select $STACK 2>/dev/null

VPS_HOST=$(pulumi config get vpsHost 2>/dev/null)
VPS_USER=$(pulumi config get vpsUser 2>/dev/null)

if [ -z "$VPS_HOST" ] || [ -z "$VPS_USER" ]; then
    echo "stack not configured"
    exit 1
fi

echo "tunnel: http://localhost:19999"
echo ""

ssh -L 19999:localhost:19999 "$VPS_USER@$VPS_HOST"
