// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtest

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Virtual Network within a Dev Test Lab.
type VirtualNetwork struct {
	s *pulumi.ResourceState
}

// NewVirtualNetwork registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetwork(ctx *pulumi.Context,
	name string, args *VirtualNetworkArgs, opts ...pulumi.ResourceOpt) (*VirtualNetwork, error) {
	if args == nil || args.LabName == nil {
		return nil, errors.New("missing required argument 'LabName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["labName"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["subnet"] = nil
		inputs["tags"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["labName"] = args.LabName
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["subnet"] = args.Subnet
		inputs["tags"] = args.Tags
	}
	inputs["uniqueIdentifier"] = nil
	s, err := ctx.RegisterResource("azure:devtest/virtualNetwork:VirtualNetwork", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VirtualNetwork{s: s}, nil
}

// GetVirtualNetwork gets an existing VirtualNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetwork(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VirtualNetworkState, opts ...pulumi.ResourceOpt) (*VirtualNetwork, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["labName"] = state.LabName
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["subnet"] = state.Subnet
		inputs["tags"] = state.Tags
		inputs["uniqueIdentifier"] = state.UniqueIdentifier
	}
	s, err := ctx.ReadResource("azure:devtest/virtualNetwork:VirtualNetwork", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VirtualNetwork{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VirtualNetwork) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VirtualNetwork) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A description for the Virtual Network.
func (r *VirtualNetwork) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
func (r *VirtualNetwork) LabName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["labName"])
}

// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
func (r *VirtualNetwork) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
func (r *VirtualNetwork) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A `subnet` block as defined below.
func (r *VirtualNetwork) Subnet() *pulumi.Output {
	return r.s.State["subnet"]
}

// A mapping of tags to assign to the resource.
func (r *VirtualNetwork) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The unique immutable identifier of the Dev Test Virtual Network.
func (r *VirtualNetwork) UniqueIdentifier() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["uniqueIdentifier"])
}

// Input properties used for looking up and filtering VirtualNetwork resources.
type VirtualNetworkState struct {
	// A description for the Virtual Network.
	Description interface{}
	// Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
	LabName interface{}
	// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `subnet` block as defined below.
	Subnet interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The unique immutable identifier of the Dev Test Virtual Network.
	UniqueIdentifier interface{}
}

// The set of arguments for constructing a VirtualNetwork resource.
type VirtualNetworkArgs struct {
	// A description for the Virtual Network.
	Description interface{}
	// Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
	LabName interface{}
	// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `subnet` block as defined below.
	Subnet interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
