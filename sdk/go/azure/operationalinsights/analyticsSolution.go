// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package operationalinsights

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Log Analytics (formally Operational Insights) Solution.
type AnalyticsSolution struct {
	s *pulumi.ResourceState
}

// NewAnalyticsSolution registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsSolution(ctx *pulumi.Context,
	name string, args *AnalyticsSolutionArgs, opts ...pulumi.ResourceOpt) (*AnalyticsSolution, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Plan == nil {
		return nil, errors.New("missing required argument 'Plan'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SolutionName == nil {
		return nil, errors.New("missing required argument 'SolutionName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil || args.WorkspaceResourceId == nil {
		return nil, errors.New("missing required argument 'WorkspaceResourceId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["location"] = nil
		inputs["plan"] = nil
		inputs["resourceGroupName"] = nil
		inputs["solutionName"] = nil
		inputs["workspaceName"] = nil
		inputs["workspaceResourceId"] = nil
	} else {
		inputs["location"] = args.Location
		inputs["plan"] = args.Plan
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["solutionName"] = args.SolutionName
		inputs["workspaceName"] = args.WorkspaceName
		inputs["workspaceResourceId"] = args.WorkspaceResourceId
	}
	s, err := ctx.RegisterResource("azure:operationalinsights/analyticsSolution:AnalyticsSolution", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AnalyticsSolution{s: s}, nil
}

// GetAnalyticsSolution gets an existing AnalyticsSolution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsSolution(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AnalyticsSolutionState, opts ...pulumi.ResourceOpt) (*AnalyticsSolution, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["location"] = state.Location
		inputs["plan"] = state.Plan
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["solutionName"] = state.SolutionName
		inputs["workspaceName"] = state.WorkspaceName
		inputs["workspaceResourceId"] = state.WorkspaceResourceId
	}
	s, err := ctx.ReadResource("azure:operationalinsights/analyticsSolution:AnalyticsSolution", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AnalyticsSolution{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AnalyticsSolution) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AnalyticsSolution) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *AnalyticsSolution) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// A `plan` block as documented below.
func (r *AnalyticsSolution) Plan() *pulumi.Output {
	return r.s.State["plan"]
}

// The name of the resource group in which the Log Analytics solution is created. Changing this forces a new resource to be created. Note: The solution and it's related workspace can only exist in the same resource group.
func (r *AnalyticsSolution) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Specifies the name of the solution to be deployed. See [here for options](https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-add-solutions).Changing this forces a new resource to be created.
func (r *AnalyticsSolution) SolutionName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["solutionName"])
}

func (r *AnalyticsSolution) WorkspaceName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["workspaceName"])
}

// The full resource ID of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
func (r *AnalyticsSolution) WorkspaceResourceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["workspaceResourceId"])
}

// Input properties used for looking up and filtering AnalyticsSolution resources.
type AnalyticsSolutionState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// A `plan` block as documented below.
	Plan interface{}
	// The name of the resource group in which the Log Analytics solution is created. Changing this forces a new resource to be created. Note: The solution and it's related workspace can only exist in the same resource group.
	ResourceGroupName interface{}
	// Specifies the name of the solution to be deployed. See [here for options](https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-add-solutions).Changing this forces a new resource to be created.
	SolutionName interface{}
	WorkspaceName interface{}
	// The full resource ID of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
	WorkspaceResourceId interface{}
}

// The set of arguments for constructing a AnalyticsSolution resource.
type AnalyticsSolutionArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// A `plan` block as documented below.
	Plan interface{}
	// The name of the resource group in which the Log Analytics solution is created. Changing this forces a new resource to be created. Note: The solution and it's related workspace can only exist in the same resource group.
	ResourceGroupName interface{}
	// Specifies the name of the solution to be deployed. See [here for options](https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-add-solutions).Changing this forces a new resource to be created.
	SolutionName interface{}
	WorkspaceName interface{}
	// The full resource ID of the Log Analytics workspace with which the solution will be linked. Changing this forces a new resource to be created.
	WorkspaceResourceId interface{}
}
