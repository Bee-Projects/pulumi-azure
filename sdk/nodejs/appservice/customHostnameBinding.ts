// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Hostname Binding within an App Service.
 */
export class CustomHostnameBinding extends pulumi.CustomResource {
    /**
     * Get an existing CustomHostnameBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomHostnameBindingState): CustomHostnameBinding {
        return new CustomHostnameBinding(name, <any>state, { id });
    }

    /**
     * The name of the App Service in which to add the Custom Hostname Binding. Changing this forces a new resource to be created.
     */
    public readonly appServiceName: pulumi.Output<string>;
    /**
     * Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
     */
    public readonly hostname: pulumi.Output<string>;
    /**
     * The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName: pulumi.Output<string>;

    /**
     * Create a CustomHostnameBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomHostnameBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomHostnameBindingArgs | CustomHostnameBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: CustomHostnameBindingState = argsOrState as CustomHostnameBindingState | undefined;
            inputs["appServiceName"] = state ? state.appServiceName : undefined;
            inputs["hostname"] = state ? state.hostname : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as CustomHostnameBindingArgs | undefined;
            if (!args || args.appServiceName === undefined) {
                throw new Error("Missing required property 'appServiceName'");
            }
            if (!args || args.hostname === undefined) {
                throw new Error("Missing required property 'hostname'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["appServiceName"] = args ? args.appServiceName : undefined;
            inputs["hostname"] = args ? args.hostname : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        super("azure:appservice/customHostnameBinding:CustomHostnameBinding", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomHostnameBinding resources.
 */
export interface CustomHostnameBindingState {
    /**
     * The name of the App Service in which to add the Custom Hostname Binding. Changing this forces a new resource to be created.
     */
    readonly appServiceName?: pulumi.Input<string>;
    /**
     * Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
     */
    readonly hostname?: pulumi.Input<string>;
    /**
     * The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomHostnameBinding resource.
 */
export interface CustomHostnameBindingArgs {
    /**
     * The name of the App Service in which to add the Custom Hostname Binding. Changing this forces a new resource to be created.
     */
    readonly appServiceName: pulumi.Input<string>;
    /**
     * Specifies the Custom Hostname to use for the App Service, example `www.example.com`. Changing this forces a new resource to be created.
     */
    readonly hostname: pulumi.Input<string>;
    /**
     * The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
