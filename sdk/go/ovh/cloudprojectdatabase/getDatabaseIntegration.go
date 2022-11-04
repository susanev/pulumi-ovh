// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about an integration of a database cluster associated with a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/CloudProjectDatabase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			integration, err := CloudProjectDatabase.GetDatabaseIntegration(ctx, &cloudprojectdatabase.GetDatabaseIntegrationArgs{
//				ServiceName: "XXX",
//				Engine:      "YYY",
//				ClusterId:   "ZZZ",
//				Id:          "UUU",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("integrationType", integration.Type)
//			return nil
//		})
//	}
//
// ```
func GetDatabaseIntegration(ctx *pulumi.Context, args *GetDatabaseIntegrationArgs, opts ...pulumi.InvokeOption) (*GetDatabaseIntegrationResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDatabaseIntegrationResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getDatabaseIntegration:getDatabaseIntegration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseIntegration.
type GetDatabaseIntegrationArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// All engines available exept `mongodb`
	Engine string `pulumi:"engine"`
	// Integration ID
	Id string `pulumi:"id"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getDatabaseIntegration.
type GetDatabaseIntegrationResult struct {
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// ID of the destination service.
	DestinationServiceId string `pulumi:"destinationServiceId"`
	// See Argument Reference above.
	Engine string `pulumi:"engine"`
	// See Argument Reference above.
	Id string `pulumi:"id"`
	// Parameters for the integration.
	Parameters map[string]string `pulumi:"parameters"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
	// ID of the source service.
	SourceServiceId string `pulumi:"sourceServiceId"`
	// Current status of the integration.
	Status string `pulumi:"status"`
	// Type of the integration.
	Type string `pulumi:"type"`
}

func GetDatabaseIntegrationOutput(ctx *pulumi.Context, args GetDatabaseIntegrationOutputArgs, opts ...pulumi.InvokeOption) GetDatabaseIntegrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabaseIntegrationResult, error) {
			args := v.(GetDatabaseIntegrationArgs)
			r, err := GetDatabaseIntegration(ctx, &args, opts...)
			var s GetDatabaseIntegrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatabaseIntegrationResultOutput)
}

// A collection of arguments for invoking getDatabaseIntegration.
type GetDatabaseIntegrationOutputArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// All engines available exept `mongodb`
	Engine pulumi.StringInput `pulumi:"engine"`
	// Integration ID
	Id pulumi.StringInput `pulumi:"id"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetDatabaseIntegrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseIntegrationArgs)(nil)).Elem()
}

// A collection of values returned by getDatabaseIntegration.
type GetDatabaseIntegrationResultOutput struct{ *pulumi.OutputState }

func (GetDatabaseIntegrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseIntegrationResult)(nil)).Elem()
}

func (o GetDatabaseIntegrationResultOutput) ToGetDatabaseIntegrationResultOutput() GetDatabaseIntegrationResultOutput {
	return o
}

func (o GetDatabaseIntegrationResultOutput) ToGetDatabaseIntegrationResultOutputWithContext(ctx context.Context) GetDatabaseIntegrationResultOutput {
	return o
}

// See Argument Reference above.
func (o GetDatabaseIntegrationResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseIntegrationResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// ID of the destination service.
func (o GetDatabaseIntegrationResultOutput) DestinationServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseIntegrationResult) string { return v.DestinationServiceId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetDatabaseIntegrationResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseIntegrationResult) string { return v.Engine }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetDatabaseIntegrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseIntegrationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Parameters for the integration.
func (o GetDatabaseIntegrationResultOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetDatabaseIntegrationResult) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

// See Argument Reference above.
func (o GetDatabaseIntegrationResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseIntegrationResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// ID of the source service.
func (o GetDatabaseIntegrationResultOutput) SourceServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseIntegrationResult) string { return v.SourceServiceId }).(pulumi.StringOutput)
}

// Current status of the integration.
func (o GetDatabaseIntegrationResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseIntegrationResult) string { return v.Status }).(pulumi.StringOutput)
}

// Type of the integration.
func (o GetDatabaseIntegrationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseIntegrationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabaseIntegrationResultOutput{})
}
