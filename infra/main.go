package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := load(config.New(ctx, ""))

		network, err := setup_network(ctx, cfg)
		if err != nil {
			return err
		}

		compute, err := setup_compute(ctx, cfg, network)
		if err != nil {
			return err
		}

		ctx.Export("vcn", network.VCN.ID())
		ctx.Export("subnet", network.Subnet.ID())
		ctx.Export("instance", compute.Instance.ID())
		ctx.Export("publicIp", compute.Instance.PublicIp)
		ctx.Export("privateIp", compute.Instance.PrivateIp)

		return nil
	})
}
