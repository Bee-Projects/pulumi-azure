# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Share(pulumi.CustomResource):
    """
    Manage an Azure Storage File Share.
    """
    def __init__(__self__, __name__, __opts__=None, name=None, quota=None, resource_group_name=None, storage_account_name=None):
        """Create a Share resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['name'] = name

        __props__['quota'] = quota

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resourceGroupName'] = resource_group_name

        if not storage_account_name:
            raise TypeError('Missing required property storage_account_name')
        __props__['storageAccountName'] = storage_account_name

        __props__['url'] = None

        super(Share, __self__).__init__(
            'azure:storage/share:Share',
            __name__,
            __props__,
            __opts__)

