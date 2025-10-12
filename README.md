<div align="center">

<pre style="background: transparent;">
██╗     ██╗███╗   ██╗██╗  ██╗███████╗
██║     ██║████╗  ██║██║ ██╔╝██╔════╝
██║     ██║██╔██╗ ██║█████╔╝ ███████╗
██║     ██║██║╚██╗██║██╔═██╗ ╚════██║
███████╗██║██║ ╚████║██║  ██╗███████║
╚══════╝╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝
</pre>
[![GitHub](https://img.shields.io/badge/GitHub-m4nyu-181717?style=flat&logo=github)](https://github.com/m4nyu)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Manuel%20Szedlak-0A66C2?style=flat&logo=linkedin)](https://www.linkedin.com/in/manuel-szedlak)
[![X](https://img.shields.io/badge/X-ManuelSzedlak-1DA1F2?style=flat&logo=x)](https://x.ManuelSzedlak)

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=flat&logo=go&logoColor=white)
![Chi](https://img.shields.io/badge/Chi-v5-00ADD8?style=flat)
![gomponents](https://img.shields.io/badge/gomponents-1.2.0-00ADD8?style=flat)
![Docker](https://img.shields.io/badge/Docker-enabled-2496ED?style=flat&logo=docker&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=flat)

**Personal portfolio website built with Go, Chi router, and gomponents for server-side rendering**

</div>

## ▲ Installation

Go 1.24 version verification or higher:

```bash
go version
```

Air installation for hot-reloading during development:

```bash
go install github.com/air-verse/air@latest
```

Project dependencies download:

```bash
go mod download
```

## ▶ Run

### Development (with hot-reload)

Development server start with automatic reloading on file changes:

```bash
make dev
```

### Build and Run

Application compilation:

```bash
make build
```

Compiled binary execution:

```bash
make run
```

Build artifacts removal:

```bash
make clean
```

## ▲ Deploy

### Deploy to Oracle Cloud (OCI)

This project includes Pulumi infrastructure for automated deployment to Oracle Cloud Free Tier.

1. **Prerequisites**

   OCI CLI configuration with credentials:

   ```bash
   oci setup config
   ```

   SSH key pair generation for instance access:

   ```bash
   ssh-keygen -t ed25519 -f ~/.ssh/oci
   ```

   Pulumi installation for infrastructure management:

   ```bash
   curl -fsSL https://get.pulumi.com | sh
   ```

2. **Configure Infrastructure**

   Infrastructure directory navigation and environment setup:

   ```bash
   cd app/infra
   cp .env.example .env
   source .env
   ```

   Pulumi authentication:

   ```bash
   pulumi login
   ```

   New Pulumi stack creation for development:

   ```bash
   pulumi stack init dev
   ```

   OCI settings configuration (replace placeholders with your actual values):

   ```bash
   pulumi config set oci:region us-phoenix-1
   pulumi config set compartmentId <your-compartment-ocid>
   pulumi config set tenancyId <your-tenancy-ocid>
   pulumi config set availabilityDomain <your-domain>
   pulumi config set --secret sshPublicKey "$(cat ~/.ssh/oci.pub)"
   ```

3. **Deploy**

   Infrastructure preview and deployment:

   ```bash
   pulumi up
   ```

4. **Access Your Instance**

   Public IP retrieval of deployed instance:

   ```bash
   pulumi stack output publicIp
   ```

   Instance SSH connection:

   ```bash
   ssh -i ~/.ssh/oci ubuntu@$(pulumi stack output publicIp)
   ```

5. **Destroy Infrastructure**

   Resource teardown:

   ```bash
   pulumi destroy
   ```

### Deploy with Docker

Docker image build:

```bash
docker build -t me-app .
```

Container execution with port 3000 exposure:

```bash
docker run -p 3000:3000 me-app
```