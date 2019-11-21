package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/go-fic/fic/eri/v1/routers"
	"github.com/nttcom/go-fic/pagination"
	"github.com/nttcom/go-fic/testhelper/client"

	th "github.com/nttcom/go-fic/testhelper"
	fakeclient "github.com/nttcom/go-fic/testhelper/client"
)

const TokenID = client.TokenID

func TestListRouters(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/routers",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprintf(w, listResponse)
		})

	client := fakeclient.ServiceClient()
	count := 0

	routers.List(client, routers.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := routers.ExtractRouters(page)
		if err != nil {
			t.Errorf("Failed to extract routers: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, expectedRoutersSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListRouterAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/routers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := routers.List(fakeclient.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	allRouters, err := routers.ExtractRouters(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allRouters))
}

func TestGetRouter(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s", idRouter1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, getResponse)
	})

	ap, err := routers.Get(fakeclient.ServiceClient(), idRouter1).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &router1, ap)
}

func TestCreateRouter(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/routers",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, createRequest)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, createResponse)
		})

	createOpts := routers.CreateOpts{
		Name:          "YourRouterName",
		Area:          "JPEAST",
		UserIPAddress: "10.100.0.0/27",
		Redundant:     &bTrue,
	}
	r, err := routers.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, r.OperationStatus, "Processing")
	th.AssertDeepEquals(t, &router1Created, r)
}

func TestDeleteRouter(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s", idRouter1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := routers.Delete(fakeclient.ServiceClient(), idRouter1)
	th.AssertNoErr(t, res.Err)
}
