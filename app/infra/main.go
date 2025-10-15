package main

import (
	"infra/lib"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := lib.Load(config.New(ctx, ""))

		firewall, err := lib.SetupFirewall(ctx, cfg)
		if err != nil {
			return err
		}

		dns, err := lib.SetupDNS(ctx, cfg)
		if err != nil {
			return err
		}

		cdn, err := lib.SetupCDN(ctx, cfg)
		if err != nil {
			return err
		}

		security, err := lib.SetupSecurity(ctx, cfg)
		if err != nil {
			return err
		}

		container, err := lib.SetupContainer(ctx, cfg)
		if err != nil {
			return err
		}

		monitoring, err := lib.SetupMonitoring(ctx, cfg, container.Provider)
		if err != nil {
			return err
		}

		ctx.Export("domain", pulumi.String(cfg.Domain))
		ctx.Export("vpsHost", pulumi.String(cfg.VPSHost))
		ctx.Export("sshTunnel", pulumi.Sprintf("ssh -L 19999:localhost:19999 %s@%s", cfg.VPSUser, cfg.VPSHost))
		ctx.Export("firewallRulesId", firewall.Rules.ID())
		ctx.Export("rootRecordId", dns.RootRecord.ID())
		ctx.Export("wwwRecordId", dns.WWWRecord.ID())
		ctx.Export("cdnSettingsId", cdn.Settings.ID())
		ctx.Export("rateLimitId", security.RateLimiting.ID())
		ctx.Export("containerId", container.Container.ID())
		ctx.Export("imageId", container.Image.ID())
		ctx.Export("netdataId", monitoring.Netdata.ID())

		return nil
	})
}
