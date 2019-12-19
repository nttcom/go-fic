package client

import (
	"github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/testhelper"
)

// Fake token to use.
const TokenID = "cbc36478b0bd8e67e89469c7749d4127"

// ServiceClient returns a generic service client for use in tests.
func ServiceClient() *fic.ServiceClient {
	return &fic.ServiceClient{
		ProviderClient: &fic.ProviderClient{TokenID: TokenID},
		Endpoint:       testhelper.Endpoint(),
	}
}
