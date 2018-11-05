// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Key Vault Certificate.
type Certifiate struct {
	s *pulumi.ResourceState
}

// NewCertifiate registers a new resource with the given unique name, arguments, and options.
func NewCertifiate(ctx *pulumi.Context,
	name string, args *CertifiateArgs, opts ...pulumi.ResourceOpt) (*Certifiate, error) {
	if args == nil || args.CertificatePolicy == nil {
		return nil, errors.New("missing required argument 'CertificatePolicy'")
	}
	if args == nil || args.VaultUri == nil {
		return nil, errors.New("missing required argument 'VaultUri'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["certificate"] = nil
		inputs["certificatePolicy"] = nil
		inputs["name"] = nil
		inputs["tags"] = nil
		inputs["vaultUri"] = nil
	} else {
		inputs["certificate"] = args.Certificate
		inputs["certificatePolicy"] = args.CertificatePolicy
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
		inputs["vaultUri"] = args.VaultUri
	}
	inputs["certificateData"] = nil
	inputs["secretId"] = nil
	inputs["thumbprint"] = nil
	inputs["version"] = nil
	s, err := ctx.RegisterResource("azure:keyvault/certifiate:Certifiate", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Certifiate{s: s}, nil
}

// GetCertifiate gets an existing Certifiate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertifiate(ctx *pulumi.Context,
	name string, id pulumi.ID, state *CertifiateState, opts ...pulumi.ResourceOpt) (*Certifiate, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["certificate"] = state.Certificate
		inputs["certificateData"] = state.CertificateData
		inputs["certificatePolicy"] = state.CertificatePolicy
		inputs["name"] = state.Name
		inputs["secretId"] = state.SecretId
		inputs["tags"] = state.Tags
		inputs["thumbprint"] = state.Thumbprint
		inputs["vaultUri"] = state.VaultUri
		inputs["version"] = state.Version
	}
	s, err := ctx.ReadResource("azure:keyvault/certifiate:Certifiate", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Certifiate{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Certifiate) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Certifiate) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A `certificate` block as defined below, used to Import an existing certificate.
func (r *Certifiate) Certificate() *pulumi.Output {
	return r.s.State["certificate"]
}

// The raw Key Vault Certificate.
func (r *Certifiate) CertificateData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["certificateData"])
}

// A `certificate_policy` block as defined below.
func (r *Certifiate) CertificatePolicy() *pulumi.Output {
	return r.s.State["certificatePolicy"]
}

// The name of the Certificate Issuer. Possible values include `Self`, or the name of a certificate issuing authority supported by Azure. Changing this forces a new resource to be created.
func (r *Certifiate) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the associated Key Vault Secret.
func (r *Certifiate) SecretId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["secretId"])
}

// A mapping of tags to assign to the resource.
func (r *Certifiate) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The X509 Thumbprint of the Key Vault Certificate returned as hex string.
func (r *Certifiate) Thumbprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["thumbprint"])
}

// Specifies the URI used to access the Key Vault instance, available on the `azurerm_key_vault` resource.
func (r *Certifiate) VaultUri() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vaultUri"])
}

// The current version of the Key Vault Certificate.
func (r *Certifiate) Version() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["version"])
}

// Input properties used for looking up and filtering Certifiate resources.
type CertifiateState struct {
	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate interface{}
	// The raw Key Vault Certificate.
	CertificateData interface{}
	// A `certificate_policy` block as defined below.
	CertificatePolicy interface{}
	// The name of the Certificate Issuer. Possible values include `Self`, or the name of a certificate issuing authority supported by Azure. Changing this forces a new resource to be created.
	Name interface{}
	// The ID of the associated Key Vault Secret.
	SecretId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The X509 Thumbprint of the Key Vault Certificate returned as hex string.
	Thumbprint interface{}
	// Specifies the URI used to access the Key Vault instance, available on the `azurerm_key_vault` resource.
	VaultUri interface{}
	// The current version of the Key Vault Certificate.
	Version interface{}
}

// The set of arguments for constructing a Certifiate resource.
type CertifiateArgs struct {
	// A `certificate` block as defined below, used to Import an existing certificate.
	Certificate interface{}
	// A `certificate_policy` block as defined below.
	CertificatePolicy interface{}
	// The name of the Certificate Issuer. Possible values include `Self`, or the name of a certificate issuing authority supported by Azure. Changing this forces a new resource to be created.
	Name interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Specifies the URI used to access the Key Vault instance, available on the `azurerm_key_vault` resource.
	VaultUri interface{}
}
