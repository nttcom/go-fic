/*
Package firewalls contains functionality for working with
FIC Firewall resources.

Example to List Firewalls

	listOpts := firewalls.ListOpts{}

	idRouter = "router1-id"
	allPages, err := .List(client, idRouter, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allFirewalls, err := firewalls.ExtractFirewalls(allPages)
	if err != nil {
		panic(err)
	}

	for _, f := range allFirewalls {
		fmt.Printf("%+v", f)
	}

Example to Show a Firewall

	idRouter = "router1-id"
	idFirewall = "firewall1-id"
	fw, err := firewalls.Get(fakeclient.ServiceClient(), idRouter, idFirewall).Extract()
	if err != nil {
		panic(err)
	}


Example to Activate a Firewall

	activateOpts := firewalls.ActivateOpts{
		UserIPAddresses: []string{
			"192.168.0.0/30",
			"192.168.0.4/30",
			"192.168.0.8/30",
			"192.168.0.12/30",
		},
	}

	idRouter = "router1-id"
	idFirewall = "firewall1-id"
	f, err := firewalls.Activate(fakeclient.ServiceClient(), idRouter, idFirewall, activateOpts).Extract()
	if err != nil {
		panic(err)
	}


Example to Deactivate a Firewall

	idRouter = "router1-id"
	idFirewall = "firewall1-id"
	f, err := firewalls.Deactivate(fakeclient.ServiceClient(), idRouter, idFirewall).Extract()
	if err != nil {
		panic(err)
	}


Example to Update a Firewall

	updateOpts := firewalls.UpdateOpts{
		Rules: []firewalls.RuleInRequest{
			firewalls.RuleInRequest{
				From: "group_1",
				To:   "group_2",
				Entries: []firewalls.EntryInRequest{
					firewalls.EntryInRequest{
						Name: "rule-01",
						Match: firewalls.MatchInRequest{
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
		CustomApplications: []firewalls.CustomApplicationInRequest{
			firewalls.CustomApplicationInRequest{
				Name:            "google-drive-web",
				Protocol:        "tcp",
				DestinationPort: "443",
			},
		},
		ApplicationSets: []firewalls.ApplicationSetInRequest{
			firewalls.ApplicationSetInRequest{
				Name: "app_set_1",
				Applications: []string{
					"google-drive-web",
					"pre-defined-ftp",
				},
			},
		},
		RoutingGroupSettings: []firewalls.RoutingGroupSettingInRequest{
			firewalls.RoutingGroupSettingInRequest{
				GroupName: "group_1",
				AddressSets: []firewalls.AddressSetInRequest{
					firewalls.AddressSetInRequest{
						Name: "group1_addset_1",
						Addresses: []string{
							"172.18.1.0/24",
						},
					},
				},
			},
			firewalls.RoutingGroupSettingInRequest{
				GroupName: "group_2",
				AddressSets: []firewalls.AddressSetInRequest{
					firewalls.AddressSetInRequest{
						Name: "group2_addset_1",
						Addresses: []string{
							"192.168.1.0/24",
						},
					},
				},
			},
			firewalls.RoutingGroupSettingInRequest{
				GroupName:   "group_3",
				AddressSets: []firewalls.AddressSetInRequest{},
			},
			firewalls.RoutingGroupSettingInRequest{
				GroupName:   "group_4",
				AddressSets: []firewalls.AddressSetInRequest{},
			},
		},
	}

	idRouter = "router1-id"
	idFirewall = "firewall1-id"
	f, err := firewalls.Update(fakeclient.ServiceClient(), idRouter, idFirewall, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

*/
package firewalls
