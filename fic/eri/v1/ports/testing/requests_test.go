package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/go-fic/fic/eri/v1/ports"
	"github.com/nttcom/go-fic/pagination"
	"github.com/nttcom/go-fic/testhelper/client"

	th "github.com/nttcom/go-fic/testhelper"
	fakeclient "github.com/nttcom/go-fic/testhelper/client"
)

const TokenID = client.TokenID

func TestListPorts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/ports",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprintf(w, listResponse)
		})

	client := fakeclient.ServiceClient()
	count := 0

	ports.List(client, ports.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ports.ExtractPorts(page)
		if err != nil {
			t.Errorf("Failed to extract ports: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, expectedPortsSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListPortAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/ports", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := ports.List(fakeclient.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	allPorts, err := ports.ExtractPorts(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allPorts))
}

func TestGetPort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/ports/%s", idPort1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, getResponse)
	})

	ap, err := ports.Get(fakeclient.ServiceClient(), idPort1).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &port1, ap)
}

func TestCreatePortWithVLAN(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/ports",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, createRequestWithVLAN)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, createResponseWithVLAN)
		})

	createOpts := ports.CreateOpts{
		Name:          "YourPortName",
		SwitchName:    "SwitchName",
		NumberOfVLANs: 32,
		PortType:      "1G",
	}
	p, err := ports.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, p.OperationStatus, "Processing")
	th.AssertDeepEquals(t, &port1CreatedWithVLAN, p)
}

func TestCreatePortWithVLANRanges(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/ports",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, createRequestWithVLANRanges)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, createResponseWithVLANRanges)
		})

	createOpts := ports.CreateOpts{
		Name:       "YourPortName",
		SwitchName: "SwitchName",
		VLANRanges: []string{
			"113-128",
		},
		PortType: "1G",
	}
	p, err := ports.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, p.OperationStatus, "Processing")
	th.AssertDeepEquals(t, &port1CreatedWithVLANRanges, p)
}

func TestDeletePort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/ports/%s", idPort1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := ports.Delete(fakeclient.ServiceClient(), idPort1)
	th.AssertNoErr(t, res.Err)
}

func TestActivatePort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/ports/%s/activate", idPort1)
	th.Mux.HandleFunc(url,
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, activateRequest)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, createResponseWithVLAN)
		})

	p, err := ports.Activate(fakeclient.ServiceClient(), idPort1).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, p.OperationStatus, "Processing")
	th.AssertDeepEquals(t, &port1CreatedWithVLAN, p)
}
