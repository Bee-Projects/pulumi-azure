# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetAppServicePlanResult(object):
    """
    A collection of values returned by getAppServicePlan.
    """
    def __init__(__self__, kind=None, location=None, maximum_number_of_workers=None, properties=None, sku=None, tags=None, id=None):
        if kind and not isinstance(kind, str):
            raise TypeError('Expected argument kind to be a str')
        __self__.kind = kind
        """
        The Operating System type of the App Service Plan
        """
        if location and not isinstance(location, str):
            raise TypeError('Expected argument location to be a str')
        __self__.location = location
        """
        The Azure location where the App Service Plan exists
        """
        if maximum_number_of_workers and not isinstance(maximum_number_of_workers, int):
            raise TypeError('Expected argument maximum_number_of_workers to be a int')
        __self__.maximum_number_of_workers = maximum_number_of_workers
        """
        Maximum number of instances that can be assigned to this App Service plan.
        """
        if properties and not isinstance(properties, list):
            raise TypeError('Expected argument properties to be a list')
        __self__.properties = properties
        """
        A `properties` block as documented below.
        """
        if sku and not isinstance(sku, dict):
            raise TypeError('Expected argument sku to be a dict')
        __self__.sku = sku
        """
        A `sku` block as documented below.
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

def get_app_service_plan(name=None, resource_group_name=None):
    """
    Use this data source to access information about an existing App Service Plan (formerly known as a `Server Farm`).
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __ret__ = pulumi.runtime.invoke('azure:appservice/getAppServicePlan:getAppServicePlan', __args__)

    return GetAppServicePlanResult(
        kind=__ret__.get('kind'),
        location=__ret__.get('location'),
        maximum_number_of_workers=__ret__.get('maximumNumberOfWorkers'),
        properties=__ret__.get('properties'),
        sku=__ret__.get('sku'),
        tags=__ret__.get('tags'),
        id=__ret__.get('id'))
