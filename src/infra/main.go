package main

import (
	"infra/lib"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := lib.Load(config.New(ctx, ""))

		isProd := ctx.Stack() == "prod"

		if isProd {
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

			ctx.Export("firewallRulesId", firewall.Rules.ID())
			ctx.Export("rootRecordId", dns.RootRecord.ID())
			ctx.Export("wwwRecordId", dns.WWWRecord.ID())
			ctx.Export("cdnSettingsId", cdn.Settings.ID())
			ctx.Export("rateLimitId", security.RateLimiting.ID())
		}

		container, err := lib.SetupContainer(ctx, cfg)
		if err != nil {
			return err
		}

		ctx.Export("domain", pulumi.String(cfg.Domain))
		ctx.Export("vpsHost", pulumi.String(cfg.VPSHost))
		ctx.Export("containerId", container.Container.ID())
		ctx.Export("imageId", container.Image.ID())

		return nil
	})
}
