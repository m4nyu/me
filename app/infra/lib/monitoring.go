package lib

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Monitoring struct {
	Provider *docker.Provider
	Netdata  *docker.Container
}

func SetupMonitoring(ctx *pulumi.Context, cfg *Config, provider *docker.Provider) (*Monitoring, error) {
	netdata, err := docker.NewContainer(ctx, "netdata", &docker.ContainerArgs{
		Name:  pulumi.String("netdata"),
		Image: pulumi.String("netdata/netdata:latest"),
		Ports: docker.ContainerPortArray{
			&docker.ContainerPortArgs{
				Internal: pulumi.Int(19999),
				External: pulumi.Int(19999),
				Ip:       pulumi.String("127.0.0.1"),
			},
		},
		Capabilities: &docker.ContainerCapabilitiesArgs{
			Adds: pulumi.StringArray{
				pulumi.String("SYS_PTRACE"),
				pulumi.String("SYS_ADMIN"),
			},
		},
		SecurityOpts: pulumi.StringArray{
			pulumi.String("apparmor=unconfined"),
		},
		Volumes: docker.ContainerVolumeArray{
			&docker.ContainerVolumeArgs{
				HostPath:      pulumi.String("/proc"),
				ContainerPath: pulumi.String("/host/proc"),
				ReadOnly:      pulumi.Bool(true),
			},
			&docker.ContainerVolumeArgs{
				HostPath:      pulumi.String("/sys"),
				ContainerPath: pulumi.String("/host/sys"),
				ReadOnly:      pulumi.Bool(true),
			},
			&docker.ContainerVolumeArgs{
				HostPath:      pulumi.String("/var/run/docker.sock"),
				ContainerPath: pulumi.String("/var/run/docker.sock"),
				ReadOnly:      pulumi.Bool(true),
			},
			&docker.ContainerVolumeArgs{
				HostPath:      pulumi.String("/etc/os-release"),
				ContainerPath: pulumi.String("/host/etc/os-release"),
				ReadOnly:      pulumi.Bool(true),
			},
		},
		Restart:     pulumi.String("always"),
		NetworkMode: pulumi.String("host"),
	}, pulumi.Provider(provider))
	if err != nil {
		return nil, err
	}

	return &Monitoring{
		Provider: provider,
		Netdata:  netdata,
	}, nil
}
