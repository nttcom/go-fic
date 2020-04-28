package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	con "github.com/nttcom/go-fic/fic/eri/v1/router_single_to_port_connections"
	"github.com/nttcom/go-fic/pagination"
	"github.com/nttcom/go-fic/testhelper/client"

	th "github.com/nttcom/go-fic/testhelper"
	fakeclient "github.com/nttcom/go-fic/testhelper/client"
)

const TokenID = client.TokenID

func TestListConnections(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/router-to-port-connections",
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

		expectedJSON, _ := json.Marshal(expectedConnectionsSlice)
		actualJSON, _ := json.Marshal(actual)
		th.AssertEquals(t, string(expectedJSON), string(actualJSON))

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListConnectionAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/router-to-port-connections", func(w http.ResponseWriter, r *http.Request) {
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

	url := fmt.Sprintf("/router-to-port-connections/%s", idConnection1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, getResponse)
	})

	c, err := con.Get(fakeclient.ServiceClient(), idConnection1).Extract()
	th.AssertNoErr(t, err)
	expectedJSON, _ := json.Marshal(&connection1)
	actualJSON, _ := json.Marshal(c)
	th.AssertEquals(t, string(expectedJSON), string(actualJSON))
}

func TestCreateConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/router-to-port-connections",
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
			RouterID:  "F020123456789",
			GroupName: "group_1",
			Primary: con.SourceHAInfo{
				IPAddress: "10.0.0.1/30",
				ASPathPrepend: con.ASPathPrepend{
					In:  &four,
					Out: &four,
				},
				MED: &con.MED{
					Out: 10,
				},
			},
			// Secondary: con.SourceHAInfo{
			// 	IPAddress: "10.0.0.5/30",
			// 	ASPathPrepend: con.ASPathPrepend{
			// 		In:  &two,
			// 		Out: &one,
			// 	},
			// },
			RouteFilter: con.RouteFilter{
				In:  "fullRoute",
				Out: "fullRouteWithDefaultRoute",
			},
		},
		Destination: con.Destination{
			Primary: con.DestinationHAInfo{
				PortID:    "F010123456789",
				VLAN:      101,
				IPAddress: "10.0.0.2/30",
				ASN:       "65000",
			},
			// Secondary: con.DestinationHAInfo{
			// 	PortID:    "F019876543210",
			// 	VLAN:      102,
			// 	IPAddress: "10.0.0.6/30",
			// 	ASN:       "65000",
			// },
		},
		Bandwidth: "100M",
	}
	c, err := con.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, c.OperationStatus, "Processing")

	expectedJSON, _ := json.Marshal(&connectionCreated)
	actualJSON, _ := json.Marshal(c)
	th.AssertEquals(t, string(expectedJSON), string(actualJSON))
}

func TestDeleteConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/router-to-port-connections/%s", idConnection1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := con.Delete(fakeclient.ServiceClient(), idConnection1)
	th.AssertNoErr(t, res.Err)
}

func TestUpdateConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/router-to-port-connections/%s", idConnection1)
	th.Mux.HandleFunc(url,
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PATCH")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, updateRequest)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, updateResponse)
		})

	updateOpts := con.UpdateOpts{
		Source: con.SourceForUpdate{
			Primary: con.SourceHAInfoForUpdate{
				ASPathPrepend: con.ASPathPrepend{
					In:  &null,
					Out: &two,
				},
				MED: &con.MED{
					Out: 30,
				},
			},
			// Secondary: con.SourceHAInfoForUpdate{
			// 	ASPathPrepend: con.ASPathPrepend{
			// 		In:  &null,
			// 		Out: &two,
			// 	},
			// },
			RouteFilter: con.RouteFilter{
				In:  "noRoute",
				Out: "fullRoute",
			},
		},
	}
	c, err := con.Update(fakeclient.ServiceClient(), idConnection1, updateOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, c.OperationStatus, "Processing")

	expectedJSON, _ := json.Marshal(&connectionUpdated)
	actualJSON, _ := json.Marshal(c)
	th.AssertEquals(t, string(expectedJSON), string(actualJSON))
}
