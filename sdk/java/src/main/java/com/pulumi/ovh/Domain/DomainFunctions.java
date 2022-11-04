// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Domain;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.ovh.Domain.inputs.GetZoneArgs;
import com.pulumi.ovh.Domain.inputs.GetZonePlainArgs;
import com.pulumi.ovh.Domain.outputs.GetZoneResult;
import com.pulumi.ovh.Utilities;
import java.util.concurrent.CompletableFuture;

public final class DomainFunctions {
    /**
     * Use this data source to retrieve information about a domain zone.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Domain.DomainFunctions;
     * import com.pulumi.ovh.Domain.inputs.GetZoneArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var rootzone = DomainFunctions.getZone(GetZoneArgs.builder()
     *             .name(&#34;mysite.ovh&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetZoneResult> getZone(GetZoneArgs args) {
        return getZone(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about a domain zone.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Domain.DomainFunctions;
     * import com.pulumi.ovh.Domain.inputs.GetZoneArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var rootzone = DomainFunctions.getZone(GetZoneArgs.builder()
     *             .name(&#34;mysite.ovh&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetZoneResult> getZonePlain(GetZonePlainArgs args) {
        return getZonePlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information about a domain zone.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Domain.DomainFunctions;
     * import com.pulumi.ovh.Domain.inputs.GetZoneArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var rootzone = DomainFunctions.getZone(GetZoneArgs.builder()
     *             .name(&#34;mysite.ovh&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetZoneResult> getZone(GetZoneArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Domain/getZone:getZone", TypeShape.of(GetZoneResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information about a domain zone.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Domain.DomainFunctions;
     * import com.pulumi.ovh.Domain.inputs.GetZoneArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var rootzone = DomainFunctions.getZone(GetZoneArgs.builder()
     *             .name(&#34;mysite.ovh&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetZoneResult> getZonePlain(GetZonePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Domain/getZone:getZone", TypeShape.of(GetZoneResult.class), args, Utilities.withVersion(options));
    }
}
