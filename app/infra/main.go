package main

import (
	"infra/lib"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := lib.Load(config.New(ctx, ""))

		network, err := lib.SetupNetwork(ctx, cfg)
		if err != nil {
			return err
		}

		compute, err := lib.SetupCompute(ctx, cfg, network)
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
