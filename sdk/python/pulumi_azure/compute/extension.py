# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Extension(pulumi.CustomResource):
    """
    Manages a Virtual Machine Extension to provide post deployment configuration
    and run automated tasks.
    
    ~> **NOTE:** Custom Script Extensions for Linux & Windows require that the `commandToExecute` returns a `0` exit code to be classified as successfully deployed. You can achieve this by appending `exit 0` to the end of your `commandToExecute`.
    
    -> **NOTE:** Custom Script Extensions require that the Azure Virtual Machine Guest Agent is running on the Virtual Machine.
    """
    def __init__(__self__, __name__, __opts__=None, auto_upgrade_minor_version=None, location=None, name=None, protected_settings=None, publisher=None, resource_group_name=None, settings=None, tags=None, type=None, type_handler_version=None, virtual_machine_name=None):
        """Create a Extension resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['autoUpgradeMinorVersion'] = auto_upgrade_minor_version

        if not location:
            raise TypeError('Missing required property location')
        __props__['location'] = location

        __props__['name'] = name

        __props__['protectedSettings'] = protected_settings

        if not publisher:
            raise TypeError('Missing required property publisher')
        __props__['publisher'] = publisher

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resourceGroupName'] = resource_group_name

        __props__['settings'] = settings

        __props__['tags'] = tags

        if not type:
            raise TypeError('Missing required property type')
        __props__['type'] = type

        if not type_handler_version:
            raise TypeError('Missing required property type_handler_version')
        __props__['typeHandlerVersion'] = type_handler_version

        if not virtual_machine_name:
            raise TypeError('Missing required property virtual_machine_name')
        __props__['virtualMachineName'] = virtual_machine_name

        super(Extension, __self__).__init__(
            'azure:compute/extension:Extension',
            __name__,
            __props__,
            __opts__)

