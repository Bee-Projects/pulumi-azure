# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Lab(pulumi.CustomResource):
    """
    Manages a Dev Test Lab.
    """
    def __init__(__self__, __name__, __opts__=None, location=None, name=None, resource_group_name=None, storage_type=None, tags=None):
        """Create a Lab resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not location:
            raise TypeError('Missing required property location')
        __props__['location'] = location

        __props__['name'] = name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resourceGroupName'] = resource_group_name

        __props__['storageType'] = storage_type

        __props__['tags'] = tags

        __props__['artifacts_storage_account_id'] = None
        __props__['default_premium_storage_account_id'] = None
        __props__['default_storage_account_id'] = None
        __props__['key_vault_id'] = None
        __props__['premium_data_disk_storage_account_id'] = None
        __props__['unique_identifier'] = None

        super(Lab, __self__).__init__(
            'azure:devtest/lab:Lab',
            __name__,
            __props__,
            __opts__)

