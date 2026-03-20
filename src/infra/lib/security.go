package lib

import (
	"github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Security struct {
	WAF      *cloudflare.Ruleset
	RateLimit *cloudflare.Ruleset
}

func SetupSecurity(ctx *pulumi.Context, cfg *Config) (*Security, error) {
	waf, err := cloudflare.NewRuleset(ctx, "waf-custom-rules", &cloudflare.RulesetArgs{
		ZoneId:      pulumi.String(cfg.CloudflareZoneID),
		Kind:        pulumi.String("zone"),
		Name:        pulumi.String("Custom WAF rules"),
		Description: pulumi.String("Block bad bots and enforce security"),
		Phase:       pulumi.String("http_request_firewall_custom"),
		Rules: cloudflare.RulesetRuleArray{
			&cloudflare.RulesetRuleArgs{
				Action:      pulumi.String("block"),
				Expression:  pulumi.String("(cf.client.bot) and not (cf.verified_bot_category in {\"Search Engine Crawler\" \"Page Preview Service\"})"),
				Description: pulumi.String("Block bad bots"),
				Enabled:     pulumi.Bool(true),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	rateLimit, err := cloudflare.NewRuleset(ctx, "rate-limiting", &cloudflare.RulesetArgs{
		ZoneId:      pulumi.String(cfg.CloudflareZoneID),
		Kind:        pulumi.String("zone"),
		Name:        pulumi.String("Rate limiting rules"),
		Description: pulumi.String("Rate limit 100 requests per minute"),
		Phase:       pulumi.String("http_ratelimit"),
		Rules: cloudflare.RulesetRuleArray{
			&cloudflare.RulesetRuleArgs{
				Action:      pulumi.String("block"),
				Expression:  pulumi.String("true"),
				Description: pulumi.String("Rate limit all requests"),
				Enabled:     pulumi.Bool(true),
				Ratelimit: &cloudflare.RulesetRuleRatelimitArgs{
					Characteristics: pulumi.StringArray{
						pulumi.String("cf.colo.id"),
						pulumi.String("ip.src"),
					},
					Period:            pulumi.Int(10),
					RequestsPerPeriod: pulumi.Int(100),
					MitigationTimeout: pulumi.Int(10),
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &Security{
		WAF:      waf,
		RateLimit: rateLimit,
	}, nil
}
