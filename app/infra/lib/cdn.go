package lib

import (
	"github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CDN struct {
	Settings        *cloudflare.ZoneSettingsOverride
	StaticCacheRule *cloudflare.PageRule
	HTMLCacheRule   *cloudflare.PageRule
}

func SetupCDN(ctx *pulumi.Context, cfg *Config) (*CDN, error) {
	settings, err := cloudflare.NewZoneSettingsOverride(ctx, "cdn-settings", &cloudflare.ZoneSettingsOverrideArgs{
		ZoneId: pulumi.String(cfg.CloudflareZoneID),
		Settings: &cloudflare.ZoneSettingsOverrideSettingsArgs{
			AlwaysOnline:              pulumi.String("on"),
			BrowserCacheTtl:           pulumi.Int(14400),
			Brotli:                    pulumi.String("on"),
			CacheLevel:                pulumi.String("aggressive"),
			DevelopmentMode:           pulumi.String("off"),
			Http2:                     pulumi.String("on"),
			Http3:                     pulumi.String("on"),
			MinTlsVersion:             pulumi.String("1.2"),
			SecurityLevel:             pulumi.String("high"),
			Ssl:                       pulumi.String("full"),
			Tls13:                     pulumi.String("on"),
			Websockets:                pulumi.String("on"),
			AlwaysUseHttps:            pulumi.String("on"),
			AutomaticHttpsRewrites:    pulumi.String("on"),
			OpportunisticEncryption:   pulumi.String("on"),
			UniversalSsl:              pulumi.String("on"),
			BrowserCheck:              pulumi.String("on"),
			ChallengeTtl:              pulumi.Int(1800),
			EmailObfuscation:          pulumi.String("on"),
			HotlinkProtection:         pulumi.String("on"),
			IpGeolocation:             pulumi.String("on"),
			Ipv6:                      pulumi.String("on"),
			ServerSideExclude:         pulumi.String("on"),
			TrueClientIpHeader:        pulumi.String("on"),
		},
	})
	if err != nil {
		return nil, err
	}

	staticCache, err := createCacheRule(ctx, "static-cache", cfg, "/static/*", 1)
	if err != nil {
		return nil, err
	}

	htmlCache, err := createCacheRule(ctx, "html-cache", cfg, "/*", 2)
	if err != nil {
		return nil, err
	}

	return &CDN{
		Settings:        settings,
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
