# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Credential(pulumi.CustomResource):
    """
    Manages a Automation Credential.
    """
    def __init__(__self__, __name__, __opts__=None, account_name=None, description=None, name=None, password=None, resource_group_name=None, username=None):
        """Create a Credential resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not account_name:
            raise TypeError('Missing required property account_name')
        __props__['account_name'] = account_name

        __props__['description'] = description

        __props__['name'] = name

        if not password:
            raise TypeError('Missing required property password')
        __props__['password'] = password

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resource_group_name'] = resource_group_name

        if not username:
            raise TypeError('Missing required property username')
        __props__['username'] = username

        super(Credential, __self__).__init__(
            'azure:automation/credential:Credential',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

