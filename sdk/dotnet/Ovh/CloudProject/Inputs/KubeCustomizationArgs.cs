// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.CloudProject.Inputs
{

    public sealed class KubeCustomizationArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiserver")]
        public Input<Inputs.KubeCustomizationApiserverArgs>? Apiserver { get; set; }

        public KubeCustomizationArgs()
        {
        }
        public static new KubeCustomizationArgs Empty => new KubeCustomizationArgs();
    }
}
