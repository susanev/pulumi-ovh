// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vrack

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attach a Dedicated Server Network Interface to a VRack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/Vrack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vrack.NewDedicatedServerInterface(ctx, "vdsi", &Vrack.DedicatedServerInterfaceArgs{
//				InterfaceId: pulumi.String("67890"),
//				ServiceName: pulumi.String("12345"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DedicatedServerInterface struct {
	pulumi.CustomResourceState

	// The id of dedicated server network interface.
	InterfaceId pulumi.StringOutput `pulumi:"interfaceId"`
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewDedicatedServerInterface registers a new resource with the given unique name, arguments, and options.
func NewDedicatedServerInterface(ctx *pulumi.Context,
	name string, args *DedicatedServerInterfaceArgs, opts ...pulumi.ResourceOption) (*DedicatedServerInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'InterfaceId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DedicatedServerInterface
	err := ctx.RegisterResource("ovh:Vrack/dedicatedServerInterface:DedicatedServerInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedServerInterface gets an existing DedicatedServerInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedServerInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedServerInterfaceState, opts ...pulumi.ResourceOption) (*DedicatedServerInterface, error) {
	var resource DedicatedServerInterface
	err := ctx.ReadResource("ovh:Vrack/dedicatedServerInterface:DedicatedServerInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedServerInterface resources.
type dedicatedServerInterfaceState struct {
	// The id of dedicated server network interface.
	InterfaceId *string `pulumi:"interfaceId"`
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
}

type DedicatedServerInterfaceState struct {
	// The id of dedicated server network interface.
	InterfaceId pulumi.StringPtrInput
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
}

func (DedicatedServerInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedServerInterfaceState)(nil)).Elem()
}

type dedicatedServerInterfaceArgs struct {
	// The id of dedicated server network interface.
	InterfaceId string `pulumi:"interfaceId"`
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a DedicatedServerInterface resource.
type DedicatedServerInterfaceArgs struct {
	// The id of dedicated server network interface.
	InterfaceId pulumi.StringInput
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (DedicatedServerInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedServerInterfaceArgs)(nil)).Elem()
}

type DedicatedServerInterfaceInput interface {
	pulumi.Input

	ToDedicatedServerInterfaceOutput() DedicatedServerInterfaceOutput
	ToDedicatedServerInterfaceOutputWithContext(ctx context.Context) DedicatedServerInterfaceOutput
}

func (*DedicatedServerInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedServerInterface)(nil)).Elem()
}

func (i *DedicatedServerInterface) ToDedicatedServerInterfaceOutput() DedicatedServerInterfaceOutput {
	return i.ToDedicatedServerInterfaceOutputWithContext(context.Background())
}

func (i *DedicatedServerInterface) ToDedicatedServerInterfaceOutputWithContext(ctx context.Context) DedicatedServerInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedServerInterfaceOutput)
}

// DedicatedServerInterfaceArrayInput is an input type that accepts DedicatedServerInterfaceArray and DedicatedServerInterfaceArrayOutput values.
// You can construct a concrete instance of `DedicatedServerInterfaceArrayInput` via:
//
//	DedicatedServerInterfaceArray{ DedicatedServerInterfaceArgs{...} }
type DedicatedServerInterfaceArrayInput interface {
	pulumi.Input

	ToDedicatedServerInterfaceArrayOutput() DedicatedServerInterfaceArrayOutput
	ToDedicatedServerInterfaceArrayOutputWithContext(context.Context) DedicatedServerInterfaceArrayOutput
}

type DedicatedServerInterfaceArray []DedicatedServerInterfaceInput

func (DedicatedServerInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedServerInterface)(nil)).Elem()
}

func (i DedicatedServerInterfaceArray) ToDedicatedServerInterfaceArrayOutput() DedicatedServerInterfaceArrayOutput {
	return i.ToDedicatedServerInterfaceArrayOutputWithContext(context.Background())
}

func (i DedicatedServerInterfaceArray) ToDedicatedServerInterfaceArrayOutputWithContext(ctx context.Context) DedicatedServerInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedServerInterfaceArrayOutput)
}

