package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/pulumi/pulumi-oci/sdk/go/oci/core"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Compute struct {
	Instance *core.Instance
}

type Image struct {
	OS      string
	Version string
	Shape   string
}

type Shape struct {
	Name   string
	Cpus   float64
	Memory float64
}

var ubuntu = Image{
	OS:      "Canonical Ubuntu",
	Version: "22.04",
	Shape:   "VM.Standard.A1.Flex",
}

var shape = Shape{
	Name:   "VM.Standard.A1.Flex",
	Cpus:   1,
	Memory: 6,
}

func setup_compute(ctx *pulumi.Context, cfg *Config, net *Network) (*Compute, error) {
	image, err := find_image(ctx, cfg)
	if err != nil {
		return nil, err
	}

	script, err := load_script()
	if err != nil {
		return nil, err
	}

	instance, err := core.NewInstance(ctx, "me-instance", &core.InstanceArgs{
		AvailabilityDomain: pulumi.String(cfg.Domain),
		CompartmentId:      pulumi.String(cfg.Compartment),
		DisplayName:        pulumi.String("me-server"),
		Shape:              pulumi.String(shape.Name),
		ShapeConfig: &core.InstanceShapeConfigArgs{
			Ocpus:       pulumi.Float64(shape.Cpus),
			MemoryInGbs: pulumi.Float64(shape.Memory),
		},
		SourceDetails: &core.InstanceSourceDetailsArgs{
			SourceType: pulumi.String("image"),
			SourceId:   pulumi.String(image),
		},
		CreateVnicDetails: &core.InstanceCreateVnicDetailsArgs{
			SubnetId:               net.Subnet.ID(),
			AssignPublicIp:         pulumi.String("true"),
			DisplayName:            pulumi.String("me-vnic"),
			SkipSourceDestCheck:    pulumi.Bool(false),
			AssignPrivateDnsRecord: pulumi.Bool(true),
		},
		Metadata: pulumi.StringMap{
			"ssh_authorized_keys": pulumi.String(cfg.Key),
			"user_data":           pulumi.String(script),
		},
	})
	if err != nil {
		return nil, err
	}

	return &Compute{
		Instance: instance,
	}, nil
}

func find_image(ctx *pulumi.Context, cfg *Config) (string, error) {
	images, err := core.GetImages(ctx, &core.GetImagesArgs{
		CompartmentId:          cfg.Compartment,
		OperatingSystem:        pulumi.StringRef(ubuntu.OS),
		OperatingSystemVersion: pulumi.StringRef(ubuntu.Version),
		Shape:                  pulumi.StringRef(ubuntu.Shape),
		SortBy:                 pulumi.StringRef("TIMECREATED"),
		SortOrder:              pulumi.StringRef("DESC"),
	})
	if err != nil {
		return "", err
	}

	if len(images.Images) == 0 {
		return "", fmt.Errorf("no images found")
	}

	return images.Images[0].Id, nil
}

func load_script() (string, error) {
	bytes, err := os.ReadFile("../scripts/init.sh")
	if err != nil {
		return "", fmt.Errorf("script read failed: %w", err)
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}
