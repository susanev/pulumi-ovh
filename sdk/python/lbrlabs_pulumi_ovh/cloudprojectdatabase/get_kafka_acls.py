# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetKafkaAclsResult',
    'AwaitableGetKafkaAclsResult',
    'get_kafka_acls',
    'get_kafka_acls_output',
]

@pulumi.output_type
class GetKafkaAclsResult:
    """
    A collection of values returned by getKafkaAcls.
    """
    def __init__(__self__, acl_ids=None, cluster_id=None, id=None, service_name=None):
        if acl_ids and not isinstance(acl_ids, list):
            raise TypeError("Expected argument 'acl_ids' to be a list")
        pulumi.set(__self__, "acl_ids", acl_ids)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="aclIds")
    def acl_ids(self) -> Sequence[str]:
        """
        The list of ACLs ids of the kafka cluster associated with the project.
        """
        return pulumi.get(self, "acl_ids")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")


class AwaitableGetKafkaAclsResult(GetKafkaAclsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKafkaAclsResult(
            acl_ids=self.acl_ids,
            cluster_id=self.cluster_id,
            id=self.id,
            service_name=self.service_name)


def get_kafka_acls(cluster_id: Optional[str] = None,
                   service_name: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKafkaAclsResult:
    """
    Use this data source to get the list of ACLs of a kafka cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    acls = ovh.CloudProjectDatabase.get_kafka_acls(service_name="XXX",
        cluster_id="YYY")
    pulumi.export("aclIds", acls.acl_ids)
    ```


    :param str cluster_id: Cluster ID
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProjectDatabase/getKafkaAcls:getKafkaAcls', __args__, opts=opts, typ=GetKafkaAclsResult).value

    return AwaitableGetKafkaAclsResult(
        acl_ids=__ret__.acl_ids,
        cluster_id=__ret__.cluster_id,
        id=__ret__.id,
        service_name=__ret__.service_name)


@_utilities.lift_output_func(get_kafka_acls)
def get_kafka_acls_output(cluster_id: Optional[pulumi.Input[str]] = None,
                          service_name: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKafkaAclsResult]:
    """
    Use this data source to get the list of ACLs of a kafka cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    acls = ovh.CloudProjectDatabase.get_kafka_acls(service_name="XXX",
        cluster_id="YYY")
    pulumi.export("aclIds", acls.acl_ids)
    ```


    :param str cluster_id: Cluster ID
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...
