// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an IPXE Script.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//	"os"
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/Me"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Me.NewIpxeScript(ctx, "script", &Me.IpxeScriptArgs{
//				Script: readFileOrPanic(fmt.Sprintf("%v/boot.ipxe", path.Module)),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type IpxeScript struct {
	pulumi.CustomResourceState

	// For documentation purpose only. This attribute is not passed to the OVHcloud API as it cannot be retrieved back. Instead a fake description ('$name auto description') is passed at creation time.
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the IPXE Script.
	Name pulumi.StringOutput `pulumi:"name"`
	// The content of the script.
	Script pulumi.StringOutput `pulumi:"script"`
}

// NewIpxeScript registers a new resource with the given unique name, arguments, and options.
func NewIpxeScript(ctx *pulumi.Context,
	name string, args *IpxeScriptArgs, opts ...pulumi.ResourceOption) (*IpxeScript, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Script == nil {
		return nil, errors.New("invalid value for required argument 'Script'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IpxeScript
	err := ctx.RegisterResource("ovh:Me/ipxeScript:IpxeScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpxeScript gets an existing IpxeScript resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpxeScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpxeScriptState, opts ...pulumi.ResourceOption) (*IpxeScript, error) {
	var resource IpxeScript
	err := ctx.ReadResource("ovh:Me/ipxeScript:IpxeScript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpxeScript resources.
type ipxeScriptState struct {
	// For documentation purpose only. This attribute is not passed to the OVHcloud API as it cannot be retrieved back. Instead a fake description ('$name auto description') is passed at creation time.
	Description *string `pulumi:"description"`
	// The name of the IPXE Script.
	Name *string `pulumi:"name"`
	// The content of the script.
	Script *string `pulumi:"script"`
}

type IpxeScriptState struct {
	// For documentation purpose only. This attribute is not passed to the OVHcloud API as it cannot be retrieved back. Instead a fake description ('$name auto description') is passed at creation time.
	Description pulumi.StringPtrInput
	// The name of the IPXE Script.
	Name pulumi.StringPtrInput
	// The content of the script.
	Script pulumi.StringPtrInput
}

func (IpxeScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipxeScriptState)(nil)).Elem()
}

type ipxeScriptArgs struct {
	// For documentation purpose only. This attribute is not passed to the OVHcloud API as it cannot be retrieved back. Instead a fake description ('$name auto description') is passed at creation time.
	Description *string `pulumi:"description"`
	// The name of the IPXE Script.
	Name *string `pulumi:"name"`
	// The content of the script.
	Script string `pulumi:"script"`
}

// The set of arguments for constructing a IpxeScript resource.
type IpxeScriptArgs struct {
	// For documentation purpose only. This attribute is not passed to the OVHcloud API as it cannot be retrieved back. Instead a fake description ('$name auto description') is passed at creation time.
	Description pulumi.StringPtrInput
	// The name of the IPXE Script.
	Name pulumi.StringPtrInput
	// The content of the script.
	Script pulumi.StringInput
}

func (IpxeScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipxeScriptArgs)(nil)).Elem()
}

type IpxeScriptInput interface {
	pulumi.Input

	ToIpxeScriptOutput() IpxeScriptOutput
	ToIpxeScriptOutputWithContext(ctx context.Context) IpxeScriptOutput
}

func (*IpxeScript) ElementType() reflect.Type {
	return reflect.TypeOf((**IpxeScript)(nil)).Elem()
}

func (i *IpxeScript) ToIpxeScriptOutput() IpxeScriptOutput {
	return i.ToIpxeScriptOutputWithContext(context.Background())
}

func (i *IpxeScript) ToIpxeScriptOutputWithContext(ctx context.Context) IpxeScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpxeScriptOutput)
}

// IpxeScriptArrayInput is an input type that accepts IpxeScriptArray and IpxeScriptArrayOutput values.
// You can construct a concrete instance of `IpxeScriptArrayInput` via:
//
//	IpxeScriptArray{ IpxeScriptArgs{...} }
type IpxeScriptArrayInput interface {
	pulumi.Input

	ToIpxeScriptArrayOutput() IpxeScriptArrayOutput
	ToIpxeScriptArrayOutputWithContext(context.Context) IpxeScriptArrayOutput
}