// DedicatedServerInterfaceMapInput is an input type that accepts DedicatedServerInterfaceMap and DedicatedServerInterfaceMapOutput values.
// You can construct a concrete instance of `DedicatedServerInterfaceMapInput` via:
//
//	DedicatedServerInterfaceMap{ "key": DedicatedServerInterfaceArgs{...} }
type DedicatedServerInterfaceMapInput interface {
	pulumi.Input

	ToDedicatedServerInterfaceMapOutput() DedicatedServerInterfaceMapOutput
	ToDedicatedServerInterfaceMapOutputWithContext(context.Context) DedicatedServerInterfaceMapOutput
}

type DedicatedServerInterfaceMap map[string]DedicatedServerInterfaceInput

func (DedicatedServerInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedServerInterface)(nil)).Elem()
}

func (i DedicatedServerInterfaceMap) ToDedicatedServerInterfaceMapOutput() DedicatedServerInterfaceMapOutput {
	return i.ToDedicatedServerInterfaceMapOutputWithContext(context.Background())
}

func (i DedicatedServerInterfaceMap) ToDedicatedServerInterfaceMapOutputWithContext(ctx context.Context) DedicatedServerInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedServerInterfaceMapOutput)
}

type DedicatedServerInterfaceOutput struct{ *pulumi.OutputState }

func (DedicatedServerInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedServerInterface)(nil)).Elem()
}

func (o DedicatedServerInterfaceOutput) ToDedicatedServerInterfaceOutput() DedicatedServerInterfaceOutput {
	return o
}

func (o DedicatedServerInterfaceOutput) ToDedicatedServerInterfaceOutputWithContext(ctx context.Context) DedicatedServerInterfaceOutput {
	return o
}

// The id of dedicated server network interface.
func (o DedicatedServerInterfaceOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerInterface) pulumi.StringOutput { return v.InterfaceId }).(pulumi.StringOutput)
}

// The id of the vrack. If omitted,
// the `OVH_VRACK_SERVICE` environment variable is used.
func (o DedicatedServerInterfaceOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServerInterface) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type DedicatedServerInterfaceArrayOutput struct{ *pulumi.OutputState }

func (DedicatedServerInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedServerInterface)(nil)).Elem()
}

func (o DedicatedServerInterfaceArrayOutput) ToDedicatedServerInterfaceArrayOutput() DedicatedServerInterfaceArrayOutput {
	return o
}

func (o DedicatedServerInterfaceArrayOutput) ToDedicatedServerInterfaceArrayOutputWithContext(ctx context.Context) DedicatedServerInterfaceArrayOutput {
	return o
}

func (o DedicatedServerInterfaceArrayOutput) Index(i pulumi.IntInput) DedicatedServerInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DedicatedServerInterface {
		return vs[0].([]*DedicatedServerInterface)[vs[1].(int)]
	}).(DedicatedServerInterfaceOutput)
}

type DedicatedServerInterfaceMapOutput struct{ *pulumi.OutputState }

func (DedicatedServerInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedServerInterface)(nil)).Elem()
}

func (o DedicatedServerInterfaceMapOutput) ToDedicatedServerInterfaceMapOutput() DedicatedServerInterfaceMapOutput {
	return o
}

func (o DedicatedServerInterfaceMapOutput) ToDedicatedServerInterfaceMapOutputWithContext(ctx context.Context) DedicatedServerInterfaceMapOutput {
	return o
}

func (o DedicatedServerInterfaceMapOutput) MapIndex(k pulumi.StringInput) DedicatedServerInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DedicatedServerInterface {
		return vs[0].(map[string]*DedicatedServerInterface)[vs[1].(string)]
	}).(DedicatedServerInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedServerInterfaceInput)(nil)).Elem(), &DedicatedServerInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedServerInterfaceArrayInput)(nil)).Elem(), DedicatedServerInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedServerInterfaceMapInput)(nil)).Elem(), DedicatedServerInterfaceMap{})
	pulumi.RegisterOutputType(DedicatedServerInterfaceOutput{})
	pulumi.RegisterOutputType(DedicatedServerInterfaceArrayOutput{})
	pulumi.RegisterOutputType(DedicatedServerInterfaceMapOutput{})
}
