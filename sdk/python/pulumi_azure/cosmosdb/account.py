# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Account(pulumi.CustomResource):
    """
    Manages a CosmosDB (formally DocumentDB) Account.
    """
    def __init__(__self__, __name__, __opts__=None, capabilities=None, consistency_policy=None, enable_automatic_failover=None, failover_policies=None, geo_locations=None, ip_range_filter=None, is_virtual_network_filter_enabled=None, kind=None, location=None, name=None, offer_type=None, resource_group_name=None, tags=None, virtual_network_rules=None):
        """Create a Account resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['capabilities'] = capabilities

        if not consistency_policy:
            raise TypeError('Missing required property consistency_policy')
        __props__['consistencyPolicy'] = consistency_policy

        __props__['enableAutomaticFailover'] = enable_automatic_failover

        __props__['failoverPolicies'] = failover_policies

        __props__['geoLocations'] = geo_locations

        __props__['ipRangeFilter'] = ip_range_filter

        __props__['isVirtualNetworkFilterEnabled'] = is_virtual_network_filter_enabled

        __props__['kind'] = kind

        if not location:
            raise TypeError('Missing required property location')
        __props__['location'] = location

        __props__['name'] = name

        if not offer_type:
            raise TypeError('Missing required property offer_type')
        __props__['offerType'] = offer_type

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resourceGroupName'] = resource_group_name

        __props__['tags'] = tags

        __props__['virtualNetworkRules'] = virtual_network_rules

        __props__['connection_strings'] = None
        __props__['endpoint'] = None
        __props__['primary_master_key'] = None
        __props__['primary_readonly_master_key'] = None
        __props__['read_endpoints'] = None
        __props__['secondary_master_key'] = None
        __props__['secondary_readonly_master_key'] = None
        __props__['write_endpoints'] = None

        super(Account, __self__).__init__(
            'azure:cosmosdb/account:Account',
            __name__,
            __props__,
            __opts__)

