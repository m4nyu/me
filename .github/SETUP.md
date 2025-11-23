# GitHub Actions Setup

This document outlines the required GitHub secrets for automated deployments.

## Required Secrets

Configure these secrets in your GitHub repository settings (`Settings` → `Secrets and variables` → `Actions`):

### 1. `VPS_SSH_KEY`
**Description:** Private SSH key (ed25519) for accessing the OVH VPS.

**How to generate:**
```bash
ssh-keygen -t ed25519 -C "github-actions@deploy" -f ~/.ssh/github_deploy
cat ~/.ssh/github_deploy
```

**Value:** The entire private key content including:
```
-----BEGIN OPENSSH PRIVATE KEY-----
...
-----END OPENSSH PRIVATE KEY-----
```

**Setup on VPS:**
```bash
# Add the public key to VPS
ssh debian@158.69.218.225
echo "ssh-ed25519 AAAA... github-actions@deploy" >> ~/.ssh/authorized_keys
```

### 2. `VPS_HOST`
**Description:** IP address of the OVH VPS.

**Value:** `158.69.218.225`

### 3. `PULUMI_CONFIG_PASSPHRASE`
**Description:** Passphrase used to encrypt/decrypt Pulumi stack configurations.

**Value:** Get from `/home/manuel/code/personal/me/src/infra/.env` file (line 1, after `PULUMI_CONFIG_PASSPHRASE=`)

**Note:** This is the same passphrase used locally. Keep it secure.

### 4. `PULUMI_ACCESS_TOKEN`
**Description:** Pulumi Cloud access token for managing infrastructure state.

**How to generate:**
1. Visit https://app.pulumi.com/
2. Go to `Settings` → `Access Tokens`
3. Click `Create token`
4. Copy the token (starts with `pul-`)

**Value:** Get from `/home/manuel/code/personal/me/src/infra/.env` file (line 2, after `PULUMI_ACCESS_TOKEN=`)

## Deployment Workflows

### Production Deployment (`prod.yml`)
- **Trigger:** Push to `main` branch
- **Target:** Production stack (`prod`)
- **Infrastructure:**
  - Cloudflare DNS (m4nuel.net)
  - Cloudflare CDN and security settings
  - UFW firewall (Cloudflare IP whitelist)
  - Docker container on port 80 (public)
- **Result:** Live website at https://m4nuel.net

### Staging Deployment (`staging.yml`)
- **Trigger:** Pull request to `main` branch
- **Target:** Staging stack (`staging`)
- **Infrastructure:**
  - Docker container on port 8080 (localhost-only)
  - No Cloudflare setup (private environment)
  - No firewall changes
- **Result:** Private staging environment accessible only via SSH tunnel

**Access staging:**
```bash
ssh -L 8080:localhost:8080 debian@158.69.218.225
# Then visit: http://localhost:8080
```

## Security Notes

- The staging environment is NOT accessible from the internet
- Only port 80 is exposed publicly (for production via Cloudflare)
- Staging runs on localhost:8080 to prevent accidental public exposure
- All secrets should be rotated if compromised

## Troubleshooting

### SSH connection fails
Check that the public key is added to VPS:
```bash
ssh debian@158.69.218.225 'cat ~/.ssh/authorized_keys'
```

### Pulumi state conflicts
If multiple deployments run simultaneously, Pulumi may lock. Wait for the lock to release or cancel conflicting workflows.

### Container port conflicts
Production (port 80) and staging (port 8080) should not conflict. If they do, check running containers:
```bash
ssh debian@158.69.218.225 'docker ps'
```
