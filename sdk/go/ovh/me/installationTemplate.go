// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create a custom installation template available for dedicated servers.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/Me"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Me.NewInstallationTemplate(ctx, "mytemplate", &Me.InstallationTemplateArgs{
//				BaseTemplateName: pulumi.String("centos7_64"),
//				DefaultLanguage:  pulumi.String("en"),
//				TemplateName:     pulumi.String("mytemplate"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Custom installation template available for dedicated servers can be imported using the `base_template_name`, `template_name` of the cluster, separated by "/" E.g., bash
//
// ```sh
//
//	$ pulumi import ovh:Me/installationTemplate:InstallationTemplate mytemplate base_template_name/template_name
//
// ```
type InstallationTemplate struct {
	pulumi.CustomResourceState

	// List of all language available for this template.
	AvailableLanguages pulumi.StringArrayOutput `pulumi:"availableLanguages"`
	// The name of an existing installation template, choose one among the list given by `getInstallationTemplates` datasource.
	BaseTemplateName pulumi.StringOutput `pulumi:"baseTemplateName"`
	// This distribution is new and, although tested and functional, may still display odd behaviour.
	Beta pulumi.BoolOutput `pulumi:"beta"`
	// This template bit format (32 or 64).
	BitFormat pulumi.IntOutput `pulumi:"bitFormat"`
	// Category of this template (informative only). (basic, customer, hosting, other, readyToUse, virtualisation).
	Category      pulumi.StringOutput                        `pulumi:"category"`
	Customization InstallationTemplateCustomizationPtrOutput `pulumi:"customization"`
	// The default language of this template.
	DefaultLanguage pulumi.StringOutput `pulumi:"defaultLanguage"`
	// is this distribution deprecated.
	Deprecated pulumi.BoolOutput `pulumi:"deprecated"`
	// information about this template.
	Description pulumi.StringOutput `pulumi:"description"`
	// the distribution this template is based on.
	Distribution pulumi.StringOutput `pulumi:"distribution"`
	// this template family type (bsd,linux,solaris,windows).
	Family pulumi.StringOutput `pulumi:"family"`
	// Filesystems available (btrfs,ext3,ext4,ntfs,reiserfs,swap,ufs,xfs,zfs).
	Filesystems pulumi.StringArrayOutput `pulumi:"filesystems"`
	// This distribution supports hardware raid configuration through the OVHcloud API.
	HardRaidConfiguration pulumi.BoolOutput `pulumi:"hardRaidConfiguration"`
	// Date of last modification of the base image.
	LastModification pulumi.StringOutput `pulumi:"lastModification"`
	// This distribution supports Logical Volumes (Linux LVM)
	LvmReady pulumi.BoolOutput `pulumi:"lvmReady"`
	// Remove default partition schemes at creation.
	RemoveDefaultPartitionSchemes pulumi.BoolOutput `pulumi:"removeDefaultPartitionSchemes"`
	// This distribution supports installation using the distribution's native kernel instead of the recommended OVHcloud kernel.
	SupportsDistributionKernel pulumi.BoolOutput `pulumi:"supportsDistributionKernel"`
	// This distribution supports RTM software.
	SupportsRtm pulumi.BoolOutput `pulumi:"supportsRtm"`
	// This distribution supports the microsoft SQL server.
	SupportsSqlServer pulumi.BoolOutput `pulumi:"supportsSqlServer"`
	// This template name.
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
}

