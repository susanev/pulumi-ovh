// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.IpLoadBalancing.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HttpFrontendState extends com.pulumi.resources.ResourceArgs {

    public static final HttpFrontendState Empty = new HttpFrontendState();

    /**
     * Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
     * 
     */
    @Import(name="allowedSources")
    private @Nullable Output<List<String>> allowedSources;

    /**
     * @return Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
     * 
     */
    public Optional<Output<List<String>>> allowedSources() {
        return Optional.ofNullable(this.allowedSources);
    }

    /**
     * Only attach frontend on these ip. No restriction if null. List of Ip blocks.
     * 
     */
    @Import(name="dedicatedIpfos")
    private @Nullable Output<List<String>> dedicatedIpfos;

    /**
     * @return Only attach frontend on these ip. No restriction if null. List of Ip blocks.
     * 
     */
    public Optional<Output<List<String>>> dedicatedIpfos() {
        return Optional.ofNullable(this.dedicatedIpfos);
    }

    /**
     * Default TCP Farm of your frontend
     * 
     */
    @Import(name="defaultFarmId")
    private @Nullable Output<Integer> defaultFarmId;

    /**
     * @return Default TCP Farm of your frontend
     * 
     */
    public Optional<Output<Integer>> defaultFarmId() {
        return Optional.ofNullable(this.defaultFarmId);
    }

    /**
     * Default ssl served to your customer
     * 
     */
    @Import(name="defaultSslId")
    private @Nullable Output<Integer> defaultSslId;

    /**
     * @return Default ssl served to your customer
     * 
     */
    public Optional<Output<Integer>> defaultSslId() {
        return Optional.ofNullable(this.defaultSslId);
    }

    /**
     * Disable your frontend. Default: &#39;false&#39;
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return Disable your frontend. Default: &#39;false&#39;
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * Human readable name for your frontend, this field is for you
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Human readable name for your frontend, this field is for you
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * HTTP Strict Transport Security. Default: &#39;false&#39;
     * 
     */
    @Import(name="hsts")
    private @Nullable Output<Boolean> hsts;

    /**
     * @return HTTP Strict Transport Security. Default: &#39;false&#39;
     * 
     */
    public Optional<Output<Boolean>> hsts() {
        return Optional.ofNullable(this.hsts);
    }

    /**
     * HTTP headers to add to the frontend. List of string.
     * 
     */
    @Import(name="httpHeaders")
    private @Nullable Output<List<String>> httpHeaders;

    /**
     * @return HTTP headers to add to the frontend. List of string.
     * 
     */
    public Optional<Output<List<String>>> httpHeaders() {
        return Optional.ofNullable(this.httpHeaders);
    }

    /**
     * Port(s) attached to your frontend. Supports single port (numerical value),
     * range (2 dash-delimited increasing ports) and comma-separated list of &#39;single port&#39;
     * and/or &#39;range&#39;. Each port must be in the [1;49151] range
     * 
     */
    @Import(name="port")
    private @Nullable Output<String> port;

    /**
     * @return Port(s) attached to your frontend. Supports single port (numerical value),
     * range (2 dash-delimited increasing ports) and comma-separated list of &#39;single port&#39;
     * and/or &#39;range&#39;. Each port must be in the [1;49151] range
     * 
     */
    public Optional<Output<String>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * Redirection HTTP&#39;
     * 
     */
    @Import(name="redirectLocation")
    private @Nullable Output<String> redirectLocation;

    /**
     * @return Redirection HTTP&#39;
     * 
     */
    public Optional<Output<String>> redirectLocation() {
        return Optional.ofNullable(this.redirectLocation);
    }

    /**
     * The internal name of your IP load balancing
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The internal name of your IP load balancing
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * SSL deciphering. Default: &#39;false&#39;
     * 
     */
    @Import(name="ssl")
    private @Nullable Output<Boolean> ssl;

    /**
     * @return SSL deciphering. Default: &#39;false&#39;
     * 
     */
    public Optional<Output<Boolean>> ssl() {
        return Optional.ofNullable(this.ssl);
    }

    /**
     * Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private HttpFrontendState() {}

    private HttpFrontendState(HttpFrontendState $) {
        this.allowedSources = $.allowedSources;
        this.dedicatedIpfos = $.dedicatedIpfos;
        this.defaultFarmId = $.defaultFarmId;
        this.defaultSslId = $.defaultSslId;
        this.disabled = $.disabled;
        this.displayName = $.displayName;
        this.hsts = $.hsts;
        this.httpHeaders = $.httpHeaders;
        this.port = $.port;
        this.redirectLocation = $.redirectLocation;
        this.serviceName = $.serviceName;
        this.ssl = $.ssl;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HttpFrontendState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HttpFrontendState $;

        public Builder() {
            $ = new HttpFrontendState();
        }

        public Builder(HttpFrontendState defaults) {
            $ = new HttpFrontendState(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowedSources Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
         * 
         * @return builder
         * 
         */
        public Builder allowedSources(@Nullable Output<List<String>> allowedSources) {
            $.allowedSources = allowedSources;
            return this;
        }

        /**
         * @param allowedSources Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
         * 
         * @return builder
         * 
         */
        public Builder allowedSources(List<String> allowedSources) {
            return allowedSources(Output.of(allowedSources));
        }

        /**
         * @param allowedSources Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.
         * 
         * @return builder
         * 
         */
        public Builder allowedSources(String... allowedSources) {
            return allowedSources(List.of(allowedSources));
        }

        /**
         * @param dedicatedIpfos Only attach frontend on these ip. No restriction if null. List of Ip blocks.
         * 
         * @return builder
         * 
         */
        public Builder dedicatedIpfos(@Nullable Output<List<String>> dedicatedIpfos) {
            $.dedicatedIpfos = dedicatedIpfos;
            return this;
        }

        /**
         * @param dedicatedIpfos Only attach frontend on these ip. No restriction if null. List of Ip blocks.
         * 
         * @return builder
         * 
         */
        public Builder dedicatedIpfos(List<String> dedicatedIpfos) {
            return dedicatedIpfos(Output.of(dedicatedIpfos));
        }

        /**
         * @param dedicatedIpfos Only attach frontend on these ip. No restriction if null. List of Ip blocks.
         * 
         * @return builder
         * 
         */
        public Builder dedicatedIpfos(String... dedicatedIpfos) {
            return dedicatedIpfos(List.of(dedicatedIpfos));
        }

        /**
         * @param defaultFarmId Default TCP Farm of your frontend
         * 
         * @return builder
         * 
         */
        public Builder defaultFarmId(@Nullable Output<Integer> defaultFarmId) {
            $.defaultFarmId = defaultFarmId;
            return this;
        }

        /**
         * @param defaultFarmId Default TCP Farm of your frontend
         * 
         * @return builder
         * 
         */
        public Builder defaultFarmId(Integer defaultFarmId) {
            return defaultFarmId(Output.of(defaultFarmId));
        }

        /**
         * @param defaultSslId Default ssl served to your customer
         * 
         * @return builder
         * 
         */
        public Builder defaultSslId(@Nullable Output<Integer> defaultSslId) {
            $.defaultSslId = defaultSslId;
            return this;
        }

        /**
         * @param defaultSslId Default ssl served to your customer
         * 
         * @return builder
         * 
         */
        public Builder defaultSslId(Integer defaultSslId) {
            return defaultSslId(Output.of(defaultSslId));
        }

        /**
         * @param disabled Disable your frontend. Default: &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled Disable your frontend. Default: &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param displayName Human readable name for your frontend, this field is for you
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Human readable name for your frontend, this field is for you
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param hsts HTTP Strict Transport Security. Default: &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder hsts(@Nullable Output<Boolean> hsts) {
            $.hsts = hsts;
            return this;
        }

        /**
         * @param hsts HTTP Strict Transport Security. Default: &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder hsts(Boolean hsts) {
            return hsts(Output.of(hsts));
        }

        /**
         * @param httpHeaders HTTP headers to add to the frontend. List of string.
         * 
         * @return builder
         * 
         */
        public Builder httpHeaders(@Nullable Output<List<String>> httpHeaders) {
            $.httpHeaders = httpHeaders;
            return this;
        }

        /**
         * @param httpHeaders HTTP headers to add to the frontend. List of string.
         * 
         * @return builder
         * 
         */
        public Builder httpHeaders(List<String> httpHeaders) {
            return httpHeaders(Output.of(httpHeaders));
        }

        /**
         * @param httpHeaders HTTP headers to add to the frontend. List of string.
         * 
         * @return builder
         * 
         */
        public Builder httpHeaders(String... httpHeaders) {
            return httpHeaders(List.of(httpHeaders));
        }

        /**
         * @param port Port(s) attached to your frontend. Supports single port (numerical value),
         * range (2 dash-delimited increasing ports) and comma-separated list of &#39;single port&#39;
         * and/or &#39;range&#39;. Each port must be in the [1;49151] range
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<String> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Port(s) attached to your frontend. Supports single port (numerical value),
         * range (2 dash-delimited increasing ports) and comma-separated list of &#39;single port&#39;
         * and/or &#39;range&#39;. Each port must be in the [1;49151] range
         * 
         * @return builder
         * 
         */
        public Builder port(String port) {
            return port(Output.of(port));
        }

        /**
         * @param redirectLocation Redirection HTTP&#39;
         * 
         * @return builder
         * 
         */
        public Builder redirectLocation(@Nullable Output<String> redirectLocation) {
            $.redirectLocation = redirectLocation;
            return this;
        }

        /**
         * @param redirectLocation Redirection HTTP&#39;
         * 
         * @return builder
         * 
         */
        public Builder redirectLocation(String redirectLocation) {
            return redirectLocation(Output.of(redirectLocation));
        }

        /**
         * @param serviceName The internal name of your IP load balancing
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
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
         * @param ssl SSL deciphering. Default: &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder ssl(@Nullable Output<Boolean> ssl) {
            $.ssl = ssl;
            return this;
        }

        /**
         * @param ssl SSL deciphering. Default: &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder ssl(Boolean ssl) {
            return ssl(Output.of(ssl));
        }

        /**
         * @param zone Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public HttpFrontendState build() {
            return $;
        }
    }

}
