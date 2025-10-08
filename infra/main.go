package main

import (
	"encoding/base64"
	"fmt"

	"github.com/pulumi/pulumi-oci/sdk/go/oci/core"
	"github.com/pulumi/pulumi-oci/sdk/go/oci/identity"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")

		// Get configuration values
		compartmentId := cfg.Require("compartmentId")
		sshPublicKey := cfg.Require("sshPublicKey")
		availabilityDomain := cfg.Require("availabilityDomain")

		// Get the tenancy OCID for getting available shapes
		tenancyId := cfg.Require("tenancyId")

		// Get availability domains
		ads, err := identity.GetAvailabilityDomains(ctx, &identity.GetAvailabilityDomainsArgs{
			CompartmentId: compartmentId,
		})
		if err != nil {
			return err
		}

		adName := ads.AvailabilityDomains[0].Name

		// Create Virtual Cloud Network (VCN)
		vcn, err := core.NewVcn(ctx, "me-vcn", &core.VcnArgs{
			CidrBlock:     pulumi.String("10.0.0.0/16"),
			CompartmentId: pulumi.String(compartmentId),
			DisplayName:   pulumi.String("me-vcn"),
			DnsLabel:      pulumi.String("mevcn"),
		})
		if err != nil {
			return err
		}

		// Create Internet Gateway
		ig, err := core.NewInternetGateway(ctx, "me-ig", &core.InternetGatewayArgs{
			CompartmentId: pulumi.String(compartmentId),
			VcnId:         vcn.ID(),
			DisplayName:   pulumi.String("me-internet-gateway"),
			Enabled:       pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		// Create Route Table
		rt, err := core.NewRouteTable(ctx, "me-rt", &core.RouteTableArgs{
			CompartmentId: pulumi.String(compartmentId),
			VcnId:         vcn.ID(),
			DisplayName:   pulumi.String("me-route-table"),
			RouteRules: core.RouteTableRouteRuleArray{
				&core.RouteTableRouteRuleArgs{
					NetworkEntityId: ig.ID(),
					Destination:     pulumi.String("0.0.0.0/0"),
					DestinationType: pulumi.String("CIDR_BLOCK"),
				},
			},
		})
		if err != nil {
			return err
		}

		// Create Security List
		sl, err := core.NewSecurityList(ctx, "me-sl", &core.SecurityListArgs{
			CompartmentId: pulumi.String(compartmentId),
			VcnId:         vcn.ID(),
			DisplayName:   pulumi.String("me-security-list"),
			EgressSecurityRules: core.SecurityListEgressSecurityRuleArray{
				&core.SecurityListEgressSecurityRuleArgs{
					Protocol:    pulumi.String("all"),
					Destination: pulumi.String("0.0.0.0/0"),
				},
			},
			IngressSecurityRules: core.SecurityListIngressSecurityRuleArray{
				// SSH
				&core.SecurityListIngressSecurityRuleArgs{
					Protocol: pulumi.String("6"), // TCP
					Source:   pulumi.String("0.0.0.0/0"),
					TcpOptions: &core.SecurityListIngressSecurityRuleTcpOptionsArgs{
						Min: pulumi.Int(22),
						Max: pulumi.Int(22),
					},
				},
				// HTTP
				&core.SecurityListIngressSecurityRuleArgs{
					Protocol: pulumi.String("6"), // TCP
					Source:   pulumi.String("0.0.0.0/0"),
					TcpOptions: &core.SecurityListIngressSecurityRuleTcpOptionsArgs{
						Min: pulumi.Int(80),
						Max: pulumi.Int(80),
					},
				},
				// HTTPS
				&core.SecurityListIngressSecurityRuleArgs{
					Protocol: pulumi.String("6"), // TCP
					Source:   pulumi.String("0.0.0.0/0"),
					TcpOptions: &core.SecurityListIngressSecurityRuleTcpOptionsArgs{
						Min: pulumi.Int(443),
						Max: pulumi.Int(443),
					},
				},
				// Custom app port (3000)
				&core.SecurityListIngressSecurityRuleArgs{
					Protocol: pulumi.String("6"), // TCP
					Source:   pulumi.String("0.0.0.0/0"),
					TcpOptions: &core.SecurityListIngressSecurityRuleTcpOptionsArgs{
						Min: pulumi.Int(3000),
						Max: pulumi.Int(3000),
					},
				},
			},
		})
		if err != nil {
			return err
		}

		// Create Subnet
		subnet, err := core.NewSubnet(ctx, "me-subnet", &core.SubnetArgs{
			CidrBlock:        pulumi.String("10.0.1.0/24"),
			CompartmentId:    pulumi.String(compartmentId),
			VcnId:            vcn.ID(),
			DisplayName:      pulumi.String("me-subnet"),
			DnsLabel:         pulumi.String("mesubnet"),
			RouteTableId:     rt.ID(),
			SecurityListIds:  pulumi.StringArray{sl.ID()},
			ProhibitPublicIpOnVnic: pulumi.Bool(false),
		})
		if err != nil {
			return err
		}

		// Get the latest Ubuntu image
		images, err := core.GetImages(ctx, &core.GetImagesArgs{
			CompartmentId:          compartmentId,
			OperatingSystem:        pulumi.StringRef("Canonical Ubuntu"),
			OperatingSystemVersion: pulumi.StringRef("22.04"),
			Shape:                  pulumi.StringRef("VM.Standard.A1.Flex"), // ARM-based free tier
			SortBy:                 pulumi.StringRef("TIMECREATED"),
			SortOrder:              pulumi.StringRef("DESC"),
		})
		if err != nil {
			return err
		}

		if len(images.Images) == 0 {
			return fmt.Errorf("no Ubuntu images found")
		}

		imageId := images.Images[0].Id

		// Cloud-init script to setup the application
		cloudInit := `#!/bin/bash
set -e

# Update system
apt-get update
apt-get upgrade -y

# Install Go
wget https://go.dev/dl/go1.24.7.linux-arm64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.7.linux-arm64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
export PATH=$PATH:/usr/local/go/bin

# Install Git
apt-get install -y git

# Create app user
useradd -m -s /bin/bash appuser

# Clone and setup application
su - appuser << 'EOF'
cd ~
git clone https://github.com/yourusername/yourrepo.git app
cd app
/usr/local/go/bin/go mod download
/usr/local/go/bin/go build -o /home/appuser/app/server ./src/cmd/server
EOF

# Create systemd service
cat > /etc/systemd/system/me-app.service << 'EOF'
[Unit]
Description=Personal Website
After=network.target

[Service]
Type=simple
User=appuser
WorkingDirectory=/home/appuser/app
ExecStart=/home/appuser/app/server
Restart=always
RestartSec=10
Environment="PORT=3000"

[Install]
WantedBy=multi-user.target
EOF

# Enable and start service
systemctl daemon-reload
systemctl enable me-app
systemctl start me-app

# Setup firewall
ufw allow 22/tcp
ufw allow 80/tcp
ufw allow 443/tcp
ufw allow 3000/tcp
ufw --force enable

# Install and setup nginx as reverse proxy
apt-get install -y nginx
cat > /etc/nginx/sites-available/default << 'EOF'
server {
    listen 80;
    server_name _;

    location / {
        proxy_pass http://localhost:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
EOF

systemctl restart nginx
`

		cloudInitBase64 := base64.StdEncoding.EncodeToString([]byte(cloudInit))

		// Create Compute Instance (Free Tier - ARM-based Ampere A1)
		instance, err := core.NewInstance(ctx, "me-instance", &core.InstanceArgs{
			AvailabilityDomain: pulumi.String(availabilityDomain),
			CompartmentId:      pulumi.String(compartmentId),
			DisplayName:        pulumi.String("me-server"),
			Shape:              pulumi.String("VM.Standard.A1.Flex"),
			ShapeConfig: &core.InstanceShapeConfigArgs{
				Ocpus:       pulumi.Float64(1), // Free tier allows up to 4 OCPUs
				MemoryInGbs: pulumi.Float64(6), // Free tier allows up to 24 GB
			},
			SourceDetails: &core.InstanceSourceDetailsArgs{
				SourceType: pulumi.String("image"),
				SourceId:   pulumi.String(imageId),
			},
			CreateVnicDetails: &core.InstanceCreateVnicDetailsArgs{
				SubnetId:             subnet.ID(),
				AssignPublicIp:       pulumi.String("true"),
				DisplayName:          pulumi.String("me-vnic"),
				SkipSourceDestCheck:  pulumi.Bool(false),
				AssignPrivateDnsRecord: pulumi.Bool(true),
			},
			Metadata: pulumi.StringMap{
				"ssh_authorized_keys": pulumi.String(sshPublicKey),
				"user_data":           pulumi.String(cloudInitBase64),
			},
		})
		if err != nil {
			return err
		}

		// Get the public IP
		vnicAttachments, err := core.GetVnicAttachments(ctx, &core.GetVnicAttachmentsArgs{
			CompartmentId: compartmentId,
			InstanceId:    pulumi.StringRef(instance.ID().ToStringOutput().ApplyT(func(id string) string { return id }).(pulumi.StringOutput).Unwrap()),
		})
		if err != nil {
			return err
		}

		// Export outputs
		ctx.Export("vncId", vcn.ID())
		ctx.Export("subnetId", subnet.ID())
		ctx.Export("instanceId", instance.ID())
		ctx.Export("instancePublicIp", instance.PublicIp)
		ctx.Export("instancePrivateIp", instance.PrivateIp)

		return nil
	})
}
