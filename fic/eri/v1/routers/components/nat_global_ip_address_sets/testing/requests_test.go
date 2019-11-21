package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/go-fic/fic/eri/v1/routers/components/nat_global_ip_address_sets"
	"github.com/nttcom/go-fic/pagination"
	"github.com/nttcom/go-fic/testhelper/client"

	th "github.com/nttcom/go-fic/testhelper"
	fakeclient "github.com/nttcom/go-fic/testhelper/client"
)

const TokenID = client.TokenID

func TestListGlobalIPAddressSets(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/nats/%s/global-ip-address-sets", idRouter, idNAT)
	th.Mux.HandleFunc(
		url,
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprintf(w, listResponse)
		})

	client := fakeclient.ServiceClient()
	count := 0

	nat_global_ip_address_sets.List(
		client, idRouter, idNAT,
		nat_global_ip_address_sets.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := nat_global_ip_address_sets.ExtractGlobalIPAddressSets(page)
		if err != nil {
			t.Errorf("Failed to extract ports: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, expectedGlobalIPAddressSetSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListGlobalIPAddressSetAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/nats/%s/global-ip-address-sets", idRouter, idNAT)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := nat_global_ip_address_sets.List(fakeclient.ServiceClient(), idRouter, idNAT, nil).AllPages()
	th.AssertNoErr(t, err)
	allGIPs, err := nat_global_ip_address_sets.ExtractGlobalIPAddressSets(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allGIPs))
}

func TestGetGlobalIPAddressSet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/nats/%s/global-ip-address-sets/%s", idRouter, idNAT, idGIP1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, getResponse)
	})

	ip, err := nat_global_ip_address_sets.Get(fakeclient.ServiceClient(), idRouter, idNAT, idGIP1).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &gip1, ip)
}

func TestCreateGLobalIPAddressSet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/nats/%s/global-ip-address-sets", idRouter, idNAT)
	th.Mux.HandleFunc(url,
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

	createOpts := nat_global_ip_address_sets.CreateOpts{
		Name:              "src-set-02",
		Type:              "sourceNapt",
		NumberOfAddresses: 5,
	}
	p, err := nat_global_ip_address_sets.Create(fakeclient.ServiceClient(), idRouter, idNAT, createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, p.OperationStatus, "Processing")
	th.AssertDeepEquals(t, &gip1Created, p)
}

func TestDeletePort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/nats/%s/global-ip-address-sets/%s", idRouter, idNAT, idGIP1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := nat_global_ip_address_sets.Delete(fakeclient.ServiceClient(), idRouter, idNAT, idGIP1)
	th.AssertNoErr(t, res.Err)
}
