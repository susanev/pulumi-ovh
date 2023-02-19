// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Hosting.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetPrivateDatabaseUserPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPrivateDatabaseUserPlainArgs Empty = new GetPrivateDatabaseUserPlainArgs();

    /**
     * The internal name of your private database
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The internal name of your private database
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    /**
     * User name
     * 
     */
    @Import(name="userName", required=true)
    private String userName;

    /**
     * @return User name
     * 
     */
    public String userName() {
        return this.userName;
    }

    private GetPrivateDatabaseUserPlainArgs() {}

    private GetPrivateDatabaseUserPlainArgs(GetPrivateDatabaseUserPlainArgs $) {
        this.serviceName = $.serviceName;
        this.userName = $.userName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPrivateDatabaseUserPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPrivateDatabaseUserPlainArgs $;

        public Builder() {
            $ = new GetPrivateDatabaseUserPlainArgs();
        }

        public Builder(GetPrivateDatabaseUserPlainArgs defaults) {
            $ = new GetPrivateDatabaseUserPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName The internal name of your private database
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param userName User name
         * 
         * @return builder
         * 
         */
        public Builder userName(String userName) {
            $.userName = userName;
            return this;
        }

        public GetPrivateDatabaseUserPlainArgs build() {
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            $.userName = Objects.requireNonNull($.userName, "expected parameter 'userName' to be non-null");
            return $;
        }
    }

}
