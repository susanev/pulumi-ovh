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

    public sealed class KubeCustomizationApiserverGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("admissionplugins")]
        public Input<Inputs.KubeCustomizationApiserverAdmissionpluginsGetArgs>? Admissionplugins { get; set; }

        public KubeCustomizationApiserverGetArgs()
        {
        }
        public static new KubeCustomizationApiserverGetArgs Empty => new KubeCustomizationApiserverGetArgs();
    }
}
