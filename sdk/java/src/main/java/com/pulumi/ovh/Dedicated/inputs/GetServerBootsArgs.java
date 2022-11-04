// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Dedicated.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetServerBootsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServerBootsArgs Empty = new GetServerBootsArgs();

    /**
     * Filter the value of bootType property (harddisk, rescue, ipxeCustomerScript, internal, network)
     * 
     */
    @Import(name="bootType")
    private @Nullable Output<String> bootType;

    /**
     * @return Filter the value of bootType property (harddisk, rescue, ipxeCustomerScript, internal, network)
     * 
     */
    public Optional<Output<String>> bootType() {
        return Optional.ofNullable(this.bootType);
    }

    @Import(name="kernel")
    private @Nullable Output<String> kernel;

    public Optional<Output<String>> kernel() {
        return Optional.ofNullable(this.kernel);
    }

    /**
     * The internal name of your dedicated server.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The internal name of your dedicated server.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private GetServerBootsArgs() {}

    private GetServerBootsArgs(GetServerBootsArgs $) {
        this.bootType = $.bootType;
        this.kernel = $.kernel;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServerBootsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServerBootsArgs $;

        public Builder() {
            $ = new GetServerBootsArgs();
        }

        public Builder(GetServerBootsArgs defaults) {
            $ = new GetServerBootsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bootType Filter the value of bootType property (harddisk, rescue, ipxeCustomerScript, internal, network)
         * 
         * @return builder
         * 
         */
        public Builder bootType(@Nullable Output<String> bootType) {
            $.bootType = bootType;
            return this;
        }

        /**
         * @param bootType Filter the value of bootType property (harddisk, rescue, ipxeCustomerScript, internal, network)
         * 
         * @return builder
         * 
         */
        public Builder bootType(String bootType) {
            return bootType(Output.of(bootType));
        }

        public Builder kernel(@Nullable Output<String> kernel) {
            $.kernel = kernel;
            return this;
        }

        public Builder kernel(String kernel) {
            return kernel(Output.of(kernel));
        }

        /**
         * @param serviceName The internal name of your dedicated server.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The internal name of your dedicated server.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public GetServerBootsArgs build() {
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}
