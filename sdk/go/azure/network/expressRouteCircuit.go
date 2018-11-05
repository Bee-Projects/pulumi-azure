// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an ExpressRoute circuit.
type ExpressRouteCircuit struct {
	s *pulumi.ResourceState
}

// NewExpressRouteCircuit registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuit(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitArgs, opts ...pulumi.ResourceOpt) (*ExpressRouteCircuit, error) {
	if args == nil || args.BandwidthInMbps == nil {
		return nil, errors.New("missing required argument 'BandwidthInMbps'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.PeeringLocation == nil {
		return nil, errors.New("missing required argument 'PeeringLocation'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceProviderName == nil {
		return nil, errors.New("missing required argument 'ServiceProviderName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allowClassicOperations"] = nil
		inputs["bandwidthInMbps"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["peeringLocation"] = nil
		inputs["resourceGroupName"] = nil
		inputs["serviceProviderName"] = nil
		inputs["sku"] = nil
		inputs["tags"] = nil
	} else {
		inputs["allowClassicOperations"] = args.AllowClassicOperations
		inputs["bandwidthInMbps"] = args.BandwidthInMbps
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["peeringLocation"] = args.PeeringLocation
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["serviceProviderName"] = args.ServiceProviderName
		inputs["sku"] = args.Sku
		inputs["tags"] = args.Tags
	}
	inputs["serviceKey"] = nil
	inputs["serviceProviderProvisioningState"] = nil
	s, err := ctx.RegisterResource("azure:network/expressRouteCircuit:ExpressRouteCircuit", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ExpressRouteCircuit{s: s}, nil
}

// GetExpressRouteCircuit gets an existing ExpressRouteCircuit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuit(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ExpressRouteCircuitState, opts ...pulumi.ResourceOpt) (*ExpressRouteCircuit, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allowClassicOperations"] = state.AllowClassicOperations
		inputs["bandwidthInMbps"] = state.BandwidthInMbps
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["peeringLocation"] = state.PeeringLocation
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["serviceKey"] = state.ServiceKey
		inputs["serviceProviderName"] = state.ServiceProviderName
		inputs["serviceProviderProvisioningState"] = state.ServiceProviderProvisioningState
		inputs["sku"] = state.Sku
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:network/expressRouteCircuit:ExpressRouteCircuit", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ExpressRouteCircuit{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ExpressRouteCircuit) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ExpressRouteCircuit) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Allow the circuit to interact with classic (RDFE) resources. The default value is `false`.
func (r *ExpressRouteCircuit) AllowClassicOperations() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["allowClassicOperations"])
}

// The bandwidth in Mbps of the circuit being created.
func (r *ExpressRouteCircuit) BandwidthInMbps() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["bandwidthInMbps"])
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *ExpressRouteCircuit) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
func (r *ExpressRouteCircuit) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the peering location and **not** the Azure resource location.
func (r *ExpressRouteCircuit) PeeringLocation() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["peeringLocation"])
}

// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
func (r *ExpressRouteCircuit) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The string needed by the service provider to provision the ExpressRoute circuit.
func (r *ExpressRouteCircuit) ServiceKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceKey"])
}

// The name of the ExpressRoute Service Provider.
func (r *ExpressRouteCircuit) ServiceProviderName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceProviderName"])
}

// The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
func (r *ExpressRouteCircuit) ServiceProviderProvisioningState() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceProviderProvisioningState"])
}

// A `sku` block for the ExpressRoute circuit as documented below.
func (r *ExpressRouteCircuit) Sku() *pulumi.Output {
	return r.s.State["sku"]
}

// A mapping of tags to assign to the resource.
func (r *ExpressRouteCircuit) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering ExpressRouteCircuit resources.
type ExpressRouteCircuitState struct {
	// Allow the circuit to interact with classic (RDFE) resources. The default value is `false`.
	AllowClassicOperations interface{}
	// The bandwidth in Mbps of the circuit being created.
	BandwidthInMbps interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the peering location and **not** the Azure resource location.
	PeeringLocation interface{}
	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The string needed by the service provider to provision the ExpressRoute circuit.
	ServiceKey interface{}
	// The name of the ExpressRoute Service Provider.
	ServiceProviderName interface{}
	// The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".
	ServiceProviderProvisioningState interface{}
	// A `sku` block for the ExpressRoute circuit as documented below.
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a ExpressRouteCircuit resource.
type ExpressRouteCircuitArgs struct {
	// Allow the circuit to interact with classic (RDFE) resources. The default value is `false`.
	AllowClassicOperations interface{}
	// The bandwidth in Mbps of the circuit being created.
	BandwidthInMbps interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the ExpressRoute circuit. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the peering location and **not** the Azure resource location.
	PeeringLocation interface{}
	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The name of the ExpressRoute Service Provider.
	ServiceProviderName interface{}
	// A `sku` block for the ExpressRoute circuit as documented below.
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
