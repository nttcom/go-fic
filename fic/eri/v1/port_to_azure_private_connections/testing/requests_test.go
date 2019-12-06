package testing

import (
	"fmt"
	"net/http"
	"testing"

	con "github.com/nttcom/go-fic/fic/eri/v1/port_to_azure_private_connections"
	"github.com/nttcom/go-fic/pagination"
	"github.com/nttcom/go-fic/testhelper/client"

	th "github.com/nttcom/go-fic/testhelper"
	fakeclient "github.com/nttcom/go-fic/testhelper/client"
)

const TokenID = client.TokenID

func TestListConnections(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/port-to-azure-private-connections",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprintf(w, listResponse)
		})

	client := fakeclient.ServiceClient()
	count := 0

	con.List(
		client, con.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {

		count++
		actual, err := con.ExtractConnections(page)
		if err != nil {
			t.Errorf("Failed to extract connections: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, expectedConnectionsSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListConnectionAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/port-to-azure-private-connections", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := con.List(fakeclient.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	allConnections, err := con.ExtractConnections(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, len(allConnections))
}

func TestGetConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/port-to-azure-private-connections/%s", idConnection1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, getResponse)
	})

	c, err := con.Get(fakeclient.ServiceClient(), idConnection1).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &connection1, c)
}

func TestCreateConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/port-to-azure-private-connections",
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

	createOpts := con.CreateOpts{
		Name: "YourConnectionName",
		Source: con.Source{
			Primary: con.Primary{
				PortID: "F010123456789",
				VLAN:   1025,
			},
			Secondary: con.Secondary{
				PortID: "F019876543210",
				VLAN:   1057,
			},
			ASN: "65530",
		},
		Destination: con.Destination{
			Interconnect: "Tokyo-1",
			ServiceKey:   "6191af11-82f9-4c15-9894-9a69c8f8628a",
			SharedKey:    "a8268f8c96a9",
			QosType:      "guarantee",
		},
		Bandwidth:                        "100M",
		PrimaryConnectedNetworkAddress:   "10.10.0.0/30",
		SecondaryConnectedNetworkAddress: "10.20.0.0/30",
	}
	c, err := con.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, c.OperationStatus, "Processing")
	th.AssertDeepEquals(t, &connectionCreated, c)

}

func TestDeleteConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/port-to-azure-private-connections/%s", idConnection1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := con.Delete(fakeclient.ServiceClient(), idConnection1)
	th.AssertNoErr(t, res.Err)
}
