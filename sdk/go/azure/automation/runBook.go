// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Automation Runbook.
type RunBook struct {
	s *pulumi.ResourceState
}

// NewRunBook registers a new resource with the given unique name, arguments, and options.
func NewRunBook(ctx *pulumi.Context,
	name string, args *RunBookArgs, opts ...pulumi.ResourceOpt) (*RunBook, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.LogProgress == nil {
		return nil, errors.New("missing required argument 'LogProgress'")
	}
	if args == nil || args.LogVerbose == nil {
		return nil, errors.New("missing required argument 'LogVerbose'")
	}
	if args == nil || args.PublishContentLink == nil {
		return nil, errors.New("missing required argument 'PublishContentLink'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RunbookType == nil {
		return nil, errors.New("missing required argument 'RunbookType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accountName"] = nil
		inputs["content"] = nil
		inputs["description"] = nil
		inputs["location"] = nil
		inputs["logProgress"] = nil
		inputs["logVerbose"] = nil
		inputs["name"] = nil
		inputs["publishContentLink"] = nil
		inputs["resourceGroupName"] = nil
		inputs["runbookType"] = nil
		inputs["tags"] = nil
	} else {
		inputs["accountName"] = args.AccountName
		inputs["content"] = args.Content
		inputs["description"] = args.Description
		inputs["location"] = args.Location
		inputs["logProgress"] = args.LogProgress
		inputs["logVerbose"] = args.LogVerbose
		inputs["name"] = args.Name
		inputs["publishContentLink"] = args.PublishContentLink
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["runbookType"] = args.RunbookType
		inputs["tags"] = args.Tags
	}
	s, err := ctx.RegisterResource("azure:automation/runBook:RunBook", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RunBook{s: s}, nil
}

// GetRunBook gets an existing RunBook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRunBook(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RunBookState, opts ...pulumi.ResourceOpt) (*RunBook, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accountName"] = state.AccountName
		inputs["content"] = state.Content
		inputs["description"] = state.Description
		inputs["location"] = state.Location
		inputs["logProgress"] = state.LogProgress
		inputs["logVerbose"] = state.LogVerbose
		inputs["name"] = state.Name
		inputs["publishContentLink"] = state.PublishContentLink
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["runbookType"] = state.RunbookType
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:automation/runBook:RunBook", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RunBook{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RunBook) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RunBook) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
func (r *RunBook) AccountName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accountName"])
}

// The desired content of the runbook.
func (r *RunBook) Content() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["content"])
}

// A description for this credential.
func (r *RunBook) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *RunBook) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Progress log option.
func (r *RunBook) LogProgress() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["logProgress"])
}

// Verbose log option.
func (r *RunBook) LogVerbose() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["logVerbose"])
}

// Specifies the name of the Runbook. Changing this forces a new resource to be created.
func (r *RunBook) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The published runbook content link.
func (r *RunBook) PublishContentLink() *pulumi.Output {
	return r.s.State["publishContentLink"]
}

// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
func (r *RunBook) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
func (r *RunBook) RunbookType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["runbookType"])
}

func (r *RunBook) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering RunBook resources.
type RunBookState struct {
	// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
	AccountName interface{}
	// The desired content of the runbook.
	Content interface{}
	// A description for this credential.
	Description interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Progress log option.
	LogProgress interface{}
	// Verbose log option.
	LogVerbose interface{}
	// Specifies the name of the Runbook. Changing this forces a new resource to be created.
	Name interface{}
	// The published runbook content link.
	PublishContentLink interface{}
	// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
	RunbookType interface{}
	Tags interface{}
}

// The set of arguments for constructing a RunBook resource.
type RunBookArgs struct {
	// The name of the automation account in which the Runbook is created. Changing this forces a new resource to be created.
	AccountName interface{}
	// The desired content of the runbook.
	Content interface{}
	// A description for this credential.
	Description interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Progress log option.
	LogProgress interface{}
	// Verbose log option.
	LogVerbose interface{}
	// Specifies the name of the Runbook. Changing this forces a new resource to be created.
	Name interface{}
	// The published runbook content link.
	PublishContentLink interface{}
	// The name of the resource group in which the Runbook is created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The type of the runbook - can be either `Graph`, `GraphPowerShell`, `GraphPowerShellWorkflow`, `PowerShellWorkflow`, `PowerShell` or `Script`.
	RunbookType interface{}
	Tags interface{}
}
