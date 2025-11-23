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
		VPSHost:            c.Require("vpsHost"),
		VPSUser:            c.Require("vpsUser"),
		CloudflareAPIToken: c.Get("cloudflareApiToken"),
		CloudflareZoneID:   c.Get("cloudflareZoneId"),
		Domain:             c.Get("domain"),
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
	if cfg.Domain == "" {
		cfg.Domain = "m4nuel.net"
	}

	return cfg
}
