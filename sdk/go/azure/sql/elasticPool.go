// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows you to manage an Azure SQL Elastic Pool.
type ElasticPool struct {
	s *pulumi.ResourceState
}

// NewElasticPool registers a new resource with the given unique name, arguments, and options.
func NewElasticPool(ctx *pulumi.Context,
	name string, args *ElasticPoolArgs, opts ...pulumi.ResourceOpt) (*ElasticPool, error) {
	if args == nil || args.Dtu == nil {
		return nil, errors.New("missing required argument 'Dtu'")
	}
	if args == nil || args.Edition == nil {
		return nil, errors.New("missing required argument 'Edition'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["dbDtuMax"] = nil
		inputs["dbDtuMin"] = nil
		inputs["dtu"] = nil
		inputs["edition"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["poolSize"] = nil
		inputs["resourceGroupName"] = nil
		inputs["serverName"] = nil
		inputs["tags"] = nil
	} else {
		inputs["dbDtuMax"] = args.DbDtuMax
		inputs["dbDtuMin"] = args.DbDtuMin
		inputs["dtu"] = args.Dtu
		inputs["edition"] = args.Edition
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["poolSize"] = args.PoolSize
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["serverName"] = args.ServerName
		inputs["tags"] = args.Tags
	}
	inputs["creationDate"] = nil
	s, err := ctx.RegisterResource("azure:sql/elasticPool:ElasticPool", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ElasticPool{s: s}, nil
}

// GetElasticPool gets an existing ElasticPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticPool(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ElasticPoolState, opts ...pulumi.ResourceOpt) (*ElasticPool, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationDate"] = state.CreationDate
		inputs["dbDtuMax"] = state.DbDtuMax
		inputs["dbDtuMin"] = state.DbDtuMin
		inputs["dtu"] = state.Dtu
		inputs["edition"] = state.Edition
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["poolSize"] = state.PoolSize
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["serverName"] = state.ServerName
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:sql/elasticPool:ElasticPool", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ElasticPool{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ElasticPool) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ElasticPool) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The creation date of the SQL Elastic Pool.
func (r *ElasticPool) CreationDate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationDate"])
}

// The maximum DTU which will be guaranteed to all databases in the elastic pool to be created.
func (r *ElasticPool) DbDtuMax() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["dbDtuMax"])
}

// The minimum DTU which will be guaranteed to all databases in the elastic pool to be created.
func (r *ElasticPool) DbDtuMin() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["dbDtuMin"])
}

// The total shared DTU for the elastic pool. Valid values depend on the `edition` which has been defined. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for valid combinations.
func (r *ElasticPool) Dtu() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["dtu"])
}

// The edition of the elastic pool to be created. Valid values are `Basic`, `Standard`, and `Premium`. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for details. Changing this forces a new resource to be created.
func (r *ElasticPool) Edition() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["edition"])
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *ElasticPool) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
func (r *ElasticPool) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The maximum size in MB that all databases in the elastic pool can grow to. The maximum size must be consistent with combination of `edition` and `dtu` and the limits documented in [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). If not defined when creating an elastic pool, the value is set to the size implied by `edition` and `dtu`.
func (r *ElasticPool) PoolSize() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["poolSize"])
}

// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
func (r *ElasticPool) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
func (r *ElasticPool) ServerName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serverName"])
}

// A mapping of tags to assign to the resource.
func (r *ElasticPool) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering ElasticPool resources.
type ElasticPoolState struct {
	// The creation date of the SQL Elastic Pool.
	CreationDate interface{}
	// The maximum DTU which will be guaranteed to all databases in the elastic pool to be created.
	DbDtuMax interface{}
	// The minimum DTU which will be guaranteed to all databases in the elastic pool to be created.
	DbDtuMin interface{}
	// The total shared DTU for the elastic pool. Valid values depend on the `edition` which has been defined. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for valid combinations.
	Dtu interface{}
	// The edition of the elastic pool to be created. Valid values are `Basic`, `Standard`, and `Premium`. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for details. Changing this forces a new resource to be created.
	Edition interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
	Name interface{}
	// The maximum size in MB that all databases in the elastic pool can grow to. The maximum size must be consistent with combination of `edition` and `dtu` and the limits documented in [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). If not defined when creating an elastic pool, the value is set to the size implied by `edition` and `dtu`.
	PoolSize interface{}
	// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
	ResourceGroupName interface{}
	// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
	ServerName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a ElasticPool resource.
type ElasticPoolArgs struct {
	// The maximum DTU which will be guaranteed to all databases in the elastic pool to be created.
	DbDtuMax interface{}
	// The minimum DTU which will be guaranteed to all databases in the elastic pool to be created.
	DbDtuMin interface{}
	// The total shared DTU for the elastic pool. Valid values depend on the `edition` which has been defined. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for valid combinations.
	Dtu interface{}
	// The edition of the elastic pool to be created. Valid values are `Basic`, `Standard`, and `Premium`. Refer to [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus) for details. Changing this forces a new resource to be created.
	Edition interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the elastic pool. This needs to be globally unique. Changing this forces a new resource to be created.
	Name interface{}
	// The maximum size in MB that all databases in the elastic pool can grow to. The maximum size must be consistent with combination of `edition` and `dtu` and the limits documented in [Azure SQL Database Service Tiers](https://docs.microsoft.com/en-gb/azure/sql-database/sql-database-service-tiers#elastic-pool-service-tiers-and-performance-in-edtus). If not defined when creating an elastic pool, the value is set to the size implied by `edition` and `dtu`.
	PoolSize interface{}
	// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server.
	ResourceGroupName interface{}
	// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
	ServerName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
