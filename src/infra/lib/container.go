package lib

import (
	"fmt"

	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Container struct {
	Build *remote.Command
	Run   *remote.Command
}

func SetupContainer(ctx *pulumi.Context, cfg *Config) (*Container, error) {
	conn := remote.ConnectionArgs{
		Host: pulumi.String(cfg.VPSHost),
		User: pulumi.String(cfg.VPSUser),
	}

	containerName := ctx.Stack()
	imageName := fmt.Sprintf("%s:%s", containerName, cfg.GitBranch)
	gitURL := fmt.Sprintf("https://github.com/%s.git#%s", cfg.GitRepo, cfg.GitBranch)

	build, err := remote.NewCommand(ctx, "docker-build", &remote.CommandArgs{
		Connection: conn,
		Create:     pulumi.Sprintf("docker build --platform linux/amd64 -t %s %s 2>&1", imageName, gitURL),
		Delete:     pulumi.Sprintf("docker rmi %s 2>/dev/null || true", imageName),
		Triggers:   pulumi.Array{pulumi.String(imageName)},
	})
	if err != nil {
		return nil, err
	}

	var portFlag string
	if ctx.Stack() == "prod" {
		portFlag = "-p 80:3000"
	} else {
		portFlag = "-p 127.0.0.1:8080:3000"
	}

	run, err := remote.NewCommand(ctx, "docker-run", &remote.CommandArgs{
		Connection: conn,
		Create: pulumi.Sprintf(`
			docker rm -f %s 2>/dev/null || true
			docker run -d \
				--name %s \
				%s \
				--restart always \
				--network bridge \
				--memory 512m \
				--memory-swap 1g \
				--cpu-shares 1024 \
				--health-cmd "wget --spider http://localhost:3000 || exit 1" \
				--health-interval 30s \
				--health-timeout 10s \
				--health-start-period 40s \
				--health-retries 3 \
				--security-opt no-new-privileges:true \
				%s
			docker ps --filter name=%s --format '{{.ID}}'
		`, containerName, containerName, portFlag, imageName, containerName),
		Delete: pulumi.Sprintf("docker rm -f %s 2>/dev/null || true", containerName),
	}, pulumi.DependsOn([]pulumi.Resource{build}))
	if err != nil {
		return nil, err
	}

	return &Container{
		Build: build,
		Run:   run,
	}, nil
}
