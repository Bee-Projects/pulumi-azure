# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class EventHubAuthorizationRule(pulumi.CustomResource):
    """
    Manages a Event Hubs authorization Rule within an Event Hub.
    """
    def __init__(__self__, __name__, __opts__=None, eventhub_name=None, listen=None, location=None, manage=None, name=None, namespace_name=None, resource_group_name=None, send=None):
        """Create a EventHubAuthorizationRule resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not eventhub_name:
            raise TypeError('Missing required property eventhub_name')
        __props__['eventhub_name'] = eventhub_name

        __props__['listen'] = listen

        __props__['location'] = location

        __props__['manage'] = manage

        __props__['name'] = name

        if not namespace_name:
            raise TypeError('Missing required property namespace_name')
        __props__['namespace_name'] = namespace_name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resource_group_name'] = resource_group_name

        __props__['send'] = send

        __props__['primary_connection_string'] = None
        __props__['primary_key'] = None
        __props__['secondary_connection_string'] = None
        __props__['secondary_key'] = None

        super(EventHubAuthorizationRule, __self__).__init__(
            'azure:eventhub/eventHubAuthorizationRule:EventHubAuthorizationRule',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

