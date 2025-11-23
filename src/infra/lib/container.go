package lib

import (
	"fmt"

	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Container struct {
	Provider  *docker.Provider
	Image     *docker.RemoteImage
	Container *docker.Container
}

func SetupContainer(ctx *pulumi.Context, cfg *Config) (*Container, error) {
	provider, err := docker.NewProvider(ctx, "docker", &docker.ProviderArgs{
		Host: pulumi.Sprintf("ssh://%s@%s", cfg.VPSUser, cfg.VPSHost),
	})
	if err != nil {
		return nil, err
	}

	containerName := ctx.Stack()

	image, err := docker.NewRemoteImage(ctx, "image", &docker.RemoteImageArgs{
		Name: pulumi.String(fmt.Sprintf("%s:%s", containerName, cfg.GitBranch)),
		Build: docker.RemoteImageBuildArgs{
			Context:    pulumi.Sprintf("https://github.com/%s.git#%s", cfg.GitRepo, cfg.GitBranch),
			Dockerfile: pulumi.String("Dockerfile"),
			Platform:   pulumi.String("linux/amd64"),
		},
		KeepLocally: pulumi.Bool(true),
	}, pulumi.Provider(provider))
	if err != nil {
		return nil, err
	}

	var portConfig docker.ContainerPortArray
	if ctx.Stack() == "prod" {
		portConfig = docker.ContainerPortArray{
			&docker.ContainerPortArgs{
				Internal: pulumi.Int(3000),
				External: pulumi.Int(80),
			},
		}
	} else {
		portConfig = docker.ContainerPortArray{
			&docker.ContainerPortArgs{
				Internal: pulumi.Int(3000),
				External: pulumi.Int(8080),
				Ip:       pulumi.String("127.0.0.1"),
			},
		}
	}

	container, err := docker.NewContainer(ctx, "container", &docker.ContainerArgs{
		Name:  pulumi.String(containerName),
		Image: image.ImageId,
		Ports: portConfig,
		Restart:       pulumi.String("always"),
		MustRun:       pulumi.Bool(true),
		NetworkMode:   pulumi.String("bridge"),
		RemoveVolumes: pulumi.Bool(false),
		ReadOnly:      pulumi.Bool(false),
		Memory:        pulumi.Int(512),
		MemorySwap:    pulumi.Int(1024),
		CpuShares:     pulumi.Int(1024),
		Healthcheck: &docker.ContainerHealthcheckArgs{
			Tests: pulumi.StringArray{
				pulumi.String("CMD"),
				pulumi.String("wget"),
				pulumi.String("--spider"),
				pulumi.String("http://localhost:3000"),
			},
			Interval:    pulumi.String("30s"),
			Timeout:     pulumi.String("10s"),
			StartPeriod: pulumi.String("40s"),
			Retries:     pulumi.Int(3),
		},
		SecurityOpts: pulumi.StringArray{
			pulumi.String("no-new-privileges:true"),
		},
	}, pulumi.Provider(provider), pulumi.ReplaceOnChanges([]string{"image"}))
	if err != nil {
		return nil, err
	}

	return &Container{
		Provider:  provider,
		Image:     image,
		Container: container,
	}, nil
}
