package testing

import (
	"testing"

	fic "github.com/nttcom/go-fic"
	th "github.com/nttcom/go-fic/testhelper"
)

func TestApplyDefaultsToEndpointOpts(t *testing.T) {
	eo := fic.EndpointOpts{Availability: fic.AvailabilityPublic}
	eo.ApplyDefaults("compute")
	expected := fic.EndpointOpts{Availability: fic.AvailabilityPublic, Type: "compute"}
	th.CheckDeepEquals(t, expected, eo)

	eo = fic.EndpointOpts{Type: "compute"}
	eo.ApplyDefaults("object-store")
	expected = fic.EndpointOpts{Availability: fic.AvailabilityPublic, Type: "compute"}
	th.CheckDeepEquals(t, expected, eo)
}
