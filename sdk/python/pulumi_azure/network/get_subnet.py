# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetSubnetResult(object):
    """
    A collection of values returned by getSubnet.
    """
    def __init__(__self__, address_prefix=None, ip_configurations=None, network_security_group_id=None, route_table_id=None, id=None):
        if address_prefix and not isinstance(address_prefix, str):
            raise TypeError('Expected argument address_prefix to be a str')
        __self__.address_prefix = address_prefix
        """
        The address prefix used for the subnet.
        """
        if ip_configurations and not isinstance(ip_configurations, list):
            raise TypeError('Expected argument ip_configurations to be a list')
        __self__.ip_configurations = ip_configurations
        """
        The collection of IP Configurations with IPs within this subnet.
        """
        if network_security_group_id and not isinstance(network_security_group_id, str):
            raise TypeError('Expected argument network_security_group_id to be a str')
        __self__.network_security_group_id = network_security_group_id
        """
        The ID of the Network Security Group associated with the subnet.
        """
        if route_table_id and not isinstance(route_table_id, str):
            raise TypeError('Expected argument route_table_id to be a str')
        __self__.route_table_id = route_table_id
        """
        The ID of the Route Table associated with this subnet.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_subnet(name=None, resource_group_name=None, virtual_network_name=None):
    """
    Use this data source to access information about an existing Subnet within a Virtual Network.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['virtualNetworkName'] = virtual_network_name
    __ret__ = pulumi.runtime.invoke('azure:network/getSubnet:getSubnet', __args__)

    return GetSubnetResult(
        address_prefix=__ret__.get('addressPrefix'),
        ip_configurations=__ret__.get('ipConfigurations'),
        network_security_group_id=__ret__.get('networkSecurityGroupId'),
        route_table_id=__ret__.get('routeTableId'),
        id=__ret__.get('id'))