// NewInstallationTemplate registers a new resource with the given unique name, arguments, and options.
func NewInstallationTemplate(ctx *pulumi.Context,
	name string, args *InstallationTemplateArgs, opts ...pulumi.ResourceOption) (*InstallationTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaseTemplateName == nil {
		return nil, errors.New("invalid value for required argument 'BaseTemplateName'")
	}
	if args.DefaultLanguage == nil {
		return nil, errors.New("invalid value for required argument 'DefaultLanguage'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InstallationTemplate
	err := ctx.RegisterResource("ovh:Me/installationTemplate:InstallationTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstallationTemplate gets an existing InstallationTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstallationTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstallationTemplateState, opts ...pulumi.ResourceOption) (*InstallationTemplate, error) {
	var resource InstallationTemplate
	err := ctx.ReadResource("ovh:Me/installationTemplate:InstallationTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstallationTemplate resources.
type installationTemplateState struct {
	// List of all language available for this template.
	AvailableLanguages []string `pulumi:"availableLanguages"`
	// The name of an existing installation template, choose one among the list given by `getInstallationTemplates` datasource.
	BaseTemplateName *string `pulumi:"baseTemplateName"`
	// This distribution is new and, although tested and functional, may still display odd behaviour.
	Beta *bool `pulumi:"beta"`
	// This template bit format (32 or 64).
	BitFormat *int `pulumi:"bitFormat"`
	// Category of this template (informative only). (basic, customer, hosting, other, readyToUse, virtualisation).
	Category      *string                            `pulumi:"category"`
	Customization *InstallationTemplateCustomization `pulumi:"customization"`
	// The default language of this template.
	DefaultLanguage *string `pulumi:"defaultLanguage"`
	// is this distribution deprecated.
	Deprecated *bool `pulumi:"deprecated"`
	// information about this template.
	Description *string `pulumi:"description"`
	// the distribution this template is based on.
	Distribution *string `pulumi:"distribution"`
	// this template family type (bsd,linux,solaris,windows).
	Family *string `pulumi:"family"`
	// Filesystems available (btrfs,ext3,ext4,ntfs,reiserfs,swap,ufs,xfs,zfs).
	Filesystems []string `pulumi:"filesystems"`
	// This distribution supports hardware raid configuration through the OVHcloud API.
	HardRaidConfiguration *bool `pulumi:"hardRaidConfiguration"`
	// Date of last modification of the base image.
	LastModification *string `pulumi:"lastModification"`
	// This distribution supports Logical Volumes (Linux LVM)
	LvmReady *bool `pulumi:"lvmReady"`
	// Remove default partition schemes at creation.
	RemoveDefaultPartitionSchemes *bool `pulumi:"removeDefaultPartitionSchemes"`
	// This distribution supports installation using the distribution's native kernel instead of the recommended OVHcloud kernel.
	SupportsDistributionKernel *bool `pulumi:"supportsDistributionKernel"`
	// This distribution supports RTM software.
	SupportsRtm *bool `pulumi:"supportsRtm"`
	// This distribution supports the microsoft SQL server.
	SupportsSqlServer *bool `pulumi:"supportsSqlServer"`
	// This template name.
	TemplateName *string `pulumi:"templateName"`
}

type InstallationTemplateState struct {
	// List of all language available for this template.
	AvailableLanguages pulumi.StringArrayInput
	// The name of an existing installation template, choose one among the list given by `getInstallationTemplates` datasource.
	BaseTemplateName pulumi.StringPtrInput
	// This distribution is new and, although tested and functional, may still display odd behaviour.
	Beta pulumi.BoolPtrInput
	// This template bit format (32 or 64).
	BitFormat pulumi.IntPtrInput
	// Category of this template (informative only). (basic, customer, hosting, other, readyToUse, virtualisation).
	Category      pulumi.StringPtrInput
	Customization InstallationTemplateCustomizationPtrInput
	// The default language of this template.
	DefaultLanguage pulumi.StringPtrInput
	// is this distribution deprecated.
	Deprecated pulumi.BoolPtrInput
	// information about this template.
	Description pulumi.StringPtrInput
	// the distribution this template is based on.
	Distribution pulumi.StringPtrInput
	// this template family type (bsd,linux,solaris,windows).
	Family pulumi.StringPtrInput
	// Filesystems available (btrfs,ext3,ext4,ntfs,reiserfs,swap,ufs,xfs,zfs).
	Filesystems pulumi.StringArrayInput
	// This distribution supports hardware raid configuration through the OVHcloud API.
	HardRaidConfiguration pulumi.BoolPtrInput
	// Date of last modification of the base image.
	LastModification pulumi.StringPtrInput
	// This distribution supports Logical Volumes (Linux LVM)
	LvmReady pulumi.BoolPtrInput
	// Remove default partition schemes at creation.
	RemoveDefaultPartitionSchemes pulumi.BoolPtrInput
	// This distribution supports installation using the distribution's native kernel instead of the recommended OVHcloud kernel.
	SupportsDistributionKernel pulumi.BoolPtrInput
	// This distribution supports RTM software.
	SupportsRtm pulumi.BoolPtrInput
	// This distribution supports the microsoft SQL server.
	SupportsSqlServer pulumi.BoolPtrInput
	// This template name.
	TemplateName pulumi.StringPtrInput
}

func (InstallationTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*installationTemplateState)(nil)).Elem()
}

type installationTemplateArgs struct {
	// The name of an existing installation template, choose one among the list given by `getInstallationTemplates` datasource.
	BaseTemplateName string                             `pulumi:"baseTemplateName"`
	Customization    *InstallationTemplateCustomization `pulumi:"customization"`
	// The default language of this template.
	DefaultLanguage string `pulumi:"defaultLanguage"`
	// Remove default partition schemes at creation.
	RemoveDefaultPartitionSchemes *bool `pulumi:"removeDefaultPartitionSchemes"`
	// This template name.
	TemplateName string `pulumi:"templateName"`
}

// The set of arguments for constructing a InstallationTemplate resource.
type InstallationTemplateArgs struct {
	// The name of an existing installation template, choose one among the list given by `getInstallationTemplates` datasource.
	BaseTemplateName pulumi.StringInput
	Customization    InstallationTemplateCustomizationPtrInput
	// The default language of this template.
	DefaultLanguage pulumi.StringInput
	// Remove default partition schemes at creation.
	RemoveDefaultPartitionSchemes pulumi.BoolPtrInput
	// This template name.
	TemplateName pulumi.StringInput
}

func (InstallationTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*installationTemplateArgs)(nil)).Elem()
}

type InstallationTemplateInput interface {
	pulumi.Input

	ToInstallationTemplateOutput() InstallationTemplateOutput
	ToInstallationTemplateOutputWithContext(ctx context.Context) InstallationTemplateOutput
}

func (*InstallationTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**InstallationTemplate)(nil)).Elem()
}

