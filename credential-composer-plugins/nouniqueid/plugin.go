package main

import (
	"context"

	"github.com/spiffe/spire-plugin-sdk/pluginmain"
	credentialcomposerv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/credentialcomposer/v1"
)

// Plugin implements the CredentialComposer plugin
type Plugin struct {
	// UnimplementedCredentialComposerServer is embedded to satisfy gRPC
	credentialcomposerv1.UnimplementedCredentialComposerServer
}

func (p *Plugin) ComposeServerX509CA(ctx context.Context, req *credentialcomposerv1.ComposeServerX509CARequest) (*credentialcomposerv1.ComposeServerX509CAResponse, error) {
	return &credentialcomposerv1.ComposeServerX509CAResponse{}, nil
}

func (p *Plugin) ComposeServerX509SVID(ctx context.Context, req *credentialcomposerv1.ComposeServerX509SVIDRequest) (*credentialcomposerv1.ComposeServerX509SVIDResponse, error) {
	return &credentialcomposerv1.ComposeServerX509SVIDResponse{}, nil
}

func (p *Plugin) ComposeAgentX509SVID(ctx context.Context, req *credentialcomposerv1.ComposeAgentX509SVIDRequest) (*credentialcomposerv1.ComposeAgentX509SVIDResponse, error) {
	return &credentialcomposerv1.ComposeAgentX509SVIDResponse{}, nil
}

func (p *Plugin) ComposeWorkloadX509SVID(ctx context.Context, req *credentialcomposerv1.ComposeWorkloadX509SVIDRequest) (*credentialcomposerv1.ComposeWorkloadX509SVIDResponse, error) {
	d := credentialcomposerv1.DistinguishedName{
		Country:            req.Attributes.Subject.Country,
		Organization:       req.Attributes.Subject.Organization,
		OrganizationalUnit: req.Attributes.Subject.OrganizationalUnit,
		Locality:           req.Attributes.Subject.Locality,
		Province:           req.Attributes.Subject.Province,
		StreetAddress:      req.Attributes.Subject.StreetAddress,
		PostalCode:         req.Attributes.Subject.PostalCode,
		SerialNumber:       req.Attributes.Subject.SerialNumber,
		CommonName:         req.Attributes.Subject.CommonName,
		ExtraNames:         nil,
	}

	a := &credentialcomposerv1.X509SVIDAttributes{
		Subject:         &d,
		DnsSans:         req.Attributes.DnsSans,
		ExtraExtensions: req.Attributes.ExtraExtensions,
	}
	r := &credentialcomposerv1.ComposeWorkloadX509SVIDResponse{Attributes: a}

	return r, nil
}

func (p *Plugin) ComposeWorkloadJWTSVID(ctx context.Context, req *credentialcomposerv1.ComposeWorkloadJWTSVIDRequest) (*credentialcomposerv1.ComposeWorkloadJWTSVIDResponse, error) {
	return &credentialcomposerv1.ComposeWorkloadJWTSVIDResponse{}, nil
}

func main() {
	plugin := new(Plugin)
	// Serve the plugin. This function call will not return. If there is a
	// failure to serve, the process will exit with a non-zero exit code.
	pluginmain.Serve(
		credentialcomposerv1.CredentialComposerPluginServer(plugin),
	)
}
