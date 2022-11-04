// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.CloudProject
{
    public static class GetUserS3Credentials
    {
        /// <summary>
        /// Use this data source to retrieve the list of all the S3 access_key_id associated with a public cloud project's user.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myS3Credentials = Ovh.CloudProject.GetUserS3Credentials.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         UserId = "1234",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["accessKeyIds"] = myS3Credentials.Apply(getUserS3CredentialsResult =&gt; getUserS3CredentialsResult.AccessKeyIds),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUserS3CredentialsResult> InvokeAsync(GetUserS3CredentialsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserS3CredentialsResult>("ovh:CloudProject/getUserS3Credentials:getUserS3Credentials", args ?? new GetUserS3CredentialsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the list of all the S3 access_key_id associated with a public cloud project's user.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myS3Credentials = Ovh.CloudProject.GetUserS3Credentials.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         UserId = "1234",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["accessKeyIds"] = myS3Credentials.Apply(getUserS3CredentialsResult =&gt; getUserS3CredentialsResult.AccessKeyIds),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUserS3CredentialsResult> Invoke(GetUserS3CredentialsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserS3CredentialsResult>("ovh:CloudProject/getUserS3Credentials:getUserS3Credentials", args ?? new GetUserS3CredentialsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserS3CredentialsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        /// <summary>
        /// The ID of a public cloud project's user.
        /// </summary>
        [Input("userId", required: true)]
        public string UserId { get; set; } = null!;

        public GetUserS3CredentialsArgs()
        {
        }
        public static new GetUserS3CredentialsArgs Empty => new GetUserS3CredentialsArgs();
    }

    public sealed class GetUserS3CredentialsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// The ID of a public cloud project's user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public GetUserS3CredentialsInvokeArgs()
        {
        }
        public static new GetUserS3CredentialsInvokeArgs Empty => new GetUserS3CredentialsInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserS3CredentialsResult
    {
        /// <summary>
        /// The list of the Access Key ID associated with this user.
        /// </summary>
        public readonly ImmutableArray<string> AccessKeyIds;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ServiceName;
        public readonly string UserId;

        [OutputConstructor]
        private GetUserS3CredentialsResult(
            ImmutableArray<string> accessKeyIds,

            string id,

            string serviceName,

            string userId)
        {
            AccessKeyIds = accessKeyIds;
            Id = id;
            ServiceName = serviceName;
            UserId = userId;
        }
    }
}
