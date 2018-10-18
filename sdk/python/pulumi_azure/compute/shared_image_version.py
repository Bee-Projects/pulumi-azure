# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class SharedImageVersion(pulumi.CustomResource):
    """
    Manages a Version of a Shared Image within a Shared Image Gallery.
    
    -> **NOTE** Shared Image Galleries are currently in Public Preview. You can find more information, including [how to register for the Public Preview here](https://azure.microsoft.com/en-gb/blog/announcing-the-public-preview-of-shared-image-gallery/).
    """
    def __init__(__self__, __name__, __opts__=None, exclude_from_latest=None, gallery_name=None, image_name=None, location=None, managed_image_id=None, name=None, resource_group_name=None, tags=None, target_regions=None):
        """Create a SharedImageVersion resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if exclude_from_latest and not isinstance(exclude_from_latest, bool):
            raise TypeError('Expected property exclude_from_latest to be a bool')
        __self__.exclude_from_latest = exclude_from_latest
        """
        Should this Image Version be excluded from the `latest` filter? If set to `true` this Image Version won't be returned for the `latest` version. Defaults to `false`.
        """
        __props__['excludeFromLatest'] = exclude_from_latest

        if not gallery_name:
            raise TypeError('Missing required property gallery_name')
        elif not isinstance(gallery_name, basestring):
            raise TypeError('Expected property gallery_name to be a basestring')
        __self__.gallery_name = gallery_name
        """
        The name of the Shared Image Gallery in which the Shared Image exists. Changing this forces a new resource to be created.
        """
        __props__['galleryName'] = gallery_name

        if not image_name:
            raise TypeError('Missing required property image_name')
        elif not isinstance(image_name, basestring):
            raise TypeError('Expected property image_name to be a basestring')
        __self__.image_name = image_name
        """
        The name of the Shared Image within the Shared Image Gallery in which this Version should be created. Changing this forces a new resource to be created.
        """
        __props__['imageName'] = image_name

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        The Azure Region in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        """
        __props__['location'] = location

        if not managed_image_id:
            raise TypeError('Missing required property managed_image_id')
        elif not isinstance(managed_image_id, basestring):
            raise TypeError('Expected property managed_image_id to be a basestring')
        __self__.managed_image_id = managed_image_id
        """
        The ID of the Managed Image which should be used for this Shared Image Version. Changing this forces a new resource to be created.
        """
        __props__['managedImageId'] = managed_image_id

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The version number for this Image Version, such as `1.0.0`. Changing this forces a new resource to be created.
        """
        __props__['name'] = name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the Resource Group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
        """
        __props__['resourceGroupName'] = resource_group_name

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        """
        A collection of tags which should be applied to this resource.
        """
        __props__['tags'] = tags

        if not target_regions:
            raise TypeError('Missing required property target_regions')
        elif not isinstance(target_regions, list):
            raise TypeError('Expected property target_regions to be a list')
        __self__.target_regions = target_regions
        __props__['targetRegions'] = target_regions

        super(SharedImageVersion, __self__).__init__(
            'azure:compute/sharedImageVersion:SharedImageVersion',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'excludeFromLatest' in outs:
            self.exclude_from_latest = outs['excludeFromLatest']
        if 'galleryName' in outs:
            self.gallery_name = outs['galleryName']
        if 'imageName' in outs:
            self.image_name = outs['imageName']
        if 'location' in outs:
            self.location = outs['location']
        if 'managedImageId' in outs:
            self.managed_image_id = outs['managedImageId']
        if 'name' in outs:
            self.name = outs['name']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'tags' in outs:
            self.tags = outs['tags']
        if 'targetRegions' in outs:
            self.target_regions = outs['targetRegions']