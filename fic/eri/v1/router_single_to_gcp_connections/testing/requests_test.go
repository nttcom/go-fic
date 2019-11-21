package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	con "github.com/nttcom/go-fic/fic/eri/v1/router_single_to_gcp_connections"
	"github.com/nttcom/go-fic/pagination"
	"github.com/nttcom/go-fic/testhelper/client"

	th "github.com/nttcom/go-fic/testhelper"
	fakeclient "github.com/nttcom/go-fic/testhelper/client"
)

const TokenID = client.TokenID

func TestListConnections(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/router-to-gcp-connections",
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
		// th.CheckDeepEquals(t, expectedConnectionsSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListConnectionAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/router-to-gcp-connections", func(w http.ResponseWriter, r *http.Request) {
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

	url := fmt.Sprintf("/router-to-gcp-connections/%s", idConnection1)
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

	th.Mux.HandleFunc("/router-to-gcp-connections",
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
			RouteFilter: con.RouteFilter{
				In:  "fullRoute",
				Out: "fullRouteWithDefaultRoute",
			},
			Primary: con.Primary{
				ASPathPrepend: con.ASPathPrepend{
					In:  &one,
					Out: &null,
				},
			},
		},
		Destination: con.Destination{
			Primary: con.DestinationHAInfo{
				Interconnect: "Equinix-TY2-1",
				PairingKey:   "4f043662-2ed6-45b0-8c3a-77d685716e28/asia-northeast1/2",
			},
			QosType: "guarantee",
		},
		Bandwidth: "300M",
	}
	c, err := con.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, c.OperationStatus, "Processing")
	// th.AssertDeepEquals(t, &connectionCreated, c)
	expectedJSON, _ := json.Marshal(&connectionCreated)
	actualJSON, _ := json.Marshal(c)
	th.AssertEquals(t, string(expectedJSON), string(actualJSON))

}

func TestDeleteConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/router-to-gcp-connections/%s", idConnection1)
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

	url := fmt.Sprintf("/router-to-gcp-connections/%s", idConnection1)
	th.Mux.HandleFunc(url,
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
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
			RouteFilter: con.RouteFilter{
				In:  "noRoute",
				Out: "defaultRoute",
			},
			Primary: con.Primary{
				ASPathPrepend: con.ASPathPrepend{
					In:  &one,
					Out: &two,
				},
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
