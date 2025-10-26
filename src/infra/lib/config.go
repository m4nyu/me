package lib

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"

type Config struct {
	VPSHost            string
	VPSUser            string
	CloudflareAPIToken string
	CloudflareZoneID   string
	Domain             string
	GitRepo            string
	GitBranch          string
	ProdContainer      string
	StagContainer      string
}

func Load(c *config.Config) *Config {
	cfg := &Config{
		VPSHost:            c.Get("vpsHost"),
		VPSUser:            c.Get("vpsUser"),
		CloudflareAPIToken: c.Require("cloudflareApiToken"),
		CloudflareZoneID:   c.Require("cloudflareZoneId"),
		Domain:             c.Require("domain"),
		GitRepo:            c.Get("gitRepo"),
		GitBranch:          c.Get("gitBranch"),
		ProdContainer:      c.Get("prodContainer"),
		StagContainer:      c.Get("stagContainer"),
	}

	if cfg.GitRepo == "" {
		cfg.GitRepo = "m4nyu/me"
	}
	if cfg.GitBranch == "" {
		cfg.GitBranch = "main"
	}
	if cfg.ProdContainer == "" {
		cfg.ProdContainer = "prod"
	}
	if cfg.StagContainer == "" {
		cfg.StagContainer = "staging"
	}

	return cfg
}
