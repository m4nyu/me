# Workflows

## Deploy

Infrastructure deployment with automatic retry every 30 minutes until capacity available.

### Required Secrets

Repository secrets configuration (Settings → Secrets and variables → Actions):

**PULUMI_CONFIG_PASSPHRASE**:
```
dev
```

**PULUMI_ACCESS_TOKEN**:

Pulumi token generation:
```bash
https://app.pulumi.com/account/tokens
```

**OCI_CONFIG**:

OCI configuration file content:
```bash
cat ~/.oci/config
```

**OCI_PRIVATE_KEY**:

OCI private key content:
```bash
cat ~/.oci/pulumi.pem
```

### Trigger

**Automatic**: Every 30 minutes via cron schedule (stops after successful deployment)

**Manual**: Actions → Deploy → Run workflow → Select stack (prod/staging)

### Disable

Workflow disabling:
- Actions → Deploy → Disable workflow
- Or schedule section removal in workflow file
