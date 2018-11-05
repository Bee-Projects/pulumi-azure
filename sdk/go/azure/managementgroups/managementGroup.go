// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package managementgroups

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Management Group.
type ManagementGroup struct {
	s *pulumi.ResourceState
}

// NewManagementGroup registers a new resource with the given unique name, arguments, and options.
func NewManagementGroup(ctx *pulumi.Context,
	name string, args *ManagementGroupArgs, opts ...pulumi.ResourceOpt) (*ManagementGroup, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["displayName"] = nil
		inputs["groupId"] = nil
		inputs["parentManagementGroupId"] = nil
		inputs["subscriptionIds"] = nil
	} else {
		inputs["displayName"] = args.DisplayName
		inputs["groupId"] = args.GroupId
		inputs["parentManagementGroupId"] = args.ParentManagementGroupId
		inputs["subscriptionIds"] = args.SubscriptionIds
	}
	s, err := ctx.RegisterResource("azure:managementgroups/managementGroup:ManagementGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ManagementGroup{s: s}, nil
}

// GetManagementGroup gets an existing ManagementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ManagementGroupState, opts ...pulumi.ResourceOpt) (*ManagementGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["displayName"] = state.DisplayName
		inputs["groupId"] = state.GroupId
		inputs["parentManagementGroupId"] = state.ParentManagementGroupId
		inputs["subscriptionIds"] = state.SubscriptionIds
	}
	s, err := ctx.ReadResource("azure:managementgroups/managementGroup:ManagementGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ManagementGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ManagementGroup) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ManagementGroup) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A friendly name for this Management Group. If not specified, this'll be the same as the `group_id`.
func (r *ManagementGroup) DisplayName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["displayName"])
}

// The UUID for this Management Group, which needs to be unique across your tenant - which will be generated if not provided. Changing this forces a new resource to be created.
func (r *ManagementGroup) GroupId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["groupId"])
}

// The ID of the Parent Management Group. Changing this forces a new resource to be created.
func (r *ManagementGroup) ParentManagementGroupId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["parentManagementGroupId"])
}

// A list of Subscription ID's which should be assigned to the Management Group.
func (r *ManagementGroup) SubscriptionIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["subscriptionIds"])
}

// Input properties used for looking up and filtering ManagementGroup resources.
type ManagementGroupState struct {
	// A friendly name for this Management Group. If not specified, this'll be the same as the `group_id`.
	DisplayName interface{}
	// The UUID for this Management Group, which needs to be unique across your tenant - which will be generated if not provided. Changing this forces a new resource to be created.
	GroupId interface{}
	// The ID of the Parent Management Group. Changing this forces a new resource to be created.
	ParentManagementGroupId interface{}
	// A list of Subscription ID's which should be assigned to the Management Group.
	SubscriptionIds interface{}
}

// The set of arguments for constructing a ManagementGroup resource.
type ManagementGroupArgs struct {
	// A friendly name for this Management Group. If not specified, this'll be the same as the `group_id`.
	DisplayName interface{}
	// The UUID for this Management Group, which needs to be unique across your tenant - which will be generated if not provided. Changing this forces a new resource to be created.
	GroupId interface{}
	// The ID of the Parent Management Group. Changing this forces a new resource to be created.
	ParentManagementGroupId interface{}
	// A list of Subscription ID's which should be assigned to the Management Group.
	SubscriptionIds interface{}
}
