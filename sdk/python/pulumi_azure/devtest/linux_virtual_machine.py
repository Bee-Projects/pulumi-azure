# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class LinuxVirtualMachine(pulumi.CustomResource):
    """
    Manages a Linux Virtual Machine within a Dev Test Lab.
    """
    def __init__(__self__, __name__, __opts__=None, allow_claim=None, disallow_public_ip_address=None, gallery_image_reference=None, inbound_nat_rules=None, lab_name=None, lab_subnet_name=None, lab_virtual_network_id=None, location=None, name=None, notes=None, password=None, resource_group_name=None, size=None, ssh_key=None, storage_type=None, tags=None, username=None):
        """Create a LinuxVirtualMachine resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['allowClaim'] = allow_claim

        __props__['disallowPublicIpAddress'] = disallow_public_ip_address

        if not gallery_image_reference:
            raise TypeError('Missing required property gallery_image_reference')
        __props__['galleryImageReference'] = gallery_image_reference

        __props__['inboundNatRules'] = inbound_nat_rules

        if not lab_name:
            raise TypeError('Missing required property lab_name')
        __props__['labName'] = lab_name

        if not lab_subnet_name:
            raise TypeError('Missing required property lab_subnet_name')
        __props__['labSubnetName'] = lab_subnet_name

        if not lab_virtual_network_id:
            raise TypeError('Missing required property lab_virtual_network_id')
        __props__['labVirtualNetworkId'] = lab_virtual_network_id

        if not location:
            raise TypeError('Missing required property location')
        __props__['location'] = location

        __props__['name'] = name

        __props__['notes'] = notes

        __props__['password'] = password

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resourceGroupName'] = resource_group_name

        if not size:
            raise TypeError('Missing required property size')
        __props__['size'] = size

        __props__['sshKey'] = ssh_key

        if not storage_type:
            raise TypeError('Missing required property storage_type')
        __props__['storageType'] = storage_type

        __props__['tags'] = tags

        if not username:
            raise TypeError('Missing required property username')
        __props__['username'] = username

        __props__['fqdn'] = None
        __props__['unique_identifier'] = None

        super(LinuxVirtualMachine, __self__).__init__(
            'azure:devtest/linuxVirtualMachine:LinuxVirtualMachine',
            __name__,
            __props__,
            __opts__)

