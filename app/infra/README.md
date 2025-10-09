# infra

## setup

```bash
cp .env.example .env
oci setup config
ssh-keygen -t ed25519 -f ~/.ssh/oci
```

## configure

```bash
source .env
pulumi login
pulumi stack init dev
pulumi config set oci:region us-phoenix-1
pulumi config set compartmentId <ocid>
pulumi config set tenancyId <ocid>
pulumi config set availabilityDomain <domain>
pulumi config set --secret sshPublicKey "$(cat ~/.ssh/oci.pub)"
```

## deploy

```bash
source .env
pulumi up
```

## outputs

```bash
source .env
pulumi stack output publicIp
```

## access

```bash
ssh -i ~/.ssh/oci ubuntu@$(pulumi stack output publicIp)
```

## destroy

```bash
source .env
pulumi destroy
```
