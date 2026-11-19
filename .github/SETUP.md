# Deployment setup

`ci.yml` checks, builds `ghcr.io/m4nyu/me`, and deploys it with Pulumi:

| Event | Stack | Reachable at |
|---|---|---|
| Pull request into `main` (drafts too) | `me-app/staging` | `127.0.0.1:8080` on the VPS, via `ssh -N -L 8080:127.0.0.1:8080 debian@158.69.218.225` |
| Push to `main` | `me-app/prod` + `me-edge/prod` | https://m4nuel.net through a Cloudflare Tunnel |

The VPS publishes no web port: prod is reached only through the tunnel.

## Repository secrets

| Secret | What |
|---|---|
| `PULUMI_ACCESS_TOKEN` | Pulumi Cloud token (state and secret encryption) |
| `CLOUDFLARE_API_TOKEN` | Account: Cloudflare Tunnel Edit; Zone (m4nuel.net): DNS Edit, Zone WAF Edit, Zone Read |
| `VPS_SSH_KEY` | Private ed25519 key whose public half is in `~debian/.ssh/authorized_keys` |
| `PULUMI_CONFIG_PASSPHRASE` | Only until the legacy stacks are retired (see `infra/legacy`) |

## Layout

| Path | |
|---|---|
| `infra/app` | the container (`me-<stack>`), network, image pull by digest |
| `infra/edge` | tunnel, DNS, WAF and rate-limit rules, cloudflared connector (prod only) |
| `infra/legacy` | config of the old Go stacks, retired by the first deploy of each stack; delete afterwards |
