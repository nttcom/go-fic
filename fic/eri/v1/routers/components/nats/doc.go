/*
Package nats contains functionality for working with
FIC Component-NAT resources.

Example to List NATs

	listOpts := nats.ListOpts{}

	allPages, err := nats.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allNATs, err := nats.ExtractNATs(allPages)
	if err != nil {
		panic(err)
	}

	for _, nat := range allNATs {
		fmt.Printf("%+v", nat)
	}

Example to Get Nat

	idRouter := "router1-id"
	idNAT := "nat1-id"
	n, err := nats.Get(client, idRouter, idNAT).Extract()
	if err != nil {
		panic(err)
	}


Example to Activate a NAT

	globalIPAddressSets := []nats.GlobalIPAddressSetInRequest{
		nats.GlobalIPAddressSetInRequest{
			Name:              "src-set-01",
			Type:              "sourceNapt",
			NumberOfAddresses: 5,
		},
		nats.GlobalIPAddressSetInRequest{
			Name:              "dst-set-01",
			Type:              "destinationNat",
			NumberOfAddresses: 1,
		},
	}

	activateOpts := nats.ActivateOpts{
		UserIPAddresses: []string{
			"192.168.0.0/30",
			"192.168.0.4/30",
			"192.168.0.8/30",
			"192.168.0.12/30",
		},
		GlobalIPAddressSets: globalIPAddressSets,
	}

	idRouter := "router1-id"
	idNAT := "nat1-id"

	n, err := nats.Activate(client, idRouter, idNAT, activateOpts).Extract()
	if err != nil {
		panic(err)
	}


	Example to deactivate a NAT

	idRouter := "router1-id"
	idNAT := "nat1-id"
	n, err := nats.Deactivate(client, idRouter, idNAT).Extract()
	if err != nil {
		panic(err)
	}


Example to Update a NAT

	sourceNAPTRules := []nats.SourceNAPTRuleInRequest{
		nats.SourceNAPTRuleInRequest{
			From: []string{
				"group_1",
			},
			To: "group_2",
			Entries: []nats.EntryInSourceNATRule{
				nats.EntryInSourceNATRule{
					Then: []string{
						"src-set-01",
						"src-set-02",
						"src-set-03",
						"src-set-04",
					},
				},
			},
		},
		nats.SourceNAPTRuleInRequest{
			From: []string{
				"group_2",
			},
			To: "group_1",
			Entries: []nats.EntryInSourceNATRule{
				nats.EntryInSourceNATRule{
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

	destinationNATRules := []nats.DestinationNATRuleInRequest{
		nats.DestinationNATRuleInRequest{
			From: "group_1",
			To:   "group_2",
			Entries: []nats.EntryInDestinationNATRule{
				nats.EntryInDestinationNATRule{
					Match: nats.MatchInRequest{
						DestinationAddress: "dst-set-01",
					},
					Then: "192.168.0.1/32",
				},
				nats.EntryInDestinationNATRule{
					Match: nats.MatchInRequest{
						DestinationAddress: "dst-set-02",
					},
					Then: "192.168.0.2/32",
				},
			},
		},
		nats.DestinationNATRuleInRequest{
			From: "group_2",
			To:   "group_1",
			Entries: []nats.EntryInDestinationNATRule{
				nats.EntryInDestinationNATRule{
					Match: nats.MatchInRequest{
						DestinationAddress: "dst-set-03",
					},
					Then: "192.168.0.3/32",
				},
				nats.EntryInDestinationNATRule{
					Match: nats.MatchInRequest{
						DestinationAddress: "dst-set-04",
					},
					Then: "192.168.0.4/32",
				},
			},
		},
	}

	updateOpts := &nats.UpdateOpts{
		SourceNAPTRules:     sourceNAPTRules,
		DestinationNATRules: destinationNATRules,
	}

	idRouter := "router1-id"
	idNAT := "nat1-id"
	n, err := nats.Update(client, idRouter, idNAT, opts).Extract()
	if err != nil {
		panic(err)
	}

*/
package nats
