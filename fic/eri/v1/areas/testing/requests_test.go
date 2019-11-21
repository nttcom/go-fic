package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/go-fic/fic/eri/v1/areas"
	"github.com/nttcom/go-fic/pagination"
	"github.com/nttcom/go-fic/testhelper/client"

	th "github.com/nttcom/go-fic/testhelper"
	fakeclient "github.com/nttcom/go-fic/testhelper/client"
)

const TokenID = client.TokenID

func TestListAreas(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/areas",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprintf(w, listResponse)
		})

	client := fakeclient.ServiceClient()
	count := 0

	areas.List(client, areas.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := areas.ExtractAreas(page)
		if err != nil {
			t.Errorf("Failed to extract areas: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, expectedAreasSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListAreaAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/areas", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := areas.List(fakeclient.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	allAreas, err := areas.ExtractAreas(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allAreas))
}
