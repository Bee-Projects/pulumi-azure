# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class ActionCustom(pulumi.CustomResource):
    """
    Manages a Custom Action within a Logic App Workflow
    """
    def __init__(__self__, __name__, __opts__=None, body=None, logic_app_id=None, name=None):
        """Create a ActionCustom resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not body:
            raise TypeError('Missing required property body')
        __props__['body'] = body

        if not logic_app_id:
            raise TypeError('Missing required property logic_app_id')
        __props__['logic_app_id'] = logic_app_id

        __props__['name'] = name

        super(ActionCustom, __self__).__init__(
            'azure:logicapps/actionCustom:ActionCustom',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

