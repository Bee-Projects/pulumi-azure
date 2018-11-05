// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Hostname Binding within an App Service.
type CustomHostnameBinding struct {
	s *pulumi.ResourceState
}

// NewCustomHostnameBinding registers a new resource with the given unique name, arguments, and options.
func NewCustomHostnameBinding(ctx *pulumi.Context,
	name string, args *CustomHostnameBindingArgs, opts ...pulumi.ResourceOpt) (*CustomHostnameBinding, error) {
	if args == nil || args.AppServiceName == nil {
		return nil, errors.New("missing required argument 'AppServiceName'")
	}
	if args == nil || args.Hostname == nil {
		return nil, errors.New("missing required argument 'Hostname'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["appServiceName"] = nil
		inputs["hostname"] = nil
		inputs["resourceGroupName"] = nil
	} else {
		inputs["appServiceName"] = args.AppServiceName
		inputs["hostname"] = args.Hostname
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	s, err := ctx.RegisterResource("azure:appservice/customHostnameBinding:CustomHostnameBinding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CustomHostnameBinding{s: s}, nil
}

// GetCustomHostnameBinding gets an existing CustomHostnameBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomHostnameBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *CustomHostnameBindingState, opts ...pulumi.ResourceOpt) (*CustomHostnameBinding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appServiceName"] = state.AppServiceName
		inputs["hostname"] = state.Hostname
		inputs["resourceGroupName"] = state.ResourceGroupName
	}
	s, err := ctx.ReadResource("azure:appservice/customHostnameBinding:CustomHostnameBinding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CustomHostnameBinding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *CustomHostnameBinding) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *CustomHostnameBinding) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The name of the App Service in which to add the Custom Hostname Binding. Changing this forces a new resource to be created.
func (r *CustomHostnameBinding) AppServiceName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["appServiceName"])
}

// Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
func (r *CustomHostnameBinding) Hostname() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["hostname"])
}

// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
func (r *CustomHostnameBinding) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Input properties used for looking up and filtering CustomHostnameBinding resources.
type CustomHostnameBindingState struct {
	// The name of the App Service in which to add the Custom Hostname Binding. Changing this forces a new resource to be created.
	AppServiceName interface{}
	// Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
	Hostname interface{}
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
}

// The set of arguments for constructing a CustomHostnameBinding resource.
type CustomHostnameBindingArgs struct {
	// The name of the App Service in which to add the Custom Hostname Binding. Changing this forces a new resource to be created.
	AppServiceName interface{}
	// Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
	Hostname interface{}
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
}
