package lib

import (
	"github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Security struct {
	WAF           *cloudflare.RulesetArgs
	RateLimiting  *cloudflare.RateLimit
	SecurityLevel *cloudflare.ZoneSettingsOverride
}

func SetupSecurity(ctx *pulumi.Context, cfg *Config) (*Security, error) {
	rateLimit, err := cloudflare.NewRateLimit(ctx, "rate-limit", &cloudflare.RateLimitArgs{
		ZoneId:      pulumi.String(cfg.CloudflareZoneID),
		Threshold:   pulumi.Int(100),
		Period:      pulumi.Int(60),
		Description: pulumi.String("Rate limit 100 requests per minute"),
		Match: &cloudflare.RateLimitMatchArgs{
			Request: &cloudflare.RateLimitMatchRequestArgs{
				UrlPattern: pulumi.String(cfg.Domain + "/*"),
			},
		},
		Action: &cloudflare.RateLimitActionArgs{
			Mode:    pulumi.String("challenge"),
			Timeout: pulumi.Int(60),
		},
	})
	if err != nil {
		return nil, err
	}

	firewallRules, err := createFirewallRules(ctx, cfg)
	if err != nil {
		return nil, err
	}

	securityHeaders, err := cloudflare.NewPageRule(ctx, "security-headers", &cloudflare.PageRuleArgs{
		ZoneId:   pulumi.String(cfg.CloudflareZoneID),
		Target:   pulumi.String(cfg.Domain + "/*"),
		Priority: pulumi.Int(10),
		Actions: &cloudflare.PageRuleActionsArgs{
			SecurityLevel: pulumi.String("high"),
		},
		Status: pulumi.String("active"),
	})
	if err != nil {
		return nil, err
	}

	_ = firewallRules
	_ = securityHeaders

	return &Security{
		RateLimiting: rateLimit,
	}, nil
}

func createFirewallRules(ctx *pulumi.Context, cfg *Config) (*cloudflare.Filter, error) {
	filter, err := cloudflare.NewFilter(ctx, "block-bad-bots", &cloudflare.FilterArgs{
		ZoneId:      pulumi.String(cfg.CloudflareZoneID),
		Description: pulumi.String("Block known bad bots"),
		Expression:  pulumi.String("(cf.client.bot) and not (cf.verified_bot_category in {\"Search Engine Crawler\" \"Page Preview Service\"})"),
	})
	if err != nil {
		return nil, err
	}

	_, err = cloudflare.NewFirewallRule(ctx, "firewall-rule", &cloudflare.FirewallRuleArgs{
		ZoneId:      pulumi.String(cfg.CloudflareZoneID),
		Description: pulumi.String("Block bad bots"),
		FilterId:    filter.ID(),
		Action:      pulumi.String("block"),
	})
	if err != nil {
		return nil, err
	}

	return filter, nil
}
