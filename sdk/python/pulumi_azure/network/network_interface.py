# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class NetworkInterface(pulumi.CustomResource):
    """
    Manages a Network Interface located in a Virtual Network, usually attached to a Virtual Machine.
    """
    def __init__(__self__, __name__, __opts__=None, applied_dns_servers=None, dns_servers=None, enable_accelerated_networking=None, enable_ip_forwarding=None, internal_dns_name_label=None, internal_fqdn=None, ip_configurations=None, location=None, mac_address=None, name=None, network_security_group_id=None, resource_group_name=None, tags=None, virtual_machine_id=None):
        """Create a NetworkInterface resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['appliedDnsServers'] = applied_dns_servers

        __props__['dnsServers'] = dns_servers

        __props__['enableAcceleratedNetworking'] = enable_accelerated_networking

        __props__['enableIpForwarding'] = enable_ip_forwarding

        __props__['internalDnsNameLabel'] = internal_dns_name_label

        __props__['internalFqdn'] = internal_fqdn

        if not ip_configurations:
            raise TypeError('Missing required property ip_configurations')
        __props__['ipConfigurations'] = ip_configurations

        if not location:
            raise TypeError('Missing required property location')
        __props__['location'] = location

        __props__['macAddress'] = mac_address

        __props__['name'] = name

        __props__['networkSecurityGroupId'] = network_security_group_id

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resourceGroupName'] = resource_group_name

        __props__['tags'] = tags

        __props__['virtualMachineId'] = virtual_machine_id

        __props__['private_ip_address'] = None
        __props__['private_ip_addresses'] = None

        super(NetworkInterface, __self__).__init__(
            'azure:network/networkInterface:NetworkInterface',
            __name__,
            __props__,
            __opts__)

