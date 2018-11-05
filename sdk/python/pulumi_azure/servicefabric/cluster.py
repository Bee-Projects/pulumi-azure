# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Cluster(pulumi.CustomResource):
    """
    Manage a Service Fabric Cluster.
    """
    def __init__(__self__, __name__, __opts__=None, add_on_features=None, certificate=None, client_certificate_thumbprints=None, cluster_code_version=None, diagnostics_config=None, fabric_settings=None, location=None, management_endpoint=None, name=None, node_types=None, reliability_level=None, resource_group_name=None, tags=None, upgrade_mode=None, vm_image=None):
        """Create a Cluster resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['addOnFeatures'] = add_on_features

        __props__['certificate'] = certificate

        __props__['clientCertificateThumbprints'] = client_certificate_thumbprints

        __props__['clusterCodeVersion'] = cluster_code_version

        __props__['diagnosticsConfig'] = diagnostics_config

        __props__['fabricSettings'] = fabric_settings

        if not location:
            raise TypeError('Missing required property location')
        __props__['location'] = location

        if not management_endpoint:
            raise TypeError('Missing required property management_endpoint')
        __props__['managementEndpoint'] = management_endpoint

        __props__['name'] = name

        if not node_types:
            raise TypeError('Missing required property node_types')
        __props__['nodeTypes'] = node_types

        if not reliability_level:
            raise TypeError('Missing required property reliability_level')
        __props__['reliabilityLevel'] = reliability_level

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resourceGroupName'] = resource_group_name

        __props__['tags'] = tags

        if not upgrade_mode:
            raise TypeError('Missing required property upgrade_mode')
        __props__['upgradeMode'] = upgrade_mode

        if not vm_image:
            raise TypeError('Missing required property vm_image')
        __props__['vmImage'] = vm_image

        __props__['cluster_endpoint'] = None

        super(Cluster, __self__).__init__(
            'azure:servicefabric/cluster:Cluster',
            __name__,
            __props__,
            __opts__)

