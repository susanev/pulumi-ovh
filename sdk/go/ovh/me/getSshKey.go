// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about an SSH key.
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
//			_, err := Me.GetSshKey(ctx, &me.GetSshKeyArgs{
//				KeyName: "mykey",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSshKey(ctx *pulumi.Context, args *LookupSshKeyArgs, opts ...pulumi.InvokeOption) (*LookupSshKeyResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSshKeyResult
	err := ctx.Invoke("ovh:Me/getSshKey:getSshKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSshKey.
type LookupSshKeyArgs struct {
	// The name of the SSH key.
	KeyName string `pulumi:"keyName"`
}

// A collection of values returned by getSshKey.
type LookupSshKeyResult struct {
	// True when this public SSH key is used for rescue mode and reinstallations.
	Default bool `pulumi:"default"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The content of the public key.
	// E.g.: "ssh-ed25519 AAAAC3..."
	Key string `pulumi:"key"`
	// See Argument Reference above.
	KeyName string `pulumi:"keyName"`
}

func LookupSshKeyOutput(ctx *pulumi.Context, args LookupSshKeyOutputArgs, opts ...pulumi.InvokeOption) LookupSshKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSshKeyResult, error) {
			args := v.(LookupSshKeyArgs)
			r, err := LookupSshKey(ctx, &args, opts...)
			var s LookupSshKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSshKeyResultOutput)
}

// A collection of arguments for invoking getSshKey.
type LookupSshKeyOutputArgs struct {
	// The name of the SSH key.
	KeyName pulumi.StringInput `pulumi:"keyName"`
}

func (LookupSshKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSshKeyArgs)(nil)).Elem()
}

// A collection of values returned by getSshKey.
type LookupSshKeyResultOutput struct{ *pulumi.OutputState }

func (LookupSshKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSshKeyResult)(nil)).Elem()
}

func (o LookupSshKeyResultOutput) ToLookupSshKeyResultOutput() LookupSshKeyResultOutput {
	return o
}

func (o LookupSshKeyResultOutput) ToLookupSshKeyResultOutputWithContext(ctx context.Context) LookupSshKeyResultOutput {
	return o
}

// True when this public SSH key is used for rescue mode and reinstallations.
func (o LookupSshKeyResultOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSshKeyResult) bool { return v.Default }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSshKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The content of the public key.
// E.g.: "ssh-ed25519 AAAAC3..."
func (o LookupSshKeyResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshKeyResult) string { return v.Key }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupSshKeyResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshKeyResult) string { return v.KeyName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSshKeyResultOutput{})
}
