package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/go-fic/fic/eri/v1/operations"
	"github.com/nttcom/go-fic/pagination"
	"github.com/nttcom/go-fic/testhelper/client"

	th "github.com/nttcom/go-fic/testhelper"
	fakeclient "github.com/nttcom/go-fic/testhelper/client"
)

const TokenID = client.TokenID

func TestListOperations(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/operations",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprintf(w, listResponse)
		})

	client := fakeclient.ServiceClient()
	count := 0

	operations.List(client, operations.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := operations.ExtractOperations(page)
		if err != nil {
			t.Errorf("Failed to extract areas: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, expectedOperationsSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListOperationAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/operations", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := operations.List(fakeclient.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	allOperations, err := operations.ExtractOperations(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, len(allOperations))
}

func TestGetOperation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/operations/cc43d0f05df24b1aabdea46456d46e39")
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, getResponse)
	})

	o, err := operations.Get(fakeclient.ServiceClient(), "cc43d0f05df24b1aabdea46456d46e39").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &operation1, o)
}
