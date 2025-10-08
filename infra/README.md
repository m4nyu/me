# Infrastructure Deployment with Pulumi

This directory contains Pulumi infrastructure-as-code for deploying the personal website to Oracle Cloud Free Tier.

## Prerequisites

1. **Oracle Cloud Account**: Sign up at https://www.oracle.com/cloud/free/
2. **Pulumi Account**: Sign up at https://app.pulumi.com/signup
3. **OCI CLI**: Install from https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/cliinstall.htm
4. **SSH Key Pair**: Generate with `ssh-keygen -t rsa -b 4096 -f ~/.ssh/oci_rsa`

## Oracle Cloud Free Tier

Always Free resources used:
- **Compute**: VM.Standard.A1.Flex (ARM Ampere) - Up to 4 OCPUs and 24GB RAM
- **VCN**: Virtual Cloud Network with Internet Gateway
- **Storage**: Block volumes (up to 200GB total)
- **Public IP**: 2 reserved public IPs

## Setup Steps

### 1. Configure OCI CLI

```bash
oci setup config
```

You'll need:
- **User OCID**: Found in OCI Console → Profile → User Settings
- **Tenancy OCID**: Found in OCI Console → Profile → Tenancy
- **Region**: e.g., `us-phoenix-1`
- **Compartment OCID**: Found in OCI Console → Identity → Compartments

### 2. Get Your Configuration Values

```bash
# Get tenancy OCID
oci iam tenancy get --tenancy-id <your-tenancy-ocid>

# Get compartment ID (use root compartment or create a new one)
oci iam compartment list --compartment-id <your-tenancy-ocid>

# Get availability domains
oci iam availability-domain list --compartment-id <your-compartment-ocid>
```

### 3. Update Pulumi Configuration

Edit `Pulumi.dev.yaml` and set:

```yaml
config:
  oci:region: us-phoenix-1
  me-infra:compartmentId: ocid1.compartment.oc1..YOUR_COMPARTMENT_ID
  me-infra:tenancyId: ocid1.tenancy.oc1..YOUR_TENANCY_ID
  me-infra:availabilityDomain: YOUR_AD_NAME
  me-infra:sshPublicKey: "YOUR_SSH_PUBLIC_KEY_CONTENT"
```

Or use the Pulumi CLI:

```bash
cd infra
pulumi stack init dev
pulumi config set oci:region us-phoenix-1
pulumi config set compartmentId ocid1.compartment.oc1..YOUR_ID
pulumi config set tenancyId ocid1.tenancy.oc1..YOUR_ID
pulumi config set availabilityDomain "YOUR_AD_NAME"
pulumi config set sshPublicKey "$(cat ~/.ssh/oci_rsa.pub)"
```

### 4. Initialize Pulumi Stack

```bash
cd infra
pulumi login  # Or use pulumi login --local for local state
pulumi stack init dev
```

### 5. Deploy Infrastructure

```bash
pulumi up
```

Review the changes and confirm to deploy.

### 6. Get Outputs

After deployment:

```bash
pulumi stack output instancePublicIp
```

### 7. SSH into Instance

```bash
ssh -i ~/.ssh/oci_rsa ubuntu@$(pulumi stack output instancePublicIp)
```

## Deployment Notes

### Application Deployment

The cloud-init script in `main.go` automatically:
1. Installs Go 1.24.7
2. Installs Git
3. Creates an `appuser` service account
4. Clones your repository (update the URL in main.go!)
5. Builds the Go application
6. Creates a systemd service
7. Installs and configures nginx as reverse proxy
8. Opens necessary firewall ports

### Updating the Application

To update the cloud-init script to use your actual GitHub repository:

1. Edit `infra/main.go`
2. Find the line: `git clone https://github.com/yourusername/yourrepo.git app`
3. Replace with your repository URL
4. Run `pulumi up` to update the infrastructure

### Manual Application Deployment

If you prefer manual deployment:

```bash
# SSH into the instance
ssh -i ~/.ssh/oci_rsa ubuntu@<instance-ip>

# Clone your repository
git clone https://github.com/yourusername/yourrepo.git app
cd app

# Build
go build -o server ./src/cmd/server

# Run (or setup systemd service)
./server
```

## Accessing Your Application

After deployment:
- **HTTP**: http://INSTANCE_PUBLIC_IP
- **Direct**: http://INSTANCE_PUBLIC_IP:3000
- **SSH**: ssh -i ~/.ssh/oci_rsa ubuntu@INSTANCE_PUBLIC_IP

## Updating Infrastructure

Make changes to `main.go` and run:

```bash
pulumi up
```

## Destroying Infrastructure

To tear down all resources:

```bash
pulumi destroy
```

## Cost

All resources used are within Oracle Cloud's Always Free tier, so there's **no cost** for:
- 1-4 OCPUs ARM-based compute (VM.Standard.A1.Flex)
- Up to 24 GB RAM
- 2 public IP addresses
- Outbound data transfer (10 TB/month)
- VCN and related networking

## Troubleshooting

### Check cloud-init logs

```bash
ssh -i ~/.ssh/oci_rsa ubuntu@<instance-ip>
sudo cat /var/log/cloud-init-output.log
```

### Check application logs

```bash
sudo journalctl -u me-app -f
```

### Nginx logs

```bash
sudo tail -f /var/log/nginx/access.log
sudo tail -f /var/log/nginx/error.log
```

## Security Considerations

1. **SSH Key**: Keep your private key secure
2. **Firewall**: Only necessary ports are opened (22, 80, 443, 3000)
3. **Updates**: Regularly update the OS and application
4. **HTTPS**: Consider adding Let's Encrypt SSL certificate for production

## Next Steps

1. Add SSL/TLS with Let's Encrypt
2. Set up a custom domain
3. Configure monitoring and alerts
4. Set up automated backups
5. Implement CI/CD for automatic deployments
