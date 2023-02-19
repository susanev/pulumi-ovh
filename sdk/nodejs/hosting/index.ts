// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetPrivateDatabaseArgs, GetPrivateDatabaseResult, GetPrivateDatabaseOutputArgs } from "./getPrivateDatabase";
export const getPrivateDatabase: typeof import("./getPrivateDatabase").getPrivateDatabase = null as any;
export const getPrivateDatabaseOutput: typeof import("./getPrivateDatabase").getPrivateDatabaseOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateDatabase","getPrivateDatabaseOutput"], () => require("./getPrivateDatabase"));

export { GetPrivateDatabaseAllowlistArgs, GetPrivateDatabaseAllowlistResult, GetPrivateDatabaseAllowlistOutputArgs } from "./getPrivateDatabaseAllowlist";
export const getPrivateDatabaseAllowlist: typeof import("./getPrivateDatabaseAllowlist").getPrivateDatabaseAllowlist = null as any;
export const getPrivateDatabaseAllowlistOutput: typeof import("./getPrivateDatabaseAllowlist").getPrivateDatabaseAllowlistOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateDatabaseAllowlist","getPrivateDatabaseAllowlistOutput"], () => require("./getPrivateDatabaseAllowlist"));

export { GetPrivateDatabaseDbArgs, GetPrivateDatabaseDbResult, GetPrivateDatabaseDbOutputArgs } from "./getPrivateDatabaseDb";
export const getPrivateDatabaseDb: typeof import("./getPrivateDatabaseDb").getPrivateDatabaseDb = null as any;
export const getPrivateDatabaseDbOutput: typeof import("./getPrivateDatabaseDb").getPrivateDatabaseDbOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateDatabaseDb","getPrivateDatabaseDbOutput"], () => require("./getPrivateDatabaseDb"));

export { GetPrivateDatabaseUserArgs, GetPrivateDatabaseUserResult, GetPrivateDatabaseUserOutputArgs } from "./getPrivateDatabaseUser";
export const getPrivateDatabaseUser: typeof import("./getPrivateDatabaseUser").getPrivateDatabaseUser = null as any;
export const getPrivateDatabaseUserOutput: typeof import("./getPrivateDatabaseUser").getPrivateDatabaseUserOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateDatabaseUser","getPrivateDatabaseUserOutput"], () => require("./getPrivateDatabaseUser"));

export { GetPrivateDatabaseUserGrantArgs, GetPrivateDatabaseUserGrantResult, GetPrivateDatabaseUserGrantOutputArgs } from "./getPrivateDatabaseUserGrant";
export const getPrivateDatabaseUserGrant: typeof import("./getPrivateDatabaseUserGrant").getPrivateDatabaseUserGrant = null as any;
export const getPrivateDatabaseUserGrantOutput: typeof import("./getPrivateDatabaseUserGrant").getPrivateDatabaseUserGrantOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateDatabaseUserGrant","getPrivateDatabaseUserGrantOutput"], () => require("./getPrivateDatabaseUserGrant"));

export { PrivateDatabaseArgs, PrivateDatabaseState } from "./privateDatabase";
export type PrivateDatabase = import("./privateDatabase").PrivateDatabase;
export const PrivateDatabase: typeof import("./privateDatabase").PrivateDatabase = null as any;
utilities.lazyLoad(exports, ["PrivateDatabase"], () => require("./privateDatabase"));

export { PrivateDatabaseAllowlistArgs, PrivateDatabaseAllowlistState } from "./privateDatabaseAllowlist";
export type PrivateDatabaseAllowlist = import("./privateDatabaseAllowlist").PrivateDatabaseAllowlist;
export const PrivateDatabaseAllowlist: typeof import("./privateDatabaseAllowlist").PrivateDatabaseAllowlist = null as any;
utilities.lazyLoad(exports, ["PrivateDatabaseAllowlist"], () => require("./privateDatabaseAllowlist"));

export { PrivateDatabaseDbArgs, PrivateDatabaseDbState } from "./privateDatabaseDb";
export type PrivateDatabaseDb = import("./privateDatabaseDb").PrivateDatabaseDb;
export const PrivateDatabaseDb: typeof import("./privateDatabaseDb").PrivateDatabaseDb = null as any;
utilities.lazyLoad(exports, ["PrivateDatabaseDb"], () => require("./privateDatabaseDb"));

export { PrivateDatabaseUserArgs, PrivateDatabaseUserState } from "./privateDatabaseUser";
export type PrivateDatabaseUser = import("./privateDatabaseUser").PrivateDatabaseUser;
export const PrivateDatabaseUser: typeof import("./privateDatabaseUser").PrivateDatabaseUser = null as any;
utilities.lazyLoad(exports, ["PrivateDatabaseUser"], () => require("./privateDatabaseUser"));

export { PrivateDatabaseUserGrantArgs, PrivateDatabaseUserGrantState } from "./privateDatabaseUserGrant";
export type PrivateDatabaseUserGrant = import("./privateDatabaseUserGrant").PrivateDatabaseUserGrant;
export const PrivateDatabaseUserGrant: typeof import("./privateDatabaseUserGrant").PrivateDatabaseUserGrant = null as any;
utilities.lazyLoad(exports, ["PrivateDatabaseUserGrant"], () => require("./privateDatabaseUserGrant"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "ovh:Hosting/privateDatabase:PrivateDatabase":
                return new PrivateDatabase(name, <any>undefined, { urn })
            case "ovh:Hosting/privateDatabaseAllowlist:PrivateDatabaseAllowlist":
                return new PrivateDatabaseAllowlist(name, <any>undefined, { urn })
            case "ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb":
                return new PrivateDatabaseDb(name, <any>undefined, { urn })
            case "ovh:Hosting/privateDatabaseUser:PrivateDatabaseUser":
                return new PrivateDatabaseUser(name, <any>undefined, { urn })
            case "ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant":
                return new PrivateDatabaseUserGrant(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("ovh", "Hosting/privateDatabase", _module)
pulumi.runtime.registerResourceModule("ovh", "Hosting/privateDatabaseAllowlist", _module)
pulumi.runtime.registerResourceModule("ovh", "Hosting/privateDatabaseDb", _module)
pulumi.runtime.registerResourceModule("ovh", "Hosting/privateDatabaseUser", _module)
pulumi.runtime.registerResourceModule("ovh", "Hosting/privateDatabaseUserGrant", _module)
