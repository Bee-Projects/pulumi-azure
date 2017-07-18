// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class EventHub extends lumi.NamedResource implements EventHubArgs {
    public readonly location: string;
    public readonly messageRetention: number;
    public readonly eventHubName?: string;
    public readonly namespaceName: string;
    public readonly partitionCount: number;
    public /*out*/ readonly partitionIds: string[];
    public readonly resourceGroupName: string;

    constructor(name: string, args: EventHubArgs) {
        super(name);
        if (lumirt.defaultIfComputed(args.location, "") === undefined) {
            throw new Error("Property argument 'location' is required, but was missing");
        }
        this.location = args.location;
        if (lumirt.defaultIfComputed(args.messageRetention, "") === undefined) {
            throw new Error("Property argument 'messageRetention' is required, but was missing");
        }
        this.messageRetention = args.messageRetention;
        this.eventHubName = args.eventHubName;
        if (lumirt.defaultIfComputed(args.namespaceName, "") === undefined) {
            throw new Error("Property argument 'namespaceName' is required, but was missing");
        }
        this.namespaceName = args.namespaceName;
        if (lumirt.defaultIfComputed(args.partitionCount, "") === undefined) {
            throw new Error("Property argument 'partitionCount' is required, but was missing");
        }
        this.partitionCount = args.partitionCount;
        if (lumirt.defaultIfComputed(args.resourceGroupName, "") === undefined) {
            throw new Error("Property argument 'resourceGroupName' is required, but was missing");
        }
        this.resourceGroupName = args.resourceGroupName;
    }
}

export interface EventHubArgs {
    readonly location: string;
    readonly messageRetention: number;
    readonly eventHubName?: string;
    readonly namespaceName: string;
    readonly partitionCount: number;
    readonly resourceGroupName: string;
}
