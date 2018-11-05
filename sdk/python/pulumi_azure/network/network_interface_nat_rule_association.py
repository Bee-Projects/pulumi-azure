# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class NetworkInterfaceNatRuleAssociation(pulumi.CustomResource):
    """
    Manages the association between a Network Interface and a Load Balancer's NAT Rule.
    """
    def __init__(__self__, __name__, __opts__=None, ip_configuration_name=None, nat_rule_id=None, network_interface_id=None):
        """Create a NetworkInterfaceNatRuleAssociation resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not ip_configuration_name:
            raise TypeError('Missing required property ip_configuration_name')
        __props__['ipConfigurationName'] = ip_configuration_name

        if not nat_rule_id:
            raise TypeError('Missing required property nat_rule_id')
        __props__['natRuleId'] = nat_rule_id

        if not network_interface_id:
            raise TypeError('Missing required property network_interface_id')
        __props__['networkInterfaceId'] = network_interface_id

        super(NetworkInterfaceNatRuleAssociation, __self__).__init__(
            'azure:network/networkInterfaceNatRuleAssociation:NetworkInterfaceNatRuleAssociation',
            __name__,
            __props__,
            __opts__)

