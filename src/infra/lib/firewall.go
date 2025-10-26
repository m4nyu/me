package lib

import (
	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Firewall struct {
	Rules *remote.Command
}

func SetupFirewall(ctx *pulumi.Context, cfg *Config) (*Firewall, error) {
	conn := remote.ConnectionArgs{
		Host: pulumi.String(cfg.VPSHost),
		User: pulumi.String(cfg.VPSUser),
	}

	rules, err := remote.NewCommand(ctx, "firewall-rules", &remote.CommandArgs{
		Connection: conn,
		Create: pulumi.String(`
			set -e

			sudo apt-get update
			sudo apt-get install -y ufw

			sudo ufw --force reset
			sudo ufw default deny incoming
			sudo ufw default allow outgoing

			sudo ufw allow 22/tcp

			CF_IPV4="173.245.48.0/20 103.21.244.0/22 103.22.200.0/22 103.31.4.0/22 141.101.64.0/18 108.162.192.0/18 190.93.240.0/20 188.114.96.0/20 197.234.240.0/22 198.41.128.0/17 162.158.0.0/15 104.16.0.0/13 104.24.0.0/14 172.64.0.0/13 131.0.72.0/22"
			CF_IPV6="2400:cb00::/32 2606:4700::/32 2803:f800::/32 2405:b500::/32 2405:8100::/32 2a06:98c0::/29 2c0f:f248::/32"

			for ip in $CF_IPV4; do
				sudo ufw allow from $ip to any port 80 proto tcp
				sudo ufw allow from $ip to any port 443 proto tcp
			done

			for ip in $CF_IPV6; do
				sudo ufw allow from $ip to any port 80 proto tcp
				sudo ufw allow from $ip to any port 443 proto tcp
			done

			sudo ufw --force enable
			sudo ufw status verbose
		`),
		Delete: pulumi.String(`
			sudo ufw --force reset
			sudo ufw --force disable
		`),
	})
	if err != nil {
		return nil, err
	}

	return &Firewall{
		Rules: rules,
	}, nil
}