func (i *InstallationTemplate) ToInstallationTemplateOutput() InstallationTemplateOutput {
	return i.ToInstallationTemplateOutputWithContext(context.Background())
}

func (i *InstallationTemplate) ToInstallationTemplateOutputWithContext(ctx context.Context) InstallationTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstallationTemplateOutput)
}

// InstallationTemplateArrayInput is an input type that accepts InstallationTemplateArray and InstallationTemplateArrayOutput values.
// You can construct a concrete instance of `InstallationTemplateArrayInput` via:
//
//	InstallationTemplateArray{ InstallationTemplateArgs{...} }
type InstallationTemplateArrayInput interface {
	pulumi.Input

	ToInstallationTemplateArrayOutput() InstallationTemplateArrayOutput
	ToInstallationTemplateArrayOutputWithContext(context.Context) InstallationTemplateArrayOutput
}

type InstallationTemplateArray []InstallationTemplateInput

func (InstallationTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstallationTemplate)(nil)).Elem()
}

func (i InstallationTemplateArray) ToInstallationTemplateArrayOutput() InstallationTemplateArrayOutput {
	return i.ToInstallationTemplateArrayOutputWithContext(context.Background())
}

func (i InstallationTemplateArray) ToInstallationTemplateArrayOutputWithContext(ctx context.Context) InstallationTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstallationTemplateArrayOutput)
}

// InstallationTemplateMapInput is an input type that accepts InstallationTemplateMap and InstallationTemplateMapOutput values.
// You can construct a concrete instance of `InstallationTemplateMapInput` via:
//
//	InstallationTemplateMap{ "key": InstallationTemplateArgs{...} }
type InstallationTemplateMapInput interface {
	pulumi.Input

	ToInstallationTemplateMapOutput() InstallationTemplateMapOutput
	ToInstallationTemplateMapOutputWithContext(context.Context) InstallationTemplateMapOutput
}

type InstallationTemplateMap map[string]InstallationTemplateInput

func (InstallationTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstallationTemplate)(nil)).Elem()
}

func (i InstallationTemplateMap) ToInstallationTemplateMapOutput() InstallationTemplateMapOutput {
	return i.ToInstallationTemplateMapOutputWithContext(context.Background())
}

func (i InstallationTemplateMap) ToInstallationTemplateMapOutputWithContext(ctx context.Context) InstallationTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstallationTemplateMapOutput)
}

type InstallationTemplateOutput struct{ *pulumi.OutputState }

func (InstallationTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstallationTemplate)(nil)).Elem()
}

func (o InstallationTemplateOutput) ToInstallationTemplateOutput() InstallationTemplateOutput {
	return o
}

func (o InstallationTemplateOutput) ToInstallationTemplateOutputWithContext(ctx context.Context) InstallationTemplateOutput {
	return o
}

// List of all language available for this template.
func (o InstallationTemplateOutput) AvailableLanguages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.StringArrayOutput { return v.AvailableLanguages }).(pulumi.StringArrayOutput)
}

// The name of an existing installation template, choose one among the list given by `getInstallationTemplates` datasource.
func (o InstallationTemplateOutput) BaseTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.StringOutput { return v.BaseTemplateName }).(pulumi.StringOutput)
}

// This distribution is new and, although tested and functional, may still display odd behaviour.
func (o InstallationTemplateOutput) Beta() pulumi.BoolOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.BoolOutput { return v.Beta }).(pulumi.BoolOutput)
}

// This template bit format (32 or 64).
func (o InstallationTemplateOutput) BitFormat() pulumi.IntOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.IntOutput { return v.BitFormat }).(pulumi.IntOutput)
}

