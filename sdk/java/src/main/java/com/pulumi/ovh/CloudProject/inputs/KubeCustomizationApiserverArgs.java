// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ovh.CloudProject.inputs.KubeCustomizationApiserverAdmissionpluginsArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubeCustomizationApiserverArgs extends com.pulumi.resources.ResourceArgs {

    public static final KubeCustomizationApiserverArgs Empty = new KubeCustomizationApiserverArgs();

    @Import(name="admissionplugins")
    private @Nullable Output<KubeCustomizationApiserverAdmissionpluginsArgs> admissionplugins;

    public Optional<Output<KubeCustomizationApiserverAdmissionpluginsArgs>> admissionplugins() {
        return Optional.ofNullable(this.admissionplugins);
    }

    private KubeCustomizationApiserverArgs() {}

    private KubeCustomizationApiserverArgs(KubeCustomizationApiserverArgs $) {
        this.admissionplugins = $.admissionplugins;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubeCustomizationApiserverArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubeCustomizationApiserverArgs $;

        public Builder() {
            $ = new KubeCustomizationApiserverArgs();
        }

        public Builder(KubeCustomizationApiserverArgs defaults) {
            $ = new KubeCustomizationApiserverArgs(Objects.requireNonNull(defaults));
        }

        public Builder admissionplugins(@Nullable Output<KubeCustomizationApiserverAdmissionpluginsArgs> admissionplugins) {
            $.admissionplugins = admissionplugins;
            return this;
        }

        public Builder admissionplugins(KubeCustomizationApiserverAdmissionpluginsArgs admissionplugins) {
            return admissionplugins(Output.of(admissionplugins));
        }

        public KubeCustomizationApiserverArgs build() {
            return $;
        }
    }

}
