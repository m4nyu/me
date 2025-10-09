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

```bash
# Install Go 1.24 or higher
go version  # Verify installation

# Install Air for hot-reloading (development)
go install github.com/air-verse/air@latest

# Install dependencies
go mod download
```

## ▶ Run

```bash
# Development server with hot-reload (default)
make dev        # Starts at http://localhost:3000

# Alternative: run without hot-reload
make run        # Build and run

# Build only
make build      # Output: .air/main

# Clean build artifacts
make clean
```

## ▲ Deploy

### Deploy to Oracle Cloud (OCI)

This project includes Pulumi infrastructure for automated deployment to Oracle Cloud Free Tier.

1. **Prerequisites**
   ```bash
   # Install OCI CLI and configure
   oci setup config

   # Generate SSH key for instance access
   ssh-keygen -t ed25519 -f ~/.ssh/oci

   # Install Pulumi
   curl -fsSL https://get.pulumi.com | sh
   ```

2. **Configure Infrastructure**
   ```bash
   cd app/infra
   cp .env.example .env
   source .env

   # Login to Pulumi
   pulumi login

   # Initialize stack
   pulumi stack init dev

   # Set configuration
   pulumi config set oci:region us-phoenix-1
   pulumi config set compartmentId <your-compartment-ocid>
   pulumi config set tenancyId <your-tenancy-ocid>
   pulumi config set availabilityDomain <your-domain>
   pulumi config set --secret sshPublicKey "$(cat ~/.ssh/oci.pub)"
   ```

3. **Deploy**
   ```bash
   source .env
   pulumi up
   ```

4. **Access Your Instance**
   ```bash
   # Get the public IP
   pulumi stack output publicIp

   # SSH into the instance
   ssh -i ~/.ssh/oci ubuntu@$(pulumi stack output publicIp)
   ```

5. **Destroy Infrastructure**
   ```bash
   source .env
   pulumi destroy
   ```

### Deploy with Docker

```bash
# Build image
docker build -t me-app .

# Run container
docker run -p 3000:3000 me-app
```

## 🌍 Features

- **Multi-language Support**: 12 languages (EN, DE, FR, ES, IT, PT, NL, RU, JA, KO, ZH, AR)
- **URL-based i18n**: Language routing with cookie preferences
- **Theme Switching**: Light/Dark/System modes
- **Server-Side Rendering**: Pure Go with gomponents
- **Live Reload**: WebSocket-based development workflow
- **Production Ready**: Docker containerization + OCI cloud-init

## 📚 Documentation

- **Development Guidelines**: See [CLAUDE.md](./CLAUDE.md)
- **Infrastructure Setup**: See [app/infra/README.md](./app/infra/README.md)

## 📄 License

MIT