type IpxeScriptArray []IpxeScriptInput

func (IpxeScriptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpxeScript)(nil)).Elem()
}

func (i IpxeScriptArray) ToIpxeScriptArrayOutput() IpxeScriptArrayOutput {
	return i.ToIpxeScriptArrayOutputWithContext(context.Background())
}

func (i IpxeScriptArray) ToIpxeScriptArrayOutputWithContext(ctx context.Context) IpxeScriptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpxeScriptArrayOutput)
}

// IpxeScriptMapInput is an input type that accepts IpxeScriptMap and IpxeScriptMapOutput values.
// You can construct a concrete instance of `IpxeScriptMapInput` via:
//
//	IpxeScriptMap{ "key": IpxeScriptArgs{...} }
type IpxeScriptMapInput interface {
	pulumi.Input

	ToIpxeScriptMapOutput() IpxeScriptMapOutput
	ToIpxeScriptMapOutputWithContext(context.Context) IpxeScriptMapOutput
}

type IpxeScriptMap map[string]IpxeScriptInput

func (IpxeScriptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpxeScript)(nil)).Elem()
}

func (i IpxeScriptMap) ToIpxeScriptMapOutput() IpxeScriptMapOutput {
	return i.ToIpxeScriptMapOutputWithContext(context.Background())
}

func (i IpxeScriptMap) ToIpxeScriptMapOutputWithContext(ctx context.Context) IpxeScriptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpxeScriptMapOutput)
}

type IpxeScriptOutput struct{ *pulumi.OutputState }

func (IpxeScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpxeScript)(nil)).Elem()
}

func (o IpxeScriptOutput) ToIpxeScriptOutput() IpxeScriptOutput {
	return o
}

func (o IpxeScriptOutput) ToIpxeScriptOutputWithContext(ctx context.Context) IpxeScriptOutput {
	return o
}

// For documentation purpose only. This attribute is not passed to the OVHcloud API as it cannot be retrieved back. Instead a fake description ('$name auto description') is passed at creation time.
func (o IpxeScriptOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *IpxeScript) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The name of the IPXE Script.
func (o IpxeScriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpxeScript) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The content of the script.
func (o IpxeScriptOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v *IpxeScript) pulumi.StringOutput { return v.Script }).(pulumi.StringOutput)
}

type IpxeScriptArrayOutput struct{ *pulumi.OutputState }

func (IpxeScriptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpxeScript)(nil)).Elem()
}

func (o IpxeScriptArrayOutput) ToIpxeScriptArrayOutput() IpxeScriptArrayOutput {
	return o
}

func (o IpxeScriptArrayOutput) ToIpxeScriptArrayOutputWithContext(ctx context.Context) IpxeScriptArrayOutput {
	return o
}

func (o IpxeScriptArrayOutput) Index(i pulumi.IntInput) IpxeScriptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpxeScript {
		return vs[0].([]*IpxeScript)[vs[1].(int)]
	}).(IpxeScriptOutput)
}

type IpxeScriptMapOutput struct{ *pulumi.OutputState }

func (IpxeScriptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpxeScript)(nil)).Elem()
}

func (o IpxeScriptMapOutput) ToIpxeScriptMapOutput() IpxeScriptMapOutput {
	return o
}

func (o IpxeScriptMapOutput) ToIpxeScriptMapOutputWithContext(ctx context.Context) IpxeScriptMapOutput {
	return o
}

func (o IpxeScriptMapOutput) MapIndex(k pulumi.StringInput) IpxeScriptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpxeScript {
		return vs[0].(map[string]*IpxeScript)[vs[1].(string)]
	}).(IpxeScriptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpxeScriptInput)(nil)).Elem(), &IpxeScript{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpxeScriptArrayInput)(nil)).Elem(), IpxeScriptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpxeScriptMapInput)(nil)).Elem(), IpxeScriptMap{})
	pulumi.RegisterOutputType(IpxeScriptOutput{})
	pulumi.RegisterOutputType(IpxeScriptArrayOutput{})
	pulumi.RegisterOutputType(IpxeScriptMapOutput{})
}
