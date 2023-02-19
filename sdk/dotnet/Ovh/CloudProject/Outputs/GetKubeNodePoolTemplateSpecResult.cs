// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.CloudProject.Outputs
{

    [OutputType]
    public sealed class GetKubeNodePoolTemplateSpecResult
    {
        public readonly ImmutableArray<ImmutableDictionary<string, object>> Taints;
        public readonly bool? Unschedulable;

        [OutputConstructor]
        private GetKubeNodePoolTemplateSpecResult(
            ImmutableArray<ImmutableDictionary<string, object>> taints,

            bool? unschedulable)
        {
            Taints = taints;
            Unschedulable = unschedulable;
        }
    }
}
