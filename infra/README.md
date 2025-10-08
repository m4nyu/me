# infra

## setup

```bash
oci setup config
ssh-keygen -t rsa -b 4096 -f ~/.ssh/oci
```

## configure

```bash
pulumi login
pulumi stack init dev
pulumi config set oci:region us-phoenix-1
pulumi config set compartmentId <ocid>
pulumi config set tenancyId <ocid>
pulumi config set availabilityDomain <domain>
pulumi config set sshPublicKey "$(cat ~/.ssh/oci.pub)"
```

## deploy

```bash
pulumi up
```

## outputs

```bash
pulumi stack output publicIp
```

## access

```bash
ssh -i ~/.ssh/oci ubuntu@$(pulumi stack output publicIp)
```

## destroy

```bash
pulumi destroy
```
