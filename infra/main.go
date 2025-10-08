package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Load configuration
		cfg := LoadConfig(config.New(ctx, ""))

		// Create network resources
		network, err := CreateNetwork(ctx, cfg)
		if err != nil {
			return err
		}

		// Create compute resources
		compute, err := CreateCompute(ctx, cfg, network)
		if err != nil {
			return err
		}

		// Export outputs
		ctx.Export("vncId", network.VCN.ID())
		ctx.Export("subnetId", network.Subnet.ID())
		ctx.Export("instanceId", compute.Instance.ID())
		ctx.Export("instancePublicIp", compute.Instance.PublicIp)
		ctx.Export("instancePrivateIp", compute.Instance.PrivateIp)

		return nil
	})
}
