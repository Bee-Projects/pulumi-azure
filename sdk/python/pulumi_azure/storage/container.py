# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Container(pulumi.CustomResource):
    """
    Manage an Azure Storage Container.
    """
    def __init__(__self__, __name__, __opts__=None, container_access_type=None, name=None, resource_group_name=None, storage_account_name=None):
        """Create a Container resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['containerAccessType'] = container_access_type

        __props__['name'] = name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resourceGroupName'] = resource_group_name

        if not storage_account_name:
            raise TypeError('Missing required property storage_account_name')
        __props__['storageAccountName'] = storage_account_name

        __props__['properties'] = None

        super(Container, __self__).__init__(
            'azure:storage/container:Container',
            __name__,
            __props__,
            __opts__)

