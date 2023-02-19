// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Dbaas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetLogsOutputGraylogStreamArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLogsOutputGraylogStreamArgs Empty = new GetLogsOutputGraylogStreamArgs();

    /**
     * The service name. It&#39;s the ID of your Logs Data Platform instance.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The service name. It&#39;s the ID of your Logs Data Platform instance.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     * Stream description
     * 
     */
    @Import(name="title", required=true)
    private Output<String> title;

    /**
     * @return Stream description
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    private GetLogsOutputGraylogStreamArgs() {}

    private GetLogsOutputGraylogStreamArgs(GetLogsOutputGraylogStreamArgs $) {
        this.serviceName = $.serviceName;
        this.title = $.title;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLogsOutputGraylogStreamArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLogsOutputGraylogStreamArgs $;

        public Builder() {
            $ = new GetLogsOutputGraylogStreamArgs();
        }

        public Builder(GetLogsOutputGraylogStreamArgs defaults) {
            $ = new GetLogsOutputGraylogStreamArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName The service name. It&#39;s the ID of your Logs Data Platform instance.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The service name. It&#39;s the ID of your Logs Data Platform instance.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param title Stream description
         * 
         * @return builder
         * 
         */
        public Builder title(Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title Stream description
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        public GetLogsOutputGraylogStreamArgs build() {
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            $.title = Objects.requireNonNull($.title, "expected parameter 'title' to be non-null");
            return $;
        }
    }

}
