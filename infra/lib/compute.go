package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/pulumi/pulumi-oci/sdk/go/oci/core"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ComputeResources holds compute-related resources
type ComputeResources struct {
	Instance *core.Instance
}

// CreateCompute creates the compute instance with cloud-init
func CreateCompute(ctx *pulumi.Context, cfg *Config, network *NetworkResources) (*ComputeResources, error) {
	// Get the latest Ubuntu ARM image
	imageID, err := getUbuntuImage(ctx, cfg)
	if err != nil {
		return nil, err
	}

	// Load cloud-init script
	cloudInitBase64, err := loadCloudInitScript()
	if err != nil {
		return nil, err
	}

	// Create compute instance
	instance, err := core.NewInstance(ctx, "me-instance", &core.InstanceArgs{
		AvailabilityDomain: pulumi.String(cfg.AvailabilityDomain),
		CompartmentId:      pulumi.String(cfg.CompartmentID),
		DisplayName:        pulumi.String("me-server"),
		Shape:              pulumi.String("VM.Standard.A1.Flex"),
		ShapeConfig: &core.InstanceShapeConfigArgs{
			Ocpus:       pulumi.Float64(1), // Free tier allows up to 4 OCPUs
			MemoryInGbs: pulumi.Float64(6), // Free tier allows up to 24 GB
		},
		SourceDetails: &core.InstanceSourceDetailsArgs{
			SourceType: pulumi.String("image"),
			SourceId:   pulumi.String(imageID),
		},
		CreateVnicDetails: &core.InstanceCreateVnicDetailsArgs{
			SubnetId:               network.Subnet.ID(),
			AssignPublicIp:         pulumi.String("true"),
			DisplayName:            pulumi.String("me-vnic"),
			SkipSourceDestCheck:    pulumi.Bool(false),
			AssignPrivateDnsRecord: pulumi.Bool(true),
		},
		Metadata: pulumi.StringMap{
			"ssh_authorized_keys": pulumi.String(cfg.SSHPublicKey),
			"user_data":           pulumi.String(cloudInitBase64),
		},
	})
	if err != nil {
		return nil, err
	}

	return &ComputeResources{
		Instance: instance,
	}, nil
}

func getUbuntuImage(ctx *pulumi.Context, cfg *Config) (string, error) {
	images, err := core.GetImages(ctx, &core.GetImagesArgs{
		CompartmentId:          cfg.CompartmentID,
		OperatingSystem:        pulumi.StringRef("Canonical Ubuntu"),
		OperatingSystemVersion: pulumi.StringRef("22.04"),
		Shape:                  pulumi.StringRef("VM.Standard.A1.Flex"), // ARM-based free tier
		SortBy:                 pulumi.StringRef("TIMECREATED"),
		SortOrder:              pulumi.StringRef("DESC"),
	})
	if err != nil {
		return "", err
	}

	if len(images.Images) == 0 {
		return "", fmt.Errorf("no Ubuntu images found")
	}

	return images.Images[0].Id, nil
}

func loadCloudInitScript() (string, error) {
	cloudInitBytes, err := os.ReadFile("scripts/init.sh")
	if err != nil {
		return "", fmt.Errorf("failed to read init script: %w", err)
	}
	return base64.StdEncoding.EncodeToString(cloudInitBytes), nil
}
