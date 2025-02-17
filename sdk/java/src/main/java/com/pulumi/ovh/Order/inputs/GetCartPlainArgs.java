// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Order.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCartPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCartPlainArgs Empty = new GetCartPlainArgs();

    /**
     * Assign a shopping cart to an loggedin client. Values can be `true` or `false`.
     * 
     */
    @Import(name="assign")
    private @Nullable Boolean assign;

    /**
     * @return Assign a shopping cart to an loggedin client. Values can be `true` or `false`.
     * 
     */
    public Optional<Boolean> assign() {
        return Optional.ofNullable(this.assign);
    }

    /**
     * Description of your cart
     * 
     */
    @Import(name="description")
    private @Nullable String description;

    /**
     * @return Description of your cart
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Expiration time (format: 2006-01-02T15:04:05+00:00)
     * 
     */
    @Import(name="expire")
    private @Nullable String expire;

    /**
     * @return Expiration time (format: 2006-01-02T15:04:05+00:00)
     * 
     */
    public Optional<String> expire() {
        return Optional.ofNullable(this.expire);
    }

    /**
     * OVHcloud Subsidiary
     * 
     */
    @Import(name="ovhSubsidiary", required=true)
    private String ovhSubsidiary;

    /**
     * @return OVHcloud Subsidiary
     * 
     */
    public String ovhSubsidiary() {
        return this.ovhSubsidiary;
    }

    private GetCartPlainArgs() {}

    private GetCartPlainArgs(GetCartPlainArgs $) {
        this.assign = $.assign;
        this.description = $.description;
        this.expire = $.expire;
        this.ovhSubsidiary = $.ovhSubsidiary;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCartPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCartPlainArgs $;

        public Builder() {
            $ = new GetCartPlainArgs();
        }

        public Builder(GetCartPlainArgs defaults) {
            $ = new GetCartPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param assign Assign a shopping cart to an loggedin client. Values can be `true` or `false`.
         * 
         * @return builder
         * 
         */
        public Builder assign(@Nullable Boolean assign) {
            $.assign = assign;
            return this;
        }

        /**
         * @param description Description of your cart
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        /**
         * @param expire Expiration time (format: 2006-01-02T15:04:05+00:00)
         * 
         * @return builder
         * 
         */
        public Builder expire(@Nullable String expire) {
            $.expire = expire;
            return this;
        }

        /**
         * @param ovhSubsidiary OVHcloud Subsidiary
         * 
         * @return builder
         * 
         */
        public Builder ovhSubsidiary(String ovhSubsidiary) {
            $.ovhSubsidiary = ovhSubsidiary;
            return this;
        }

        public GetCartPlainArgs build() {
            $.ovhSubsidiary = Objects.requireNonNull($.ovhSubsidiary, "expected parameter 'ovhSubsidiary' to be non-null");
            return $;
        }
    }

}
