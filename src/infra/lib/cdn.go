package lib

import (
	"github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CDN struct {
	StaticCacheRule *cloudflare.PageRule
	HTMLCacheRule   *cloudflare.PageRule
}

func SetupCDN(ctx *pulumi.Context, cfg *Config) (*CDN, error) {
	staticCache, err := createCacheRule(ctx, "static-cache", cfg, "/static/*", 1)
	if err != nil {
		return nil, err
	}

	htmlCache, err := createCacheRule(ctx, "html-cache", cfg, "/*", 2)
	if err != nil {
		return nil, err
	}

	return &CDN{
		StaticCacheRule: staticCache,
		HTMLCacheRule:   htmlCache,
	}, nil
}

func createCacheRule(ctx *pulumi.Context, name string, cfg *Config, path string, priority int) (*cloudflare.PageRule, error) {
	return cloudflare.NewPageRule(ctx, name, &cloudflare.PageRuleArgs{
		ZoneId:   pulumi.String(cfg.CloudflareZoneID),
		Target:   pulumi.String(cfg.Domain + path),
		Priority: pulumi.Int(priority),
		Actions: &cloudflare.PageRuleActionsArgs{
			CacheLevel: pulumi.String("cache_everything"),
		},
		Status: pulumi.String("active"),
	})
}
