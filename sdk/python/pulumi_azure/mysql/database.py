# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Database(pulumi.CustomResource):
    """
    Manages a MySQL Database within a MySQL Server
    """
    def __init__(__self__, __name__, __opts__=None, charset=None, collation=None, name=None, resource_group_name=None, server_name=None):
        """Create a Database resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not charset:
            raise TypeError('Missing required property charset')
        __props__['charset'] = charset

        if not collation:
            raise TypeError('Missing required property collation')
        __props__['collation'] = collation

        __props__['name'] = name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resourceGroupName'] = resource_group_name

        if not server_name:
            raise TypeError('Missing required property server_name')
        __props__['serverName'] = server_name

        super(Database, __self__).__init__(
            'azure:mysql/database:Database',
            __name__,
            __props__,
            __opts__)

