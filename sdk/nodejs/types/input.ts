// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export namespace CloudProject {
    export interface ContainerRegistryPlan {
        /**
         * Plan code from the catalog
         */
        code?: pulumi.Input<string>;
        /**
         * Plan creation date
         */
        createdAt?: pulumi.Input<string>;
        /**
         * Features of the plan
         */
        features?: pulumi.Input<pulumi.Input<inputs.CloudProject.ContainerRegistryPlanFeature>[]>;
        /**
         * Plan ID
         */
        id?: pulumi.Input<string>;
        /**
         * Registry name
         */
        name?: pulumi.Input<string>;
        /**
         * Container registry limits
         */
        registryLimits?: pulumi.Input<pulumi.Input<inputs.CloudProject.ContainerRegistryPlanRegistryLimit>[]>;
        /**
         * Registry last update date
         */
        updatedAt?: pulumi.Input<string>;
    }

    export interface ContainerRegistryPlanFeature {
        /**
         * Vulnerability scanning
         */
        vulnerability?: pulumi.Input<boolean>;
    }

    export interface ContainerRegistryPlanRegistryLimit {
        /**
         * Docker image storage limits in bytes
         */
        imageStorage?: pulumi.Input<number>;
        /**
         * Parallel requests on Docker image API (/v2 Docker registry API)
         */
        parallelRequest?: pulumi.Input<number>;
    }

    export interface DatabaseEndpoint {
        /**
         * Type of component the URI relates to.
         */
        component?: pulumi.Input<string>;
        /**
         * Domain of the cluster.
         */
        domain?: pulumi.Input<string>;
        /**
         * Path of the endpoint.
         */
        path?: pulumi.Input<string>;
        /**
         * Connection port for the endpoint.
         */
        port?: pulumi.Input<number>;
        /**
         * Scheme used to generate the URI.
         */
        scheme?: pulumi.Input<string>;
        /**
         * Defines whether the endpoint uses SSL.
         */
        ssl?: pulumi.Input<boolean>;
        /**
         * SSL mode used to connect to the service if the SSL is enabled.
         */
        sslMode?: pulumi.Input<string>;
        /**
         * URI of the endpoint.
         */
        uri?: pulumi.Input<string>;
    }

    export interface DatabaseNode {
        /**
         * Private network id in which the node should be deployed. It's the regional openstackId of the private network
         */
        networkId?: pulumi.Input<string>;
        /**
         * Public cloud region in which the node should be deployed.
         * Ex: "GRA'.
         */
        region: pulumi.Input<string>;
        /**
         * Private subnet ID in which the node is.
         */
        subnetId?: pulumi.Input<string>;
    }

    export interface GetKubeCustomization {
        apiserver?: inputs.CloudProject.GetKubeCustomizationApiserver;
    }

    export interface GetKubeCustomizationArgs {
        apiserver?: pulumi.Input<inputs.CloudProject.GetKubeCustomizationApiserverArgs>;
    }

    export interface GetKubeCustomizationApiserver {
        admissionplugins?: inputs.CloudProject.GetKubeCustomizationApiserverAdmissionplugins;
    }

    export interface GetKubeCustomizationApiserverArgs {
        admissionplugins?: pulumi.Input<inputs.CloudProject.GetKubeCustomizationApiserverAdmissionpluginsArgs>;
    }

    export interface GetKubeCustomizationApiserverAdmissionplugins {
        disableds?: string[];
        enableds?: string[];
    }

    export interface GetKubeCustomizationApiserverAdmissionpluginsArgs {
        disableds?: pulumi.Input<pulumi.Input<string>[]>;
        enableds?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface KubeCustomization {
        apiserver?: pulumi.Input<inputs.CloudProject.KubeCustomizationApiserver>;
    }

    export interface KubeCustomizationApiserver {
        admissionplugins?: pulumi.Input<inputs.CloudProject.KubeCustomizationApiserverAdmissionplugins>;
    }

    export interface KubeCustomizationApiserverAdmissionplugins {
        disableds?: pulumi.Input<pulumi.Input<string>[]>;
        enableds?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface KubeNodePoolTemplate {
        metadata?: pulumi.Input<inputs.CloudProject.KubeNodePoolTemplateMetadata>;
        spec?: pulumi.Input<inputs.CloudProject.KubeNodePoolTemplateSpec>;
    }

    export interface KubeNodePoolTemplateMetadata {
        annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
        finalizers?: pulumi.Input<pulumi.Input<string>[]>;
        labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    }

    export interface KubeNodePoolTemplateSpec {
        taints?: pulumi.Input<pulumi.Input<{[key: string]: any}>[]>;
        unschedulable?: pulumi.Input<boolean>;
    }

    export interface KubePrivateNetworkConfiguration {
        defaultVrackGateway: pulumi.Input<string>;
        privateNetworkRoutingAsDefault: pulumi.Input<boolean>;
    }

    export interface NetworkPrivateRegionsAttribute {
        openstackid?: pulumi.Input<string>;
        region?: pulumi.Input<string>;
        /**
         * the status of the network. should be normally set to 'ACTIVE'.
         */
        status: pulumi.Input<string>;
    }

    export interface NetworkPrivateRegionsStatus {
        region?: pulumi.Input<string>;
        /**
         * the status of the network. should be normally set to 'ACTIVE'.
         */
        status: pulumi.Input<string>;
    }

    export interface NetworkPrivateSubnetIpPool {
        /**
         * Enable DHCP.
         * Changing this forces a new resource to be created. Defaults to false.
         * _
         */
        dhcp?: pulumi.Input<boolean>;
        /**
         * Last ip for this region.
         * Changing this value recreates the subnet.
         */
        end?: pulumi.Input<string>;
        /**
         * Global network in CIDR format.
         * Changing this value recreates the subnet
         */
        network?: pulumi.Input<string>;
        /**
         * The region in which the network subnet will be created.
         * Ex.: "GRA1". Changing this value recreates the resource.
         */
        region?: pulumi.Input<string>;
        /**
         * First ip for this region.
         * Changing this value recreates the subnet.
         */
        start?: pulumi.Input<string>;
    }

    export interface ProjectOrder {
        /**
         * date
         */
        date?: pulumi.Input<string>;
        /**
         * Information about a Bill entry
         */
        details?: pulumi.Input<pulumi.Input<inputs.CloudProject.ProjectOrderDetail>[]>;
        /**
         * expiration date
         */
        expirationDate?: pulumi.Input<string>;
        /**
         * order id
         */
        orderId?: pulumi.Input<number>;
    }

    export interface ProjectOrderDetail {
        /**
         * A description associated with the user.
         */
        description?: pulumi.Input<string>;
        /**
         * expiration date
         */
        domain?: pulumi.Input<string>;
        /**
         * order detail id
         */
        orderDetailId?: pulumi.Input<number>;
        /**
         * quantity
         */
        quantity?: pulumi.Input<string>;
    }

    export interface ProjectPlan {
        /**
         * Catalog name
         */
        catalogName?: pulumi.Input<string>;
        /**
         * Representation of a configuration item for personalizing product
         */
        configurations?: pulumi.Input<pulumi.Input<inputs.CloudProject.ProjectPlanConfiguration>[]>;
        /**
         * duration
         */
        duration: pulumi.Input<string>;
        /**
         * Plan code
         */
        planCode: pulumi.Input<string>;
        /**
         * Pricing model identifier
         */
        pricingMode: pulumi.Input<string>;
    }

    export interface ProjectPlanConfiguration {
        /**
         * Identifier of the resource
         */
        label: pulumi.Input<string>;
        /**
         * Path to the resource in API.OVH.COM
         */
        value: pulumi.Input<string>;
    }

    export interface ProjectPlanOption {
        /**
         * Catalog name
         */
        catalogName?: pulumi.Input<string>;
        /**
         * Representation of a configuration item for personalizing product
         */
        configurations?: pulumi.Input<pulumi.Input<inputs.CloudProject.ProjectPlanOptionConfiguration>[]>;
        /**
         * duration
         */
        duration: pulumi.Input<string>;
        /**
         * Plan code
         */
        planCode: pulumi.Input<string>;
        /**
         * Pricing model identifier
         */
        pricingMode: pulumi.Input<string>;
    }

    export interface ProjectPlanOptionConfiguration {
        /**
         * Identifier of the resource
         */
        label: pulumi.Input<string>;
        /**
         * Path to the resource in API.OVH.COM
         */
        value: pulumi.Input<string>;
    }

    export interface UserRole {
        /**
         * A description associated with the user.
         */
        description?: pulumi.Input<string>;
        /**
         * id of the role
         */
        id?: pulumi.Input<string>;
        /**
         * name of the role
         */
        name?: pulumi.Input<string>;
        /**
         * list of permissions associated with the role
         */
        permissions?: pulumi.Input<pulumi.Input<string>[]>;
    }
}

export namespace CloudProjectDatabase {
    export interface OpensearchUserAcl {
        /**
         * Pattern of the ACL.
         */
        pattern: pulumi.Input<string>;
        /**
         * Permission of the ACL
         * Available permission:
         * * `admin`
         * * `read`
         * * `write`
         * * `readwrite`
         * * `deny`
         */
        permission: pulumi.Input<string>;
    }
}

export namespace Dbaas {
    export interface LogsInputConfiguration {
        /**
         * Flowgger configuration
         */
        flowgger?: pulumi.Input<inputs.Dbaas.LogsInputConfigurationFlowgger>;
        /**
         * Logstash configuration
         */
        logstash?: pulumi.Input<inputs.Dbaas.LogsInputConfigurationLogstash>;
    }

    export interface LogsInputConfigurationFlowgger {
        /**
         * Type of format to decode. One of "RFC5424", "LTSV", "GELF", "CAPNP"
         */
        logFormat: pulumi.Input<string>;
        /**
         * Indicates how messages are delimited. One of "LINE", "NUL", "SYSLEN", "CAPNP"
         */
        logFraming: pulumi.Input<string>;
    }

    export interface LogsInputConfigurationLogstash {
        /**
         * The filter section of logstash.conf
         */
        filterSection?: pulumi.Input<string>;
        /**
         * The filter section of logstash.conf
         */
        inputSection: pulumi.Input<string>;
        /**
         * The list of customs Grok patterns
         */
        patternSection?: pulumi.Input<string>;
    }
}

export namespace Dedicated {
    export interface ServiceInstallTaskDetails {
        /**
         * Template change log details.
         *
         * @deprecated field is not used anymore
         */
        changeLog?: pulumi.Input<string>;
        /**
         * Set up the server using the provided hostname instead of the default hostname.
         */
        customHostname?: pulumi.Input<string>;
        /**
         * Disk group id.
         */
        diskGroupId?: pulumi.Input<number>;
        /**
         * set to true to install RTM.
         */
        installRtm?: pulumi.Input<boolean>;
        /**
         * set to true to install sql server (Windows template only).
         */
        installSqlServer?: pulumi.Input<boolean>;
        /**
         * language.
         */
        language?: pulumi.Input<string>;
        /**
         * set to true to disable RAID.
         */
        noRaid?: pulumi.Input<boolean>;
        /**
         * Indicate the URL where your postinstall customisation script is located.
         */
        postInstallationScriptLink?: pulumi.Input<string>;
        /**
         * Indicate the string returned by your postinstall customisation script on successful execution. Advice: your script should return a unique validation string in case of succes. A good example is 'loh1Xee7eo OK OK OK UGh8Ang1Gu'.
         */
        postInstallationScriptReturn?: pulumi.Input<string>;
        /**
         * set to true to make a hardware raid reset.
         */
        resetHwRaid?: pulumi.Input<boolean>;
        /**
         * soft raid devices.
         */
        softRaidDevices?: pulumi.Input<number>;
        /**
         * Name of the ssh key that should be installed. Password login will be disabled.
         */
        sshKeyName?: pulumi.Input<string>;
        /**
         * Use the distribution's native kernel instead of the recommended OVHcloud Kernel.
         */
        useDistribKernel?: pulumi.Input<boolean>;
        /**
         * set to true to use SPLA.
         */
        useSpla?: pulumi.Input<boolean>;
    }
}

export namespace Domain {
    export interface ZoneOrder {
        /**
         * date
         */
        date?: pulumi.Input<string>;
        /**
         * Information about a Bill entry
         */
        details?: pulumi.Input<pulumi.Input<inputs.Domain.ZoneOrderDetail>[]>;
        /**
         * expiration date
         */
        expirationDate?: pulumi.Input<string>;
        /**
         * order id
         */
        orderId?: pulumi.Input<number>;
    }

    export interface ZoneOrderDetail {
        /**
         * description
         */
        description?: pulumi.Input<string>;
        /**
         * expiration date
         */
        domain?: pulumi.Input<string>;
        /**
         * order detail id
         */
        orderDetailId?: pulumi.Input<number>;
        /**
         * quantity
         */
        quantity?: pulumi.Input<string>;
    }

    export interface ZonePlan {
        /**
         * Catalog name
         */
        catalogName?: pulumi.Input<string>;
        /**
         * Representation of a configuration item for personalizing product
         */
        configurations?: pulumi.Input<pulumi.Input<inputs.Domain.ZonePlanConfiguration>[]>;
        /**
         * duration
         */
        duration: pulumi.Input<string>;
        /**
         * Plan code
         */
        planCode: pulumi.Input<string>;
        /**
         * Pricing model identifier
         */
        pricingMode: pulumi.Input<string>;
    }

    export interface ZonePlanConfiguration {
        /**
         * Identifier of the resource
         */
        label: pulumi.Input<string>;
        /**
         * Path to the resource in API.OVH.COM
         */
        value: pulumi.Input<string>;
    }

    export interface ZonePlanOption {
        /**
         * Catalog name
         */
        catalogName?: pulumi.Input<string>;
        /**
         * Representation of a configuration item for personalizing product
         */
        configurations?: pulumi.Input<pulumi.Input<inputs.Domain.ZonePlanOptionConfiguration>[]>;
        /**
         * duration
         */
        duration: pulumi.Input<string>;
        /**
         * Plan code
         */
        planCode: pulumi.Input<string>;
        /**
         * Pricing model identifier
         */
        pricingMode: pulumi.Input<string>;
    }

    export interface ZonePlanOptionConfiguration {
        /**
         * Identifier of the resource
         */
        label: pulumi.Input<string>;
        /**
         * Path to the resource in API.OVH.COM
         */
        value: pulumi.Input<string>;
    }
}

export namespace Ip {
    export interface IpServiceOrder {
        /**
         * date
         */
        date?: pulumi.Input<string>;
        /**
         * Information about a Bill entry
         */
        details?: pulumi.Input<pulumi.Input<inputs.Ip.IpServiceOrderDetail>[]>;
        /**
         * expiration date
         */
        expirationDate?: pulumi.Input<string>;
        /**
         * order id
         */
        orderId?: pulumi.Input<number>;
    }

    export interface IpServiceOrderDetail {
        /**
         * Custom description on your ip.
         */
        description?: pulumi.Input<string>;
        /**
         * expiration date
         */
        domain?: pulumi.Input<string>;
        /**
         * order detail id
         */
        orderDetailId?: pulumi.Input<number>;
        /**
         * quantity
         */
        quantity?: pulumi.Input<string>;
    }

    export interface IpServicePlan {
        /**
         * Catalog name
         */
        catalogName?: pulumi.Input<string>;
        /**
         * Representation of a configuration item for personalizing product
         */
        configurations?: pulumi.Input<pulumi.Input<inputs.Ip.IpServicePlanConfiguration>[]>;
        /**
         * duration
         */
        duration: pulumi.Input<string>;
        /**
         * Plan code
         */
        planCode: pulumi.Input<string>;
        /**
         * Pricing model identifier
         */
        pricingMode: pulumi.Input<string>;
    }

    export interface IpServicePlanConfiguration {
        /**
         * Identifier of the resource
         */
        label: pulumi.Input<string>;
        /**
         * Path to the resource in API.OVH.COM
         */
        value: pulumi.Input<string>;
    }

    export interface IpServicePlanOption {
        /**
         * Catalog name
         */
        catalogName?: pulumi.Input<string>;
        /**
         * Representation of a configuration item for personalizing product
         */
        configurations?: pulumi.Input<pulumi.Input<inputs.Ip.IpServicePlanOptionConfiguration>[]>;
        /**
         * duration
         */
        duration: pulumi.Input<string>;
        /**
         * Plan code
         */
        planCode: pulumi.Input<string>;
        /**
         * Pricing model identifier
         */
        pricingMode: pulumi.Input<string>;
    }

    export interface IpServicePlanOptionConfiguration {
        /**
         * Identifier of the resource
         */
        label: pulumi.Input<string>;
        /**
         * Path to the resource in API.OVH.COM
         */
        value: pulumi.Input<string>;
    }

    export interface IpServiceRoutedTo {
        /**
         * Service where ip is routed to
         * * `serviceName`: service name
         */
        serviceName?: pulumi.Input<string>;
    }
}

export namespace IpLoadBalancing {
    export interface HttpFarmProbe {
        /**
         * Force use of SSL (TLS)
         */
        forceSsl?: pulumi.Input<boolean>;
        /**
         * probe interval, Value between 30 and 3600 seconds, default 30
         */
        interval?: pulumi.Input<number>;
        /**
         * What to match `pattern` against (`contains`, `default`, `internal`, `matches`, `status`)
         */
        match?: pulumi.Input<string>;
        /**
         * HTTP probe method (`GET`, `HEAD`, `OPTIONS`, `internal`)
         */
        method?: pulumi.Input<string>;
        /**
         * Negate probe result
         */
        negate?: pulumi.Input<boolean>;
        /**
         * Pattern to match against `match`
         */
        pattern?: pulumi.Input<string>;
        /**
         * Port for backends to receive traffic on.
         */
        port?: pulumi.Input<number>;
        /**
         * Valid values : `http`, `internal`, `mysql`, `oco`, `pgsql`, `smtp`, `tcp`
         */
        type: pulumi.Input<string>;
        /**
         * URL for HTTP probe type.
         */
        url?: pulumi.Input<string>;
    }

    export interface HttpRouteAction {
        /**
         * HTTP status code for "redirect" and "reject" actions
         */
        status?: pulumi.Input<number>;
        /**
         * Farm ID for "farm" action type or URL template for "redirect" action. You may use ${uri}, ${protocol}, ${host}, ${port} and ${path} variables in redirect target
         */
        target?: pulumi.Input<string>;
        /**
         * Action to trigger if all the rules of this route matches
         */
        type: pulumi.Input<string>;
    }

    export interface HttpRouteRule {
        /**
         * Name of the field to match like "protocol" or "host" "/ipLoadbalancing/{serviceName}/route/availableRules" for a list of available rules
         */
        field?: pulumi.Input<string>;
        /**
         * Matching operator. Not all operators are available for all fields. See "availableRules"
         * * `negate`- Invert the matching operator effect
         */
        match?: pulumi.Input<string>;
        negate?: pulumi.Input<boolean>;
        /**
         * Value to match against this match. Interpretation if this field depends on the match and field
         */
        pattern?: pulumi.Input<string>;
        /**
         * Id of your rule
         */
        ruleId?: pulumi.Input<number>;
        /**
         * Name of sub-field, if applicable. This may be a Cookie or Header name for instance
         */
        subField?: pulumi.Input<string>;
    }

    export interface LoadBalancerOrder {
        /**
         * date
         */
        date?: pulumi.Input<string>;
        /**
         * Information about a Bill entry
         */
        details?: pulumi.Input<pulumi.Input<inputs.IpLoadBalancing.LoadBalancerOrderDetail>[]>;
        /**
         * expiration date
         */
        expirationDate?: pulumi.Input<string>;
        /**
         * order id
         */
        orderId?: pulumi.Input<number>;
    }

    export interface LoadBalancerOrderDetail {
        /**
         * description
         */
        description?: pulumi.Input<string>;
        /**
         * expiration date
         */
        domain?: pulumi.Input<string>;
        /**
         * order detail id
         */
        orderDetailId?: pulumi.Input<number>;
        /**
         * quantity
         */
        quantity?: pulumi.Input<string>;
    }

    export interface LoadBalancerOrderableZone {
        /**
         * The zone three letter code
         */
        name?: pulumi.Input<string>;
        /**
         * Plan code
         */
        planCode?: pulumi.Input<string>;
    }

    export interface LoadBalancerPlan {
        /**
         * Catalog name
         */
        catalogName?: pulumi.Input<string>;
        /**
         * Representation of a configuration item for personalizing product
         */
        configurations?: pulumi.Input<pulumi.Input<inputs.IpLoadBalancing.LoadBalancerPlanConfiguration>[]>;
        /**
         * duration
         */
        duration: pulumi.Input<string>;
        /**
         * Plan code
         */
        planCode: pulumi.Input<string>;
        /**
         * Pricing model identifier
         */
        pricingMode: pulumi.Input<string>;
    }

    export interface LoadBalancerPlanConfiguration {
        /**
         * Identifier of the resource
         */
        label: pulumi.Input<string>;
        /**
         * Path to the resource in API.OVH.COM
         */
        value: pulumi.Input<string>;
    }

    export interface LoadBalancerPlanOption {
        /**
         * Catalog name
         */
        catalogName?: pulumi.Input<string>;
        /**
         * Representation of a configuration item for personalizing product
         */
        configurations?: pulumi.Input<pulumi.Input<inputs.IpLoadBalancing.LoadBalancerPlanOptionConfiguration>[]>;
        /**
         * duration
         */
        duration: pulumi.Input<string>;
        /**
         * Plan code
         */
        planCode: pulumi.Input<string>;
        /**
         * Pricing model identifier
         */
        pricingMode: pulumi.Input<string>;
    }

    export interface LoadBalancerPlanOptionConfiguration {
        /**
         * Identifier of the resource
         */
        label: pulumi.Input<string>;
        /**
         * Path to the resource in API.OVH.COM
         */
        value: pulumi.Input<string>;
    }

    export interface TcpFarmProbe {
        /**
         * Force use of SSL (TLS)
         */
        forceSsl?: pulumi.Input<boolean>;
        /**
         * probe interval, Value between 30 and 3600 seconds, default 30
         */
        interval?: pulumi.Input<number>;
        /**
         * What to match `pattern` against (`contains`, `default`, `internal`, `matches`, `status`)
         */
        match?: pulumi.Input<string>;
        /**
         * HTTP probe method (`GET`, `HEAD`, `OPTIONS`, `internal`)
         */
        method?: pulumi.Input<string>;
        /**
         * Negate probe result
         */
        negate?: pulumi.Input<boolean>;
        /**
         * Pattern to match against `match`
         */
        pattern?: pulumi.Input<string>;
        /**
         * Port for backends to receive traffic on.
         */
        port?: pulumi.Input<number>;
        /**
         * Valid values : `http`, `internal`, `mysql`, `oco`, `pgsql`, `smtp`, `tcp`
         */
        type: pulumi.Input<string>;
        /**
         * URL for HTTP probe type.
         */
        url?: pulumi.Input<string>;
    }

    export interface TcpRouteAction {
        /**
         * Farm ID for "farm" action type, empty for others.
         */
        target?: pulumi.Input<string>;
        /**
         * Action to trigger if all the rules of this route matches
         */
        type: pulumi.Input<string>;
    }

    export interface TcpRouteRule {
        /**
         * Name of the field to match like "protocol" or "host" "/ipLoadbalancing/{serviceName}/route/availableRules" for a list of available rules
         */
        field?: pulumi.Input<string>;
        /**
         * Matching operator. Not all operators are available for all fields. See "availableRules"
         * * `negate`- Invert the matching operator effect
         */
        match?: pulumi.Input<string>;
        negate?: pulumi.Input<boolean>;
        /**
         * Value to match against this match. Interpretation if this field depends on the match and field
         */
        pattern?: pulumi.Input<string>;
        /**
         * Id of your rule
         */
        ruleId?: pulumi.Input<number>;
        /**
         * Name of sub-field, if applicable. This may be a Cookie or Header name for instance
         */
        subField?: pulumi.Input<string>;
    }
}

export namespace Me {
    export interface InstallationTemplateCustomization {
        /**
         * @deprecated field is not used anymore
         */
        changeLog?: pulumi.Input<string>;
        customHostname?: pulumi.Input<string>;
        postInstallationScriptLink?: pulumi.Input<string>;
        postInstallationScriptReturn?: pulumi.Input<string>;
        /**
         * @deprecated field is not used anymore
         */
        rating?: pulumi.Input<number>;
        sshKeyName?: pulumi.Input<string>;
        useDistributionKernel?: pulumi.Input<boolean>;
    }
}

export namespace Order {
}

export namespace Vrack {
    export interface VrackOrder {
        /**
         * date
         */
        date?: pulumi.Input<string>;
        /**
         * Information about a Bill entry
         */
        details?: pulumi.Input<pulumi.Input<inputs.Vrack.VrackOrderDetail>[]>;
        /**
         * expiration date
         */
        expirationDate?: pulumi.Input<string>;
        /**
         * order id
         */
        orderId?: pulumi.Input<number>;
    }

    export interface VrackOrderDetail {
        /**
         * yourvrackdescription
         */
        description?: pulumi.Input<string>;
        /**
         * expiration date
         */
        domain?: pulumi.Input<string>;
        /**
         * order detail id
         */
        orderDetailId?: pulumi.Input<number>;
        /**
         * quantity
         */
        quantity?: pulumi.Input<string>;
    }

    export interface VrackPlan {
        /**
         * Catalog name
         */
        catalogName?: pulumi.Input<string>;
        /**
         * Representation of a configuration item for personalizing product
         */
        configurations?: pulumi.Input<pulumi.Input<inputs.Vrack.VrackPlanConfiguration>[]>;
        /**
         * duration
         */
        duration: pulumi.Input<string>;
        /**
         * Plan code
         */
        planCode: pulumi.Input<string>;
        /**
         * Pricing model identifier
         */
        pricingMode: pulumi.Input<string>;
    }

    export interface VrackPlanConfiguration {
        /**
         * Identifier of the resource
         */
        label: pulumi.Input<string>;
        /**
         * Path to the resource in API.OVH.COM
         */
        value: pulumi.Input<string>;
    }

    export interface VrackPlanOption {
        /**
         * Catalog name
         */
        catalogName?: pulumi.Input<string>;
        /**
         * Representation of a configuration item for personalizing product
         */
        configurations?: pulumi.Input<pulumi.Input<inputs.Vrack.VrackPlanOptionConfiguration>[]>;
        /**
         * duration
         */
        duration: pulumi.Input<string>;
        /**
         * Plan code
         */
        planCode: pulumi.Input<string>;
        /**
         * Pricing model identifier
         */
        pricingMode: pulumi.Input<string>;
    }

    export interface VrackPlanOptionConfiguration {
        /**
         * Identifier of the resource
         */
        label: pulumi.Input<string>;
        /**
         * Path to the resource in API.OVH.COM
         */
        value: pulumi.Input<string>;
    }
}
