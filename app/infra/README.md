# Infrastructure

Oracle Cloud Infrastructure deployment for m4nuel.net

## ▲ Setup

OCI CLI configuration with credentials:

```bash
oci setup config
```

SSH key pair generation for instance access:

```bash
ssh-keygen -t ed25519 -f ~/.ssh/oci
```

Environment file creation from template:

```bash
cp .env.example .env
```

Pulumi authentication:

```bash
pulumi login
```

## ▶ Deploy

### Production

Production stack selection:

```bash
source .env
pulumi stack select prod
```

Infrastructure deployment:

```bash
pulumi up
```

### Staging

Staging stack selection:

```bash
source .env
pulumi stack select staging
```

Infrastructure deployment:

```bash
pulumi up
```

**Automated deployment via GitHub Actions in `.github/workflows/deploy.yml`**

## ▲ Access

Public IP retrieval:

```bash
pulumi stack output publicIp
```

SSH connection to instance:

```bash
ssh -i ~/.ssh/oci ubuntu@$(pulumi stack output publicIp)
```

## ▼ Destroy

Production teardown:

```bash
source .env
pulumi stack select prod
pulumi destroy
```

Staging teardown:

```bash
source .env
pulumi stack select staging
pulumi destroy
```

## DNS Configuration

Namecheap DNS records update after deployment:

```
Type: A
Host: @
Value: <production-public-ip>

Type: A
Host: www
Value: <production-public-ip>
```

SSL certificate automatic provisioning via Let's Encrypt.
