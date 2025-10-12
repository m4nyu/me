package lib

import (
	"github.com/pulumi/pulumi-oci/sdk/go/oci/core"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Network struct {
	VCN      *core.Vcn
	Gateway  *core.InternetGateway
	Routes   *core.RouteTable
	Security *core.SecurityList
	Subnet   *core.Subnet
}

type Port struct {
	Number int
}

var ports = []Port{
	{22},
	{80},
	{443},
	{3000},
}

func SetupNetwork(ctx *pulumi.Context, cfg *Config) (*Network, error) {
	vcn, err := createVcn(ctx, cfg)
	if err != nil {
		return nil, err
	}

	gateway, err := createGateway(ctx, cfg, vcn)
	if err != nil {
		return nil, err
	}

	routes, err := createRoutes(ctx, cfg, vcn, gateway)
	if err != nil {
		return nil, err
	}

	security, err := newSecurity(ctx, cfg, vcn)
	if err != nil {
		return nil, err
	}

	subnet, err := createSubnet(ctx, cfg, vcn, routes, security)
	if err != nil {
		return nil, err
	}

	return &Network{
		VCN:      vcn,
		Gateway:  gateway,
		Routes:   routes,
		Security: security,
		Subnet:   subnet,
	}, nil
}

func createVcn(ctx *pulumi.Context, cfg *Config) (*core.Vcn, error) {
	return core.NewVcn(ctx, "m4nuel-vcn", &core.VcnArgs{
		CidrBlock:     pulumi.String("10.0.0.0/16"),
		CompartmentId: pulumi.String(cfg.Compartment),
		DisplayName:   pulumi.String("m4nuel"),
		DnsLabel:      pulumi.String("m4nuel"),
	})
}

func createGateway(ctx *pulumi.Context, cfg *Config, vcn *core.Vcn) (*core.InternetGateway, error) {
	return core.NewInternetGateway(ctx, "m4nuel-ig", &core.InternetGatewayArgs{
		CompartmentId: pulumi.String(cfg.Compartment),
		VcnId:         vcn.ID(),
		DisplayName:   pulumi.String("m4nuel-gateway"),
		Enabled:       pulumi.Bool(true),
	})
}

func createRoutes(ctx *pulumi.Context, cfg *Config, vcn *core.Vcn, gateway *core.InternetGateway) (*core.RouteTable, error) {
	return core.NewRouteTable(ctx, "m4nuel-rt", &core.RouteTableArgs{
		CompartmentId: pulumi.String(cfg.Compartment),
		VcnId:         vcn.ID(),
		DisplayName:   pulumi.String("m4nuel-routes"),
		RouteRules: core.RouteTableRouteRuleArray{
			&core.RouteTableRouteRuleArgs{
				NetworkEntityId: gateway.ID(),
				Destination:     pulumi.String("0.0.0.0/0"),
				DestinationType: pulumi.String("CIDR_BLOCK"),
			},
		},
	})
}

func createSubnet(ctx *pulumi.Context, cfg *Config, vcn *core.Vcn, routes *core.RouteTable, security *core.SecurityList) (*core.Subnet, error) {
	return core.NewSubnet(ctx, "m4nuel-subnet", &core.SubnetArgs{
		CidrBlock:              pulumi.String("10.0.1.0/24"),
		CompartmentId:          pulumi.String(cfg.Compartment),
		VcnId:                  vcn.ID(),
		DisplayName:            pulumi.String("m4nuel-subnet"),
		DnsLabel:               pulumi.String("m4nuel"),
		RouteTableId:           routes.ID(),
		SecurityListIds:        pulumi.StringArray{security.ID()},
		ProhibitPublicIpOnVnic: pulumi.Bool(false),
	})
}

func newSecurity(ctx *pulumi.Context, cfg *Config, vcn *core.Vcn) (*core.SecurityList, error) {
	var rules core.SecurityListIngressSecurityRuleArray
	for _, p := range ports {
		rules = append(rules, &core.SecurityListIngressSecurityRuleArgs{
			Protocol: pulumi.String("6"),
			Source:   pulumi.String("0.0.0.0/0"),
			TcpOptions: &core.SecurityListIngressSecurityRuleTcpOptionsArgs{
				Min: pulumi.Int(p.Number),
				Max: pulumi.Int(p.Number),
			},
		})
	}

	return core.NewSecurityList(ctx, "m4nuel-sl", &core.SecurityListArgs{
		CompartmentId: pulumi.String(cfg.Compartment),
		VcnId:         vcn.ID(),
		DisplayName:   pulumi.String("m4nuel-security"),
		EgressSecurityRules: core.SecurityListEgressSecurityRuleArray{
			&core.SecurityListEgressSecurityRuleArgs{
				Protocol:    pulumi.String("all"),
				Destination: pulumi.String("0.0.0.0/0"),
			},
		},
		IngressSecurityRules: rules,
	})
}
