package lib

import (
	"github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DNS struct {
	RootRecord *cloudflare.Record
	WWWRecord  *cloudflare.Record
}

func SetupDNS(ctx *pulumi.Context, cfg *Config) (*DNS, error) {
	rootRecord, err := createDNSRecord(ctx, "root", cfg, "@", "Root domain")
	if err != nil {
		return nil, err
	}

	wwwRecord, err := createDNSRecord(ctx, "www", cfg, "www", "WWW subdomain")
	if err != nil {
		return nil, err
	}

	return &DNS{
		RootRecord: rootRecord,
		WWWRecord:  wwwRecord,
	}, nil
}

func createDNSRecord(ctx *pulumi.Context, name string, cfg *Config, recordName, comment string) (*cloudflare.Record, error) {
	return cloudflare.NewRecord(ctx, name, &cloudflare.RecordArgs{
		ZoneId:  pulumi.String(cfg.CloudflareZoneID),
		Name:    pulumi.String(recordName),
		Type:    pulumi.String("A"),
		Content: pulumi.String(cfg.VPSHost),
		Ttl:     pulumi.Int(1),
		Proxied: pulumi.Bool(true),
		Comment: pulumi.String(comment),
	})
}
