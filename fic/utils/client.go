package utils

import (
	"fmt"
	"os"
	"reflect"

	fic "github.com/nttcom/go-fic"
	tokens3 "github.com/nttcom/go-fic/fic/identity/v3/tokens"
)

const (
	// v3 represents Keystone v3.
	// The version can be anything from v3 to v3.x.
	v3 = "v3"
)

/*
NewClient prepares an unauthenticated ProviderClient instance.
Most users will probably prefer using the AuthenticatedClient function
instead.

This is useful if you wish to explicitly control the version of the identity
service that's used for authentication explicitly, for example.

A basic example of using this would be:

	ao, err := ecl.AuthOptionsFromEnv()
	provider, err := ecl.NewClient(ao.IdentityEndpoint)
	client, err := ecl.NewIdentityV3(provider, fic.EndpointOpts{})
*/
func NewClient(endpoint string) (*fic.ProviderClient, error) {
	base, err := BaseEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	endpoint = fic.NormalizeURL(endpoint)
	base = fic.NormalizeURL(base)

	p := new(fic.ProviderClient)
	p.IdentityBase = base
	p.IdentityEndpoint = endpoint
	p.UseTokenLock()

	return p, nil
}

/*
AuthenticatedClient logs in to an Enterprise Cloud found at the identity endpoint
specified by the options, acquires a token, and returns a Provider Client
instance that's ready to operate.

If the full path to a versioned identity endpoint was specified  (example:
http://example.com:5000/v3), that path will be used as the endpoint to query.

If a versionless endpoint was specified (example: http://example.com:5000/),
the endpoint will be queried to determine which versions of the identity service
are available, then chooses the most recent or most supported version.

Example:

	ao, err := ecl.AuthOptionsFromEnv()
	provider, err := ecl.AuthenticatedClient(ao)
	client, err := ecl.NewNetworkV2(client, fic.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
*/
func AuthenticatedClient(options fic.AuthOptions) (*fic.ProviderClient, error) {
	client, err := NewClient(options.IdentityEndpoint)
	if err != nil {
		return nil, err
	}

	err = Authenticate(client, options)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Authenticate or re-authenticate against the most recent identity service
// supported at the provided endpoint.
func Authenticate(client *fic.ProviderClient, options fic.AuthOptions) error {
	versions := []*Version{
		{ID: v3, Priority: 30, Suffix: "/v3/"},
	}

	chosen, endpoint, err := ChooseVersion(client, versions)
	if err != nil {
		return err
	}

	switch chosen.ID {
	case v3:
		return v3auth(client, endpoint, &options, fic.EndpointOpts{})
	default:
		// The switch statement must be out of date from the versions list.
		return fmt.Errorf("Unrecognized identity version: %s", chosen.ID)
	}
}

// AuthenticateV3 explicitly authenticates against the identity v3 service.
func AuthenticateV3(client *fic.ProviderClient, options tokens3.AuthOptionsBuilder, eo fic.EndpointOpts) error {
	return v3auth(client, "", options, eo)
}

func v3auth(client *fic.ProviderClient, endpoint string, opts tokens3.AuthOptionsBuilder, eo fic.EndpointOpts) error {
	// Override the generated service endpoint with the one returned by the version endpoint.
	v3Client, err := NewIdentityV3(client, eo)
	if err != nil {
		return err
	}

	if endpoint != "" {
		v3Client.Endpoint = endpoint
	}

	result := tokens3.Create(v3Client, opts)

	token, err := result.ExtractToken()
	if err != nil {
		return err
	}

	catalog, err := result.ExtractServiceCatalog()
	if err != nil {
		return err
	}

	client.TokenID = token.ID

	if opts.CanReauth() {
		// here we're creating a throw-away client (tac). it's a copy of the user's provider client, but
		// with the token and reauth func zeroed out. combined with setting `AllowReauth` to `false`,
		// this should retry authentication only once
		tac := *client
		tac.ReauthFunc = nil
		tac.TokenID = ""
		var tao tokens3.AuthOptionsBuilder
		switch ot := opts.(type) {
		case *fic.AuthOptions:
			o := *ot
			o.AllowReauth = false
			tao = &o
		case *tokens3.AuthOptions:
			o := *ot
			o.AllowReauth = false
			tao = &o
		default:
			tao = opts
		}
		client.ReauthFunc = func() error {
			err := v3auth(&tac, endpoint, tao, eo)
			if err != nil {
				return err
			}
			client.TokenID = tac.TokenID
			return nil
		}
	}
	client.EndpointLocator = func(opts fic.EndpointOpts) (string, error) {
		return V3EndpointURL(catalog, opts)
	}

	return nil
}

// NewIdentityV3 creates a ServiceClient that may be used to access the v3
// identity service.
func NewIdentityV3(client *fic.ProviderClient, eo fic.EndpointOpts) (*fic.ServiceClient, error) {
	endpoint := client.IdentityBase + "v3/"
	clientType := "identity"
	var err error
	if !reflect.DeepEqual(eo, fic.EndpointOpts{}) {
		eo.ApplyDefaults(clientType)
		endpoint, err = client.EndpointLocator(eo)
		if err != nil {
			return nil, err
		}
	}

	// Ensure endpoint still has a suffix of v3.
	// This is because EndpointLocator might have found a versionless
	// endpoint or the published endpoint is still /v2.0. In both
	// cases, we need to fix the endpoint to point to /v3.
	base, err := BaseEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	base = fic.NormalizeURL(base)

	endpoint = base + "v3/"

	return &fic.ServiceClient{
		ProviderClient: client,
		Endpoint:       endpoint,
		Type:           clientType,
	}, nil
}

func initClientOpts(client *fic.ProviderClient, eo fic.EndpointOpts, clientType string) (*fic.ServiceClient, error) {
	sc := new(fic.ServiceClient)
	eo.ApplyDefaults(clientType)
	url, err := client.EndpointLocator(eo)
	if err != nil {
		return sc, err
	}
	sc.ProviderClient = client
	sc.Endpoint = url
	sc.Type = clientType
	return sc, nil
}

// NewEriV1 creates a ServiceClient that may be used with the v1
// FIC ERI package.
func NewEriV1(client *fic.ProviderClient, eo fic.EndpointOpts) (*fic.ServiceClient, error) {
	c, err := initClientOpts(client, eo, "fic-eri")

	if err != nil {
		return c, err
	}

	if os.Getenv("STATIC_FIC_ERI_ENDPOINT") != "" {
		c.Endpoint = os.Getenv("STATIC_FIC_ERI_ENDPOINT")
	}
	c.Endpoint += "fic-eri/v1/"
	return c, err
}
