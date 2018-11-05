# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetVaultResult(object):
    """
    A collection of values returned by getVault.
    """
    def __init__(__self__, location=None, sku=None, tags=None, id=None):
        if location and not isinstance(location, str):
            raise TypeError('Expected argument location to be a str')
        __self__.location = location
        """
        The Azure location where the resource resides.
        """
        if sku and not isinstance(sku, str):
            raise TypeError('Expected argument sku to be a str')
        __self__.sku = sku
        """
        The vault's current SKU.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError('Expected argument tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags assigned to the resource.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_vault(name=None, resource_group_name=None):
    """
    Use this data source to access information about an existing Recovery Services Vault.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __ret__ = pulumi.runtime.invoke('azure:recoveryservices/getVault:getVault', __args__)

    return GetVaultResult(
        location=__ret__.get('location'),
        sku=__ret__.get('sku'),
        tags=__ret__.get('tags'),
        id=__ret__.get('id'))
