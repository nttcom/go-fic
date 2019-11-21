package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/go-fic/pagination"
	"github.com/nttcom/go-fic/testhelper/client"

	"github.com/nttcom/go-fic/fic/eri/v1/routers/components/firewalls"

	th "github.com/nttcom/go-fic/testhelper"
	fakeclient "github.com/nttcom/go-fic/testhelper/client"
)

const TokenID = client.TokenID

func TestListFirewalls(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/firewalls", idRouter)
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

	firewalls.List(client, idRouter, firewalls.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := firewalls.ExtractFirewalls(page)
		if err != nil {
			t.Errorf("Failed to extract firewalls: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, expectedFirewallsSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListFirewallAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/firewalls", idRouter)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := firewalls.List(fakeclient.ServiceClient(), idRouter, nil).AllPages()
	th.AssertNoErr(t, err)
	allFirewalls, err := firewalls.ExtractFirewalls(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, len(allFirewalls))
}

func TestGetFirewall(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/firewalls/%s", idRouter, idFirewall1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, getResponse)
	})

	fw, err := firewalls.Get(fakeclient.ServiceClient(), idRouter, idFirewall1).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &firewall1, fw)
}

func TestActivateFirewall(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/firewalls/%s/activate", idRouter, idFirewall1)
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

	activateOpts := firewalls.ActivateOpts{
		UserIPAddresses: []string{
			"192.168.0.0/30",
			"192.168.0.4/30",
			"192.168.0.8/30",
			"192.168.0.12/30",
		},
	}
	f, err := firewalls.Activate(fakeclient.ServiceClient(), idRouter, idFirewall1, activateOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, f.OperationStatus, "Processing")
	th.AssertDeepEquals(t, &firewallActivated, f)
}

func TestDeactivateFirewall(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/firewalls/%s/deactivate", idRouter, idFirewall1)
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

	f, err := firewalls.Deactivate(fakeclient.ServiceClient(), idRouter, idFirewall1).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, f.OperationStatus, "Processing")
	th.AssertDeepEquals(t, &firewallDeactivated, f)
}

func TestUpdateFirewall(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/routers/%s/firewalls/%s", idRouter, idFirewall1)
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

	updateOpts := firewalls.UpdateOpts{
		Rules: []firewalls.Rule{
			firewalls.Rule{
				From: "group_1",
				To:   "group_2",
				Entries: []firewalls.Entry{
					firewalls.Entry{
						Name: "rule-01",
						Match: firewalls.Match{
							SourceAddressSets: []string{
								"group1_addset_1",
							},
							DestinationAddressSets: []string{
								"group2_addset_1",
							},
							Application: "app_set_1",
						},
						Action: "permit",
					},
				},
			},
		},
		CustomApplications: []firewalls.CustomApplication{
			firewalls.CustomApplication{
				Name:            "google-drive-web",
				Protocol:        "tcp",
				DestinationPort: "443",
			},
		},
		ApplicationSets: []firewalls.ApplicationSet{
			firewalls.ApplicationSet{
				Name: "app_set_1",
				Applications: []string{
					"google-drive-web",
					"pre-defined-ftp",
				},
			},
		},
		RoutingGroupSettings: []firewalls.RoutingGroupSetting{
			firewalls.RoutingGroupSetting{
				GroupName: "group_1",
				AddressSets: []firewalls.AddressSet{
					firewalls.AddressSet{
						Name: "group1_addset_1",
						Addresses: []string{
							"172.18.1.0/24",
						},
					},
				},
			},
			firewalls.RoutingGroupSetting{
				GroupName: "group_2",
				AddressSets: []firewalls.AddressSet{
					firewalls.AddressSet{
						Name: "group2_addset_1",
						Addresses: []string{
							"192.168.1.0/24",
						},
					},
				},
			},
			firewalls.RoutingGroupSetting{
				GroupName:   "group_3",
				AddressSets: []firewalls.AddressSet{},
			},
			firewalls.RoutingGroupSetting{
				GroupName:   "group_4",
				AddressSets: []firewalls.AddressSet{},
			},
		},
	}

	f, err := firewalls.Update(fakeclient.ServiceClient(), idRouter, idFirewall1, updateOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, f.OperationStatus, "Processing")
	th.AssertDeepEquals(t, &firewallUpdated, f)
}
