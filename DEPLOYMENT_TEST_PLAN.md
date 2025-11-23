# Deployment Test Plan

## Overview

This document outlines the testing plan for the automated deployment infrastructure.

## Completed: Local Dry Runs ✅

### Production Stack Preview
```bash
cd src/infra
export PULUMI_CONFIG_PASSPHRASE='<your-passphrase-from-.env>'
export PULUMI_ACCESS_TOKEN='<your-access-token-from-.env>'
pulumi stack select prod
pulumi preview
```

**Result:** ✅ Success
- 11 resources to create (+ 3 unchanged)
- Cloudflare DNS, CDN, security, firewall all configured
- Docker container on port 80
- Some deprecation warnings (resources still supported until Jan 15, 2025)

### Staging Stack Preview
```bash
cd src/infra
export PULUMI_CONFIG_PASSPHRASE='<your-passphrase-from-.env>'
export PULUMI_ACCESS_TOKEN='<your-access-token-from-.env>'
pulumi stack select staging
pulumi preview
```

**Result:** ✅ Success
- 4 resources to create
- No Cloudflare infrastructure (as designed)
- Docker container on localhost:8080
- Minimal configuration

## Next: GitHub Actions Testing

### Prerequisites

Before testing with GitHub Actions, ensure these secrets are configured:

1. Go to: https://github.com/m4nyu/me/settings/secrets/actions
2. Add the following secrets:

| Secret Name | Value | Source |
|-------------|-------|--------|
| `VPS_SSH_KEY` | Private SSH key (ed25519) | Generate new or use existing |
| `VPS_HOST` | `158.69.218.225` | Static value |
| `PULUMI_CONFIG_PASSPHRASE` | Your passphrase | From `src/infra/.env` |
| `PULUMI_ACCESS_TOKEN` | Your access token | From `src/infra/.env` |

### SSH Key Setup

#### Option 1: Generate New Key for GitHub Actions
```bash
ssh-keygen -t ed25519 -C "github-actions@m4nyu-me" -f ~/.ssh/github_deploy
cat ~/.ssh/github_deploy  # Copy this to VPS_SSH_KEY secret

# Add public key to VPS
ssh debian@158.69.218.225
echo "ssh-ed25519 AAAA... github-actions@m4nyu-me" >> ~/.ssh/authorized_keys
exit
```

#### Option 2: Use Existing Key
```bash
cat ~/.ssh/id_ed25519  # Copy entire private key to VPS_SSH_KEY secret
```

### Test 1: Staging Deployment (PR)

**Purpose:** Test that staging deploys correctly when a PR is created

**Steps:**
1. Create a PR from `feat/cicd` to `main`:
   - Visit: https://github.com/m4nyu/me/pull/new/feat/cicd
   - Title: "Test: CI/CD Staging Deployment"
   - Create pull request (DO NOT MERGE YET)

2. Monitor the workflow:
   - Go to: https://github.com/m4nyu/me/actions
   - Watch the "Deploy Staging" workflow

**Expected Results:**
- ✅ Workflow completes successfully
- ✅ Comment posted on PR with SSH tunnel instructions
- ✅ Container running on VPS: `ssh debian@158.69.218.225 'docker ps | grep staging'`
- ✅ Staging accessible via SSH tunnel:
  ```bash
  ssh -L 8080:localhost:8080 debian@158.69.218.225
  # Visit http://localhost:8080 in browser
  ```
- ✅ Staging NOT accessible from internet: `curl http://158.69.218.225:8080` (should fail)

**Troubleshooting:**
- If SSH fails: Check VPS_SSH_KEY is properly formatted
- If Pulumi fails: Check PULUMI_CONFIG_PASSPHRASE and PULUMI_ACCESS_TOKEN
- If container fails: Check VPS has enough resources

### Test 2: Production Deployment (Merge to Main)

**Purpose:** Test that production deploys correctly when PR is merged

**⚠️ WARNING:** This will deploy to PRODUCTION (m4nuel.net)

**Steps:**
1. Verify staging deployment works correctly (Test 1)
2. Review all changes one final time
3. Merge the PR to `main`
4. Monitor the workflow:
   - Go to: https://github.com/m4nyu/me/actions
   - Watch the "Deploy Production" workflow

**Expected Results:**
- ✅ Workflow completes successfully
- ✅ All Cloudflare resources created/updated
- ✅ Firewall rules applied to VPS
- ✅ Container running on VPS: `ssh debian@158.69.218.225 'docker ps | grep prod'`
- ✅ Production container on port 80: `ssh debian@158.69.218.225 'curl localhost:80'`
- ✅ Site accessible via Cloudflare: `curl https://m4nuel.net`
- ✅ Site accessible via www: `curl https://www.m4nuel.net`
- ✅ HTTP redirects to HTTPS: `curl -I http://m4nuel.net` (should see 301/302)

**Verification Commands:**
```bash
# Check containers
ssh debian@158.69.218.225 'docker ps'

# Check prod container logs
ssh debian@158.69.218.225 'docker logs prod --tail 50'

# Check staging container logs
ssh debian@158.69.218.225 'docker logs staging --tail 50'

# Test production locally from VPS
ssh debian@158.69.218.225 'curl -I localhost:80'

# Test staging locally from VPS
ssh debian@158.69.218.225 'curl -I localhost:8080'

# Check firewall rules
ssh debian@158.69.218.225 'sudo ufw status verbose'
```

### Test 3: Subsequent PR (Staging Update)

**Purpose:** Test that staging updates correctly on subsequent PRs

**Steps:**
1. After merging to main, create a new branch with a small change
2. Create a PR to main
3. Verify staging deployment updates

**Expected Results:**
- ✅ Staging container recreated with new changes
- ✅ Production container unaffected
- ✅ Both containers running simultaneously

## Rollback Plan

If production deployment fails or causes issues:

### Option 1: Destroy via Pulumi
```bash
cd src/infra
export PULUMI_CONFIG_PASSPHRASE='<your-passphrase-from-.env>'
export PULUMI_ACCESS_TOKEN='<your-access-token-from-.env>'
pulumi stack select prod
pulumi destroy
```

### Option 2: Manual Container Stop
```bash
ssh debian@158.69.218.225 'docker stop prod && docker rm prod'
```

### Option 3: Revert Git Commit
```bash
git revert HEAD
git push origin main
# Wait for workflow to redeploy previous version
```

## Success Criteria

All tests pass when:
- ✅ Local dry runs succeed (production & staging)
- ✅ Staging deploys on PR creation
- ✅ Staging is NOT accessible from internet
- ✅ Staging is accessible via SSH tunnel
- ✅ Production deploys on merge to main
- ✅ Production is accessible at https://m4nuel.net
- ✅ Cloudflare CDN, security, and firewall configured
- ✅ Both staging and production can run simultaneously
- ✅ Subsequent PRs update staging without affecting production

## Current Status

- [x] Create feat/cicd branch
- [x] Implement infrastructure code changes
- [x] Test production deployment locally (dry run)
- [x] Test staging deployment locally (dry run)
- [x] Commit and push changes
- [ ] Configure GitHub secrets
- [ ] Test staging deployment via GitHub Actions (PR)
- [ ] Test production deployment via GitHub Actions (merge)
- [ ] Test subsequent staging deployments

## Notes

- Both containers run on same VPS (158.69.218.225)
- Port 80: Production (public, proxied via Cloudflare)
- Port 8080: Staging (localhost-only, SSH tunnel access)
- Cloudflare deprecation warnings are expected (still supported until Jan 15, 2025)
- Consider updating to `cloudflare_ruleset` in future iteration
