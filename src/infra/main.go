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
			ctx.Export("staticCacheId", cdn.StaticCacheRule.ID())
			ctx.Export("htmlCacheId", cdn.HTMLCacheRule.ID())
			ctx.Export("wafRulesetId", security.WAF.ID())
			ctx.Export("rateLimitRulesetId", security.RateLimit.ID())
		}

		container, err := lib.SetupContainer(ctx, cfg)
		if err != nil {
			return err
		}

		ctx.Export("domain", pulumi.String(cfg.Domain))
		ctx.Export("vpsHost", pulumi.String(cfg.VPSHost))
		ctx.Export("buildId", container.Build.ID())
		ctx.Export("runId", container.Run.ID())

		return nil
	})
}
