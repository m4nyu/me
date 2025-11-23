# Deployment Scripts

This directory contains deployment scripts for the infrastructure.

## Scripts

### `deploy-staging.sh`

Deploys a **private staging environment** for testing.

**Features:**
- Container on `localhost:8080` (not accessible from internet)
- No Cloudflare configuration
- Minimal infrastructure
- SSH tunnel required for access

**Usage:**
```bash
./src/scripts/deploy-staging.sh
```

Or via Make:
```bash
make deploy-staging
```

**Access Staging:**
```bash
ssh -L 8080:localhost:8080 debian@158.69.218.225
# Then visit: http://localhost:8080
```

### `deploy-prod.sh`

Deploys to **production** with full infrastructure.

**Features:**
- Container on port `80` (public access)
- Full Cloudflare setup (DNS, CDN, security, firewall)
- Live at your domain
- ⚠️ **WARNING:** Deploys to production!

**Usage:**
```bash
./src/scripts/deploy-prod.sh
```

Or via Make:
```bash
make deploy-prod
```

**Requirements:**
- Cloudflare API Token
- Cloudflare Zone ID
- Domain name

### `deploy.sh` (Legacy)

Legacy deployment script with interactive menu. This script allows you to:
- Choose between deploy/update/destroy operations
- Select prod or staging stack
- Configure all settings interactively

**Note:** The new `deploy-staging.sh` and `deploy-prod.sh` scripts are recommended for most use cases.

## Configuration

All scripts will:
1. Check for required commands (pulumi, ssh)
2. Prompt for SSH key generation if needed
3. Ask for configuration values (with defaults)
4. Show a preview of changes
5. Confirm before deploying

## Environment Variables

Scripts respect these environment variables if set:
- `PULUMI_CONFIG_PASSPHRASE` - Pulumi encryption passphrase
- `PULUMI_ACCESS_TOKEN` - Pulumi Cloud access token

## Automated Deployments

For automated deployments via GitHub Actions, see:
- `.github/workflows/prod.yml` - Production deployment on merge to main
- `.github/workflows/staging.yml` - Staging deployment on pull requests

These workflows don't use these scripts; they run Pulumi directly with secrets from GitHub.
