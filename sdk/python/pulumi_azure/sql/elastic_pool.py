# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class ElasticPool(pulumi.CustomResource):
    """
    Allows you to manage an Azure SQL Elastic Pool.
    """
    def __init__(__self__, __name__, __opts__=None, db_dtu_max=None, db_dtu_min=None, dtu=None, edition=None, location=None, name=None, pool_size=None, resource_group_name=None, server_name=None, tags=None):
        """Create a ElasticPool resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['dbDtuMax'] = db_dtu_max

        __props__['dbDtuMin'] = db_dtu_min

        if not dtu:
            raise TypeError('Missing required property dtu')
        __props__['dtu'] = dtu

        if not edition:
            raise TypeError('Missing required property edition')
        __props__['edition'] = edition

        if not location:
            raise TypeError('Missing required property location')
        __props__['location'] = location

        __props__['name'] = name

        __props__['poolSize'] = pool_size

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resourceGroupName'] = resource_group_name

        if not server_name:
            raise TypeError('Missing required property server_name')
        __props__['serverName'] = server_name

        __props__['tags'] = tags

        __props__['creation_date'] = None

        super(ElasticPool, __self__).__init__(
            'azure:sql/elasticPool:ElasticPool',
            __name__,
            __props__,
            __opts__)

