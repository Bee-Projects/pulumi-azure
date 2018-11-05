# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Certifiate(pulumi.CustomResource):
    """
    Manages a Key Vault Certificate.
    """
    def __init__(__self__, __name__, __opts__=None, certificate=None, certificate_policy=None, name=None, tags=None, vault_uri=None):
        """Create a Certifiate resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['certificate'] = certificate

        if not certificate_policy:
            raise TypeError('Missing required property certificate_policy')
        __props__['certificatePolicy'] = certificate_policy

        __props__['name'] = name

        __props__['tags'] = tags

        if not vault_uri:
            raise TypeError('Missing required property vault_uri')
        __props__['vaultUri'] = vault_uri

        __props__['certificate_data'] = None
        __props__['secret_id'] = None
        __props__['thumbprint'] = None
        __props__['version'] = None

        super(Certifiate, __self__).__init__(
            'azure:keyvault/certifiate:Certifiate',
            __name__,
            __props__,
            __opts__)

