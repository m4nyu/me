package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// Config holds all configuration values for the infrastructure
type Config struct {
	CompartmentID      string
	TenancyID          string
	AvailabilityDomain string
	SSHPublicKey       string
	Region             string
}

// LoadConfig reads configuration from Pulumi config
func LoadConfig(cfg *config.Config) *Config {
	return &Config{
		CompartmentID:      cfg.Require("compartmentId"),
		TenancyID:          cfg.Require("tenancyId"),
		AvailabilityDomain: cfg.Require("availabilityDomain"),
		SSHPublicKey:       cfg.Require("sshPublicKey"),
		Region:             cfg.Get("region"),
	}
}
