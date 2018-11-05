# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Topic(pulumi.CustomResource):
    """
    Manage a ServiceBus Topic.
    
    **Note** Topics can only be created in Namespaces with an SKU of `standard` or higher.
    """
    def __init__(__self__, __name__, __opts__=None, auto_delete_on_idle=None, default_message_ttl=None, duplicate_detection_history_time_window=None, enable_batched_operations=None, enable_express=None, enable_filtering_messages_before_publishing=None, enable_partitioning=None, location=None, max_size_in_megabytes=None, name=None, namespace_name=None, requires_duplicate_detection=None, resource_group_name=None, status=None, support_ordering=None):
        """Create a Topic resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['autoDeleteOnIdle'] = auto_delete_on_idle

        __props__['defaultMessageTtl'] = default_message_ttl

        __props__['duplicateDetectionHistoryTimeWindow'] = duplicate_detection_history_time_window

        __props__['enableBatchedOperations'] = enable_batched_operations

        __props__['enableExpress'] = enable_express

        __props__['enableFilteringMessagesBeforePublishing'] = enable_filtering_messages_before_publishing

        __props__['enablePartitioning'] = enable_partitioning

        __props__['location'] = location

        __props__['maxSizeInMegabytes'] = max_size_in_megabytes

        __props__['name'] = name

        if not namespace_name:
            raise TypeError('Missing required property namespace_name')
        __props__['namespaceName'] = namespace_name

        __props__['requiresDuplicateDetection'] = requires_duplicate_detection

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resourceGroupName'] = resource_group_name

        __props__['status'] = status

        __props__['supportOrdering'] = support_ordering

        super(Topic, __self__).__init__(
            'azure:eventhub/topic:Topic',
            __name__,
            __props__,
            __opts__)

