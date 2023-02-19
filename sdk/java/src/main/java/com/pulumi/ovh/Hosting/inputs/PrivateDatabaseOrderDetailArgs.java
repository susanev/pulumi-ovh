// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Hosting.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PrivateDatabaseOrderDetailArgs extends com.pulumi.resources.ResourceArgs {

    public static final PrivateDatabaseOrderDetailArgs Empty = new PrivateDatabaseOrderDetailArgs();

    /**
     * Custom description on your privatedatabase order.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Custom description on your privatedatabase order.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * expiration date
     * 
     */
    @Import(name="domain")
    private @Nullable Output<String> domain;

    /**
     * @return expiration date
     * 
     */
    public Optional<Output<String>> domain() {
        return Optional.ofNullable(this.domain);
    }

    /**
     * order detail id
     * 
     */
    @Import(name="orderDetailId")
    private @Nullable Output<Integer> orderDetailId;

    /**
     * @return order detail id
     * 
     */
    public Optional<Output<Integer>> orderDetailId() {
        return Optional.ofNullable(this.orderDetailId);
    }

    /**
     * quantity
     * 
     */
    @Import(name="quantity")
    private @Nullable Output<String> quantity;

    /**
     * @return quantity
     * 
     */
    public Optional<Output<String>> quantity() {
        return Optional.ofNullable(this.quantity);
    }

    private PrivateDatabaseOrderDetailArgs() {}

    private PrivateDatabaseOrderDetailArgs(PrivateDatabaseOrderDetailArgs $) {
        this.description = $.description;
        this.domain = $.domain;
        this.orderDetailId = $.orderDetailId;
        this.quantity = $.quantity;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PrivateDatabaseOrderDetailArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PrivateDatabaseOrderDetailArgs $;

        public Builder() {
            $ = new PrivateDatabaseOrderDetailArgs();
        }

        public Builder(PrivateDatabaseOrderDetailArgs defaults) {
            $ = new PrivateDatabaseOrderDetailArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Custom description on your privatedatabase order.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Custom description on your privatedatabase order.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param domain expiration date
         * 
         * @return builder
         * 
         */
        public Builder domain(@Nullable Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain expiration date
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param orderDetailId order detail id
         * 
         * @return builder
         * 
         */
        public Builder orderDetailId(@Nullable Output<Integer> orderDetailId) {
            $.orderDetailId = orderDetailId;
            return this;
        }

        /**
         * @param orderDetailId order detail id
         * 
         * @return builder
         * 
         */
        public Builder orderDetailId(Integer orderDetailId) {
            return orderDetailId(Output.of(orderDetailId));
        }

        /**
         * @param quantity quantity
         * 
         * @return builder
         * 
         */
        public Builder quantity(@Nullable Output<String> quantity) {
            $.quantity = quantity;
            return this;
        }

        /**
         * @param quantity quantity
         * 
         * @return builder
         * 
         */
        public Builder quantity(String quantity) {
            return quantity(Output.of(quantity));
        }

        public PrivateDatabaseOrderDetailArgs build() {
            return $;
        }
    }

}
