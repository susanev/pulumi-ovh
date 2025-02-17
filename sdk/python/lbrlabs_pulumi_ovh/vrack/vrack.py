# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['VrackArgs', 'Vrack']

@pulumi.input_type
class VrackArgs:
    def __init__(__self__, *,
                 ovh_subsidiary: pulumi.Input[str],
                 plan: pulumi.Input['VrackPlanArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 payment_mean: Optional[pulumi.Input[str]] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]]] = None):
        """
        The set of arguments for constructing a Vrack resource.
        :param pulumi.Input[str] ovh_subsidiary: OVHcloud Subsidiary
        :param pulumi.Input['VrackPlanArgs'] plan: Product Plan to order
        :param pulumi.Input[str] description: yourvrackdescription
        :param pulumi.Input[str] name: yourvrackname
        :param pulumi.Input[str] payment_mean: Ovh payment mode
        :param pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]] plan_options: Product Plan to order
        """
        pulumi.set(__self__, "ovh_subsidiary", ovh_subsidiary)
        pulumi.set(__self__, "plan", plan)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if payment_mean is not None:
            warnings.warn("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""", DeprecationWarning)
            pulumi.log.warn("""payment_mean is deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
        if payment_mean is not None:
            pulumi.set(__self__, "payment_mean", payment_mean)
        if plan_options is not None:
            pulumi.set(__self__, "plan_options", plan_options)

    @property
    @pulumi.getter(name="ovhSubsidiary")
    def ovh_subsidiary(self) -> pulumi.Input[str]:
        """
        OVHcloud Subsidiary
        """
        return pulumi.get(self, "ovh_subsidiary")

    @ovh_subsidiary.setter
    def ovh_subsidiary(self, value: pulumi.Input[str]):
        pulumi.set(self, "ovh_subsidiary", value)

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Input['VrackPlanArgs']:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: pulumi.Input['VrackPlanArgs']):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        yourvrackdescription
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        yourvrackname
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="paymentMean")
    def payment_mean(self) -> Optional[pulumi.Input[str]]:
        """
        Ovh payment mode
        """
        return pulumi.get(self, "payment_mean")

    @payment_mean.setter
    def payment_mean(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payment_mean", value)

    @property
    @pulumi.getter(name="planOptions")
    def plan_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]]]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan_options")

    @plan_options.setter
    def plan_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]]]):
        pulumi.set(self, "plan_options", value)


@pulumi.input_type
class _VrackState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 orders: Optional[pulumi.Input[Sequence[pulumi.Input['VrackOrderArgs']]]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[str]] = None,
                 payment_mean: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['VrackPlanArgs']] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Vrack resources.
        :param pulumi.Input[str] description: yourvrackdescription
        :param pulumi.Input[str] name: yourvrackname
        :param pulumi.Input[Sequence[pulumi.Input['VrackOrderArgs']]] orders: Details about an Order
        :param pulumi.Input[str] ovh_subsidiary: OVHcloud Subsidiary
        :param pulumi.Input[str] payment_mean: Ovh payment mode
        :param pulumi.Input['VrackPlanArgs'] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]] plan_options: Product Plan to order
        :param pulumi.Input[str] service_name: The internal name of your vrack
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if orders is not None:
            pulumi.set(__self__, "orders", orders)
        if ovh_subsidiary is not None:
            pulumi.set(__self__, "ovh_subsidiary", ovh_subsidiary)
        if payment_mean is not None:
            warnings.warn("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""", DeprecationWarning)
            pulumi.log.warn("""payment_mean is deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
        if payment_mean is not None:
            pulumi.set(__self__, "payment_mean", payment_mean)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if plan_options is not None:
            pulumi.set(__self__, "plan_options", plan_options)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        yourvrackdescription
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        yourvrackname
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def orders(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VrackOrderArgs']]]]:
        """
        Details about an Order
        """
        return pulumi.get(self, "orders")

    @orders.setter
    def orders(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VrackOrderArgs']]]]):
        pulumi.set(self, "orders", value)

    @property
    @pulumi.getter(name="ovhSubsidiary")
    def ovh_subsidiary(self) -> Optional[pulumi.Input[str]]:
        """
        OVHcloud Subsidiary
        """
        return pulumi.get(self, "ovh_subsidiary")

    @ovh_subsidiary.setter
    def ovh_subsidiary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ovh_subsidiary", value)

    @property
    @pulumi.getter(name="paymentMean")
    def payment_mean(self) -> Optional[pulumi.Input[str]]:
        """
        Ovh payment mode
        """
        return pulumi.get(self, "payment_mean")

    @payment_mean.setter
    def payment_mean(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payment_mean", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input['VrackPlanArgs']]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input['VrackPlanArgs']]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter(name="planOptions")
    def plan_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]]]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan_options")

    @plan_options.setter
    def plan_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]]]):
        pulumi.set(self, "plan_options", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The internal name of your vrack
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


class Vrack(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[str]] = None,
                 payment_mean: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input[pulumi.InputType['VrackPlanArgs']]] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VrackPlanOptionArgs']]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh
        import pulumi_ovh as ovh

        mycart = ovh.Order.get_cart(ovh_subsidiary="fr",
            description="my vrack order cart")
        vrack_cart_product_plan = ovh.Order.get_cart_product_plan(cart_id=mycart.id,
            price_capacity="renew",
            product="vrack",
            plan_code="vrack")
        vrack_vrack = ovh.vrack.Vrack("vrackVrack",
            ovh_subsidiary=mycart.ovh_subsidiary,
            description="my vrack",
            plan=ovh.vrack.VrackPlanArgs(
                duration=vrack_cart_product_plan.selected_prices[0].duration,
                plan_code=vrack_cart_product_plan.plan_code,
                pricing_mode=vrack_cart_product_plan.selected_prices[0].pricing_mode,
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: yourvrackdescription
        :param pulumi.Input[str] name: yourvrackname
        :param pulumi.Input[str] ovh_subsidiary: OVHcloud Subsidiary
        :param pulumi.Input[str] payment_mean: Ovh payment mode
        :param pulumi.Input[pulumi.InputType['VrackPlanArgs']] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VrackPlanOptionArgs']]]] plan_options: Product Plan to order
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VrackArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh
        import pulumi_ovh as ovh

        mycart = ovh.Order.get_cart(ovh_subsidiary="fr",
            description="my vrack order cart")
        vrack_cart_product_plan = ovh.Order.get_cart_product_plan(cart_id=mycart.id,
            price_capacity="renew",
            product="vrack",
            plan_code="vrack")
        vrack_vrack = ovh.vrack.Vrack("vrackVrack",
            ovh_subsidiary=mycart.ovh_subsidiary,
            description="my vrack",
            plan=ovh.vrack.VrackPlanArgs(
                duration=vrack_cart_product_plan.selected_prices[0].duration,
                plan_code=vrack_cart_product_plan.plan_code,
                pricing_mode=vrack_cart_product_plan.selected_prices[0].pricing_mode,
            ))
        ```

        :param str resource_name: The name of the resource.
        :param VrackArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VrackArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[str]] = None,
                 payment_mean: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input[pulumi.InputType['VrackPlanArgs']]] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VrackPlanOptionArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VrackArgs.__new__(VrackArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if ovh_subsidiary is None and not opts.urn:
                raise TypeError("Missing required property 'ovh_subsidiary'")
            __props__.__dict__["ovh_subsidiary"] = ovh_subsidiary
            if payment_mean is not None and not opts.urn:
                warnings.warn("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""", DeprecationWarning)
                pulumi.log.warn("""payment_mean is deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
            __props__.__dict__["payment_mean"] = payment_mean
            if plan is None and not opts.urn:
                raise TypeError("Missing required property 'plan'")
            __props__.__dict__["plan"] = plan
            __props__.__dict__["plan_options"] = plan_options
            __props__.__dict__["orders"] = None
            __props__.__dict__["service_name"] = None
        super(Vrack, __self__).__init__(
            'ovh:Vrack/vrack:Vrack',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            orders: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VrackOrderArgs']]]]] = None,
            ovh_subsidiary: Optional[pulumi.Input[str]] = None,
            payment_mean: Optional[pulumi.Input[str]] = None,
            plan: Optional[pulumi.Input[pulumi.InputType['VrackPlanArgs']]] = None,
            plan_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VrackPlanOptionArgs']]]]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'Vrack':
        """
        Get an existing Vrack resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: yourvrackdescription
        :param pulumi.Input[str] name: yourvrackname
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VrackOrderArgs']]]] orders: Details about an Order
        :param pulumi.Input[str] ovh_subsidiary: OVHcloud Subsidiary
        :param pulumi.Input[str] payment_mean: Ovh payment mode
        :param pulumi.Input[pulumi.InputType['VrackPlanArgs']] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VrackPlanOptionArgs']]]] plan_options: Product Plan to order
        :param pulumi.Input[str] service_name: The internal name of your vrack
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VrackState.__new__(_VrackState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["orders"] = orders
        __props__.__dict__["ovh_subsidiary"] = ovh_subsidiary
        __props__.__dict__["payment_mean"] = payment_mean
        __props__.__dict__["plan"] = plan
        __props__.__dict__["plan_options"] = plan_options
        __props__.__dict__["service_name"] = service_name
        return Vrack(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        yourvrackdescription
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        yourvrackname
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def orders(self) -> pulumi.Output[Sequence['outputs.VrackOrder']]:
        """
        Details about an Order
        """
        return pulumi.get(self, "orders")

    @property
    @pulumi.getter(name="ovhSubsidiary")
    def ovh_subsidiary(self) -> pulumi.Output[str]:
        """
        OVHcloud Subsidiary
        """
        return pulumi.get(self, "ovh_subsidiary")

    @property
    @pulumi.getter(name="paymentMean")
    def payment_mean(self) -> pulumi.Output[Optional[str]]:
        """
        Ovh payment mode
        """
        return pulumi.get(self, "payment_mean")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output['outputs.VrackPlan']:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="planOptions")
    def plan_options(self) -> pulumi.Output[Optional[Sequence['outputs.VrackPlanOption']]]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan_options")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The internal name of your vrack
        """
        return pulumi.get(self, "service_name")

