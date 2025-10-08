package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type Config struct {
	Compartment string
	Tenancy     string
	Domain      string
	Key         string
	Region      string
}

func load(c *config.Config) *Config {
	return &Config{
		Compartment: c.Require("compartmentId"),
		Tenancy:     c.Require("tenancyId"),
		Domain:      c.Require("availabilityDomain"),
		Key:         c.Require("sshPublicKey"),
		Region:      c.Get("region"),
	}
}
