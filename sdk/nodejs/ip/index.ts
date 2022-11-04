// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetServiceArgs, GetServiceResult, GetServiceOutputArgs } from "./getService";
export const getService: typeof import("./getService").getService = null as any;
export const getServiceOutput: typeof import("./getService").getServiceOutput = null as any;
utilities.lazyLoad(exports, ["getService","getServiceOutput"], () => require("./getService"));

export { IpServiceArgs, IpServiceState } from "./ipService";
export type IpService = import("./ipService").IpService;
export const IpService: typeof import("./ipService").IpService = null as any;
utilities.lazyLoad(exports, ["IpService"], () => require("./ipService"));

export { ReverseArgs, ReverseState } from "./reverse";
export type Reverse = import("./reverse").Reverse;
export const Reverse: typeof import("./reverse").Reverse = null as any;
utilities.lazyLoad(exports, ["Reverse"], () => require("./reverse"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "ovh:Ip/ipService:IpService":
                return new IpService(name, <any>undefined, { urn })
            case "ovh:Ip/reverse:Reverse":
                return new Reverse(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("ovh", "Ip/ipService", _module)
pulumi.runtime.registerResourceModule("ovh", "Ip/reverse", _module)
