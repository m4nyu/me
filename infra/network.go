package main

import (
	"github.com/pulumi/pulumi-oci/sdk/go/oci/core"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NetworkResources holds all networking-related resources
type NetworkResources struct {
	VCN             *core.Vcn
	InternetGateway *core.InternetGateway
	RouteTable      *core.RouteTable
	SecurityList    *core.SecurityList
	Subnet          *core.Subnet
}

// CreateNetwork creates all networking resources
func CreateNetwork(ctx *pulumi.Context, cfg *Config) (*NetworkResources, error) {
	// Create Virtual Cloud Network (VCN)
	vcn, err := core.NewVcn(ctx, "me-vcn", &core.VcnArgs{
		CidrBlock:     pulumi.String("10.0.0.0/16"),
		CompartmentId: pulumi.String(cfg.CompartmentID),
		DisplayName:   pulumi.String("me-vcn"),
		DnsLabel:      pulumi.String("mevcn"),
	})
	if err != nil {
		return nil, err
	}

	// Create Internet Gateway
	ig, err := core.NewInternetGateway(ctx, "me-ig", &core.InternetGatewayArgs{
		CompartmentId: pulumi.String(cfg.CompartmentID),
		VcnId:         vcn.ID(),
		DisplayName:   pulumi.String("me-internet-gateway"),
		Enabled:       pulumi.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	// Create Route Table
	rt, err := core.NewRouteTable(ctx, "me-rt", &core.RouteTableArgs{
		CompartmentId: pulumi.String(cfg.CompartmentID),
		VcnId:         vcn.ID(),
		DisplayName:   pulumi.String("me-route-table"),
		RouteRules: core.RouteTableRouteRuleArray{
			&core.RouteTableRouteRuleArgs{
				NetworkEntityId: ig.ID(),
				Destination:     pulumi.String("0.0.0.0/0"),
				DestinationType: pulumi.String("CIDR_BLOCK"),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// Create Security List
	sl, err := createSecurityList(ctx, cfg, vcn)
	if err != nil {
		return nil, err
	}

	// Create Subnet
	subnet, err := core.NewSubnet(ctx, "me-subnet", &core.SubnetArgs{
		CidrBlock:              pulumi.String("10.0.1.0/24"),
		CompartmentId:          pulumi.String(cfg.CompartmentID),
		VcnId:                  vcn.ID(),
		DisplayName:            pulumi.String("me-subnet"),
		DnsLabel:               pulumi.String("mesubnet"),
		RouteTableId:           rt.ID(),
		SecurityListIds:        pulumi.StringArray{sl.ID()},
		ProhibitPublicIpOnVnic: pulumi.Bool(false),
	})
	if err != nil {
		return nil, err
	}

	return &NetworkResources{
		VCN:             vcn,
		InternetGateway: ig,
		RouteTable:      rt,
		SecurityList:    sl,
		Subnet:          subnet,
	}, nil
}

func createSecurityList(ctx *pulumi.Context, cfg *Config, vcn *core.Vcn) (*core.SecurityList, error) {
	return core.NewSecurityList(ctx, "me-sl", &core.SecurityListArgs{
		CompartmentId: pulumi.String(cfg.CompartmentID),
		VcnId:         vcn.ID(),
		DisplayName:   pulumi.String("me-security-list"),
		EgressSecurityRules: core.SecurityListEgressSecurityRuleArray{
			&core.SecurityListEgressSecurityRuleArgs{
				Protocol:    pulumi.String("all"),
				Destination: pulumi.String("0.0.0.0/0"),
			},
		},
		IngressSecurityRules: core.SecurityListIngressSecurityRuleArray{
			// SSH
			&core.SecurityListIngressSecurityRuleArgs{
				Protocol: pulumi.String("6"), // TCP
				Source:   pulumi.String("0.0.0.0/0"),
				TcpOptions: &core.SecurityListIngressSecurityRuleTcpOptionsArgs{
					Min: pulumi.Int(22),
					Max: pulumi.Int(22),
				},
			},
			// HTTP
			&core.SecurityListIngressSecurityRuleArgs{
				Protocol: pulumi.String("6"), // TCP
				Source:   pulumi.String("0.0.0.0/0"),
				TcpOptions: &core.SecurityListIngressSecurityRuleTcpOptionsArgs{
					Min: pulumi.Int(80),
					Max: pulumi.Int(80),
				},
			},
			// HTTPS
			&core.SecurityListIngressSecurityRuleArgs{
				Protocol: pulumi.String("6"), // TCP
				Source:   pulumi.String("0.0.0.0/0"),
				TcpOptions: &core.SecurityListIngressSecurityRuleTcpOptionsArgs{
					Min: pulumi.Int(443),
					Max: pulumi.Int(443),
				},
			},
			// Custom app port (3000)
			&core.SecurityListIngressSecurityRuleArgs{
				Protocol: pulumi.String("6"), // TCP
				Source:   pulumi.String("0.0.0.0/0"),
				TcpOptions: &core.SecurityListIngressSecurityRuleTcpOptionsArgs{
					Min: pulumi.Int(3000),
					Max: pulumi.Int(3000),
				},
			},
		},
	})
}
