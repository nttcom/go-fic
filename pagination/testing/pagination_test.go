package testing

import (
	fic "github.com/nttcom/go-fic"
	"github.com/nttcom/go-fic/testhelper"
)

func createClient() *fic.ServiceClient {
	return &fic.ServiceClient{
		ProviderClient: &fic.ProviderClient{TokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
