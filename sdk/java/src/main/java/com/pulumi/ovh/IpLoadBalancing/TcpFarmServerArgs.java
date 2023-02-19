// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.IpLoadBalancing;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TcpFarmServerArgs extends com.pulumi.resources.ResourceArgs {

    public static final TcpFarmServerArgs Empty = new TcpFarmServerArgs();

    /**
     * Address of the backend server (IP from either internal or OVHcloud network)
     * 
     */
    @Import(name="address", required=true)
    private Output<String> address;

    /**
     * @return Address of the backend server (IP from either internal or OVHcloud network)
     * 
     */
    public Output<String> address() {
        return this.address;
    }

    /**
     * is it a backup server used in case of failure of all the non-backup backends
     * 
     */
    @Import(name="backup")
    private @Nullable Output<Boolean> backup;

    /**
     * @return is it a backup server used in case of failure of all the non-backup backends
     * 
     */
    public Optional<Output<Boolean>> backup() {
        return Optional.ofNullable(this.backup);
    }

    @Import(name="chain")
    private @Nullable Output<String> chain;

    public Optional<Output<String>> chain() {
        return Optional.ofNullable(this.chain);
    }

    /**
     * Label for the server
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Label for the server
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * ID of the farm this server is attached to
     * 
     */
    @Import(name="farmId", required=true)
    private Output<Integer> farmId;

    /**
     * @return ID of the farm this server is attached to
     * 
     */
    public Output<Integer> farmId() {
        return this.farmId;
    }

    /**
     * enable action when backend marked down. (`shutdown-sessions`)
     * 
     */
    @Import(name="onMarkedDown")
    private @Nullable Output<String> onMarkedDown;

    /**
     * @return enable action when backend marked down. (`shutdown-sessions`)
     * 
     */
    public Optional<Output<String>> onMarkedDown() {
        return Optional.ofNullable(this.onMarkedDown);
    }

    /**
     * Port that backend will respond on
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return Port that backend will respond on
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * defines if backend will be probed to determine health and keep as active in farm if healthy
     * 
     */
    @Import(name="probe")
    private @Nullable Output<Boolean> probe;

    /**
     * @return defines if backend will be probed to determine health and keep as active in farm if healthy
     * 
     */
    public Optional<Output<Boolean>> probe() {
        return Optional.ofNullable(this.probe);
    }

    /**
     * version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
     * 
     */
    @Import(name="proxyProtocolVersion")
    private @Nullable Output<String> proxyProtocolVersion;

    /**
     * @return version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
     * 
     */
    public Optional<Output<String>> proxyProtocolVersion() {
        return Optional.ofNullable(this.proxyProtocolVersion);
    }

    /**
     * The internal name of your IP load balancing
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The internal name of your IP load balancing
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     * is the connection ciphered with SSL (TLS)
     * 
     */
    @Import(name="ssl")
    private @Nullable Output<Boolean> ssl;

    /**
     * @return is the connection ciphered with SSL (TLS)
     * 
     */
    public Optional<Output<Boolean>> ssl() {
        return Optional.ofNullable(this.ssl);
    }

    /**
     * backend status - `active` or `inactive`
     * 
     */
    @Import(name="status", required=true)
    private Output<String> status;

    /**
     * @return backend status - `active` or `inactive`
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     * used in loadbalancing algorithm
     * 
     */
    @Import(name="weight")
    private @Nullable Output<Integer> weight;

    /**
     * @return used in loadbalancing algorithm
     * 
     */
    public Optional<Output<Integer>> weight() {
        return Optional.ofNullable(this.weight);
    }

    private TcpFarmServerArgs() {}

    private TcpFarmServerArgs(TcpFarmServerArgs $) {
        this.address = $.address;
        this.backup = $.backup;
        this.chain = $.chain;
        this.displayName = $.displayName;
        this.farmId = $.farmId;
        this.onMarkedDown = $.onMarkedDown;
        this.port = $.port;
        this.probe = $.probe;
        this.proxyProtocolVersion = $.proxyProtocolVersion;
        this.serviceName = $.serviceName;
        this.ssl = $.ssl;
        this.status = $.status;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TcpFarmServerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TcpFarmServerArgs $;

        public Builder() {
            $ = new TcpFarmServerArgs();
        }

        public Builder(TcpFarmServerArgs defaults) {
            $ = new TcpFarmServerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address Address of the backend server (IP from either internal or OVHcloud network)
         * 
         * @return builder
         * 
         */
        public Builder address(Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address Address of the backend server (IP from either internal or OVHcloud network)
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param backup is it a backup server used in case of failure of all the non-backup backends
         * 
         * @return builder
         * 
         */
        public Builder backup(@Nullable Output<Boolean> backup) {
            $.backup = backup;
            return this;
        }

        /**
         * @param backup is it a backup server used in case of failure of all the non-backup backends
         * 
         * @return builder
         * 
         */
        public Builder backup(Boolean backup) {
            return backup(Output.of(backup));
        }

        public Builder chain(@Nullable Output<String> chain) {
            $.chain = chain;
            return this;
        }

        public Builder chain(String chain) {
            return chain(Output.of(chain));
        }

        /**
         * @param displayName Label for the server
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Label for the server
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param farmId ID of the farm this server is attached to
         * 
         * @return builder
         * 
         */
        public Builder farmId(Output<Integer> farmId) {
            $.farmId = farmId;
            return this;
        }

        /**
         * @param farmId ID of the farm this server is attached to
         * 
         * @return builder
         * 
         */
        public Builder farmId(Integer farmId) {
            return farmId(Output.of(farmId));
        }

        /**
         * @param onMarkedDown enable action when backend marked down. (`shutdown-sessions`)
         * 
         * @return builder
         * 
         */
        public Builder onMarkedDown(@Nullable Output<String> onMarkedDown) {
            $.onMarkedDown = onMarkedDown;
            return this;
        }

        /**
         * @param onMarkedDown enable action when backend marked down. (`shutdown-sessions`)
         * 
         * @return builder
         * 
         */
        public Builder onMarkedDown(String onMarkedDown) {
            return onMarkedDown(Output.of(onMarkedDown));
        }

        /**
         * @param port Port that backend will respond on
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Port that backend will respond on
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param probe defines if backend will be probed to determine health and keep as active in farm if healthy
         * 
         * @return builder
         * 
         */
        public Builder probe(@Nullable Output<Boolean> probe) {
            $.probe = probe;
            return this;
        }

        /**
         * @param probe defines if backend will be probed to determine health and keep as active in farm if healthy
         * 
         * @return builder
         * 
         */
        public Builder probe(Boolean probe) {
            return probe(Output.of(probe));
        }

        /**
         * @param proxyProtocolVersion version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
         * 
         * @return builder
         * 
         */
        public Builder proxyProtocolVersion(@Nullable Output<String> proxyProtocolVersion) {
            $.proxyProtocolVersion = proxyProtocolVersion;
            return this;
        }

        /**
         * @param proxyProtocolVersion version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
         * 
         * @return builder
         * 
         */
        public Builder proxyProtocolVersion(String proxyProtocolVersion) {
            return proxyProtocolVersion(Output.of(proxyProtocolVersion));
        }

        /**
         * @param serviceName The internal name of your IP load balancing
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The internal name of your IP load balancing
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param ssl is the connection ciphered with SSL (TLS)
         * 
         * @return builder
         * 
         */
        public Builder ssl(@Nullable Output<Boolean> ssl) {
            $.ssl = ssl;
            return this;
        }

        /**
         * @param ssl is the connection ciphered with SSL (TLS)
         * 
         * @return builder
         * 
         */
        public Builder ssl(Boolean ssl) {
            return ssl(Output.of(ssl));
        }

        /**
         * @param status backend status - `active` or `inactive`
         * 
         * @return builder
         * 
         */
        public Builder status(Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status backend status - `active` or `inactive`
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param weight used in loadbalancing algorithm
         * 
         * @return builder
         * 
         */
        public Builder weight(@Nullable Output<Integer> weight) {
            $.weight = weight;
            return this;
        }

        /**
         * @param weight used in loadbalancing algorithm
         * 
         * @return builder
         * 
         */
        public Builder weight(Integer weight) {
            return weight(Output.of(weight));
        }

        public TcpFarmServerArgs build() {
            $.address = Objects.requireNonNull($.address, "expected parameter 'address' to be non-null");
            $.farmId = Objects.requireNonNull($.farmId, "expected parameter 'farmId' to be non-null");
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            $.status = Objects.requireNonNull($.status, "expected parameter 'status' to be non-null");
            return $;
        }
    }

}