// Category of this template (informative only). (basic, customer, hosting, other, readyToUse, virtualisation).
func (o InstallationTemplateOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

func (o InstallationTemplateOutput) Customization() InstallationTemplateCustomizationPtrOutput {
	return o.ApplyT(func(v *InstallationTemplate) InstallationTemplateCustomizationPtrOutput { return v.Customization }).(InstallationTemplateCustomizationPtrOutput)
}

// The default language of this template.
func (o InstallationTemplateOutput) DefaultLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.StringOutput { return v.DefaultLanguage }).(pulumi.StringOutput)
}

// is this distribution deprecated.
func (o InstallationTemplateOutput) Deprecated() pulumi.BoolOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.BoolOutput { return v.Deprecated }).(pulumi.BoolOutput)
}

// information about this template.
func (o InstallationTemplateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// the distribution this template is based on.
func (o InstallationTemplateOutput) Distribution() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.StringOutput { return v.Distribution }).(pulumi.StringOutput)
}

// this template family type (bsd,linux,solaris,windows).
func (o InstallationTemplateOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.StringOutput { return v.Family }).(pulumi.StringOutput)
}

// Filesystems available (btrfs,ext3,ext4,ntfs,reiserfs,swap,ufs,xfs,zfs).
func (o InstallationTemplateOutput) Filesystems() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.StringArrayOutput { return v.Filesystems }).(pulumi.StringArrayOutput)
}

// This distribution supports hardware raid configuration through the OVHcloud API.
func (o InstallationTemplateOutput) HardRaidConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.BoolOutput { return v.HardRaidConfiguration }).(pulumi.BoolOutput)
}

// Date of last modification of the base image.
func (o InstallationTemplateOutput) LastModification() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.StringOutput { return v.LastModification }).(pulumi.StringOutput)
}

// This distribution supports Logical Volumes (Linux LVM)
func (o InstallationTemplateOutput) LvmReady() pulumi.BoolOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.BoolOutput { return v.LvmReady }).(pulumi.BoolOutput)
}

// Remove default partition schemes at creation.
func (o InstallationTemplateOutput) RemoveDefaultPartitionSchemes() pulumi.BoolOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.BoolOutput { return v.RemoveDefaultPartitionSchemes }).(pulumi.BoolOutput)
}

// This distribution supports installation using the distribution's native kernel instead of the recommended OVHcloud kernel.
func (o InstallationTemplateOutput) SupportsDistributionKernel() pulumi.BoolOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.BoolOutput { return v.SupportsDistributionKernel }).(pulumi.BoolOutput)
}

// This distribution supports RTM software.
func (o InstallationTemplateOutput) SupportsRtm() pulumi.BoolOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.BoolOutput { return v.SupportsRtm }).(pulumi.BoolOutput)
}

// This distribution supports the microsoft SQL server.
func (o InstallationTemplateOutput) SupportsSqlServer() pulumi.BoolOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.BoolOutput { return v.SupportsSqlServer }).(pulumi.BoolOutput)
}

// This template name.
func (o InstallationTemplateOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplate) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

type InstallationTemplateArrayOutput struct{ *pulumi.OutputState }

func (InstallationTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstallationTemplate)(nil)).Elem()
}

func (o InstallationTemplateArrayOutput) ToInstallationTemplateArrayOutput() InstallationTemplateArrayOutput {
	return o
}

func (o InstallationTemplateArrayOutput) ToInstallationTemplateArrayOutputWithContext(ctx context.Context) InstallationTemplateArrayOutput {
	return o
}

func (o InstallationTemplateArrayOutput) Index(i pulumi.IntInput) InstallationTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstallationTemplate {
		return vs[0].([]*InstallationTemplate)[vs[1].(int)]
	}).(InstallationTemplateOutput)
}

type InstallationTemplateMapOutput struct{ *pulumi.OutputState }

func (InstallationTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstallationTemplate)(nil)).Elem()
}

func (o InstallationTemplateMapOutput) ToInstallationTemplateMapOutput() InstallationTemplateMapOutput {
	return o
}

func (o InstallationTemplateMapOutput) ToInstallationTemplateMapOutputWithContext(ctx context.Context) InstallationTemplateMapOutput {
	return o
}

func (o InstallationTemplateMapOutput) MapIndex(k pulumi.StringInput) InstallationTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstallationTemplate {
		return vs[0].(map[string]*InstallationTemplate)[vs[1].(string)]
	}).(InstallationTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstallationTemplateInput)(nil)).Elem(), &InstallationTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstallationTemplateArrayInput)(nil)).Elem(), InstallationTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstallationTemplateMapInput)(nil)).Elem(), InstallationTemplateMap{})
	pulumi.RegisterOutputType(InstallationTemplateOutput{})
	pulumi.RegisterOutputType(InstallationTemplateArrayOutput{})
	pulumi.RegisterOutputType(InstallationTemplateMapOutput{})
}
