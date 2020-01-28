package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/go-fic/pagination"
	"github.com/nttcom/go-fic/testhelper/client"

	"github.com/nttcom/go-fic/fic/eri/v1/routers/components/nats"

	th "github.com/nttcom/go-fic/testhelper"
	fakeclient "github.com/nttcom/go-fic/testhelper/client"
)

const TokenID = client.TokenID

func TestListNATs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/nats", idRouter)
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

	nats.List(client, idRouter, nats.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := nats.ExtractNATs(page)
		if err != nil {
			t.Errorf("Failed to extract nats: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, expectedNATsSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListNATAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/nats", idRouter)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := nats.List(fakeclient.ServiceClient(), idRouter, nil).AllPages()
	th.AssertNoErr(t, err)
	allNATs, err := nats.ExtractNATs(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, len(allNATs))
}

func TestGetNAT(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/nats/%s", idRouter, idNAT1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, getResponse)
	})

	n, err := nats.Get(fakeclient.ServiceClient(), idRouter, idNAT1).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &nat1, n)
}

func TestActivateNAT(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/nats/%s/activate", idRouter, idNAT1)
	th.Mux.HandleFunc(url,
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, activateRequest)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, activateResponse)
		})

	globalIPAddressSets := []nats.GlobalIPAddressSet{
		{
			Name:              "src-set-01",
			Type:              "sourceNapt",
			NumberOfAddresses: 5,
		},
		{
			Name:              "dst-set-01",
			Type:              "destinationNat",
			NumberOfAddresses: 1,
		},
	}

	opts := nats.ActivateOpts{
		UserIPAddresses: []string{
			"192.168.0.0/30",
			"192.168.0.4/30",
			"192.168.0.8/30",
			"192.168.0.12/30",
		},
		GlobalIPAddressSets: globalIPAddressSets,
	}
	n, err := nats.Activate(fakeclient.ServiceClient(), idRouter, idNAT1, opts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, n.OperationStatus, "Processing")
	th.AssertDeepEquals(t, &natActivated, n)
}

func TestUpdateNAT(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/nats/%s", idRouter, idNAT1)
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

	sourceNAPTRules := []nats.SourceNAPTRule{
		{
			From: []string{
				"group_1",
			},
			To: "group_2",
			Entries: []nats.EntryInSourceNAPTRule{
				{
					Then: []string{
						"src-set-01",
						"src-set-02",
						"src-set-03",
						"src-set-04",
					},
				},
			},
		},
		{
			From: []string{
				"group_2",
			},
			To: "group_1",
			Entries: []nats.EntryInSourceNAPTRule{
				{
					Then: []string{
						"src-set-05",
						"src-set-06",
						"src-set-07",
						"src-set-08",
					},
				},
			},
		},
	}

	destinationNATRules := []nats.DestinationNATRule{
		{
			From: "group_1",
			To:   "group_2",
			Entries: []nats.EntryInDestinationNATRule{
				{
					Match: nats.Match{
						DestinationAddress: "dst-set-01",
					},
					Then: "192.168.0.1/32",
				},
				{
					Match: nats.Match{
						DestinationAddress: "dst-set-02",
					},
					Then: "192.168.0.2/32",
				},
			},
		},
		{
			From: "group_2",
			To:   "group_1",
			Entries: []nats.EntryInDestinationNATRule{
				{
					Match: nats.Match{
						DestinationAddress: "dst-set-03",
					},
					Then: "192.168.0.3/32",
				},
				{
					Match: nats.Match{
						DestinationAddress: "dst-set-04",
					},
					Then: "192.168.0.4/32",
				},
			},
		},
	}

	opts := &nats.UpdateOpts{
		SourceNAPTRules:     sourceNAPTRules,
		DestinationNATRules: destinationNATRules,
	}
	n, err := nats.Update(fakeclient.ServiceClient(), idRouter, idNAT1, opts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, n.OperationStatus, "Processing")
	th.AssertDeepEquals(t, &natUpdated, n)
}

func TestDeactivateNAT(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/nats/%s/deactivate", idRouter, idNAT1)
	th.Mux.HandleFunc(url,
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, deactivateRequest)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, deactivateResponse)
		})

	n, err := nats.Deactivate(fakeclient.ServiceClient(), idRouter, idNAT1).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, n.OperationStatus, "Processing")
	th.AssertDeepEquals(t, &deactivatedNAT, n)
}
