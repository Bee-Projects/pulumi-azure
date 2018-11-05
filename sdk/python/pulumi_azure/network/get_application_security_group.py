# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetApplicationSecurityGroupResult(object):
    """
    A collection of values returned by getApplicationSecurityGroup.
    """
    def __init__(__self__, location=None, tags=None, id=None):
        if location and not isinstance(location, str):
            raise TypeError('Expected argument location to be a str')
        __self__.location = location
        """
        The supported Azure location where the Application Security Group exists.
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

def get_application_security_group(name=None, resource_group_name=None):
    """
    Use this data source to access information about an existing Application Security Group.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __ret__ = pulumi.runtime.invoke('azure:network/getApplicationSecurityGroup:getApplicationSecurityGroup', __args__)

    return GetApplicationSecurityGroupResult(
        location=__ret__.get('location'),
        tags=__ret__.get('tags'),
        id=__ret__.get('id'))
