// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get a OVHcloud Managed Kubernetes Service cluster OIDC.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * export = async () => {
 *     const oidc = await ovh.CloudProject.getKubeOidc({
 *         serviceName: "XXXXXX",
 *         kubeId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
 *     });
 *     const oidc_val = oidc.clientId;
 *     return {
 *         "oidc-val": oidc_val,
 *     };
 * }
 * ```
 */
export function getKubeOidc(args: GetKubeOidcArgs, opts?: pulumi.InvokeOptions): Promise<GetKubeOidcResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProject/getKubeOidc:getKubeOidc", {
        "clientId": args.clientId,
        "issuerUrl": args.issuerUrl,
        "kubeId": args.kubeId,
        "oidcCaContent": args.oidcCaContent,
        "oidcGroupsClaims": args.oidcGroupsClaims,
        "oidcGroupsPrefix": args.oidcGroupsPrefix,
        "oidcRequiredClaims": args.oidcRequiredClaims,
        "oidcSigningAlgs": args.oidcSigningAlgs,
        "oidcUsernameClaim": args.oidcUsernameClaim,
        "oidcUsernamePrefix": args.oidcUsernamePrefix,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getKubeOidc.
 */
export interface GetKubeOidcArgs {
    /**
     * The OIDC client ID.
     */
    clientId?: string;
    /**
     * The OIDC issuer url.
     */
    issuerUrl?: string;
    /**
     * The id of the managed kubernetes cluster.
     */
    kubeId: string;
    oidcCaContent?: string;
    oidcGroupsClaims?: string[];
    oidcGroupsPrefix?: string;
    oidcRequiredClaims?: string[];
    oidcSigningAlgs?: string[];
    oidcUsernameClaim?: string;
    oidcUsernamePrefix?: string;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getKubeOidc.
 */
export interface GetKubeOidcResult {
    /**
     * The OIDC client ID.
     */
    readonly clientId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The OIDC issuer url.
     */
    readonly issuerUrl?: string;
    /**
     * See Argument Reference above.
     */
    readonly kubeId: string;
    readonly oidcCaContent?: string;
    readonly oidcGroupsClaims?: string[];
    readonly oidcGroupsPrefix?: string;
    readonly oidcRequiredClaims?: string[];
    readonly oidcSigningAlgs?: string[];
    readonly oidcUsernameClaim?: string;
    readonly oidcUsernamePrefix?: string;
    /**
     * See Argument Reference above.
     */
    readonly serviceName: string;
}
/**
 * Use this data source to get a OVHcloud Managed Kubernetes Service cluster OIDC.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * export = async () => {
 *     const oidc = await ovh.CloudProject.getKubeOidc({
 *         serviceName: "XXXXXX",
 *         kubeId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
 *     });
 *     const oidc_val = oidc.clientId;
 *     return {
 *         "oidc-val": oidc_val,
 *     };
 * }
 * ```
 */
export function getKubeOidcOutput(args: GetKubeOidcOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKubeOidcResult> {
    return pulumi.output(args).apply((a: any) => getKubeOidc(a, opts))
}

/**
 * A collection of arguments for invoking getKubeOidc.
 */
export interface GetKubeOidcOutputArgs {
    /**
     * The OIDC client ID.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The OIDC issuer url.
     */
    issuerUrl?: pulumi.Input<string>;
    /**
     * The id of the managed kubernetes cluster.
     */
    kubeId: pulumi.Input<string>;
    oidcCaContent?: pulumi.Input<string>;
    oidcGroupsClaims?: pulumi.Input<pulumi.Input<string>[]>;
    oidcGroupsPrefix?: pulumi.Input<string>;
    oidcRequiredClaims?: pulumi.Input<pulumi.Input<string>[]>;
    oidcSigningAlgs?: pulumi.Input<pulumi.Input<string>[]>;
    oidcUsernameClaim?: pulumi.Input<string>;
    oidcUsernamePrefix?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
