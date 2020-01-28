package testing

import (
	"fmt"

	"github.com/nttcom/go-fic/fic/eri/v1/routers/components/firewalls"
)

const idRouter = "R000000000001"
const idFirewall1 = "F000000000001"

var listResponse = fmt.Sprintf(`
{
    "firewalls": [
        {
            "id": "%s",
            "tenantId": "080f290761484afabbec22938adc6a2e",
            "redundant": false,
            "isActivated": true,
            "operationStatus": "Completed",
            "customApplications": [
                {
                    "name": "google-drive-web",
                    "protocol": "tcp",
                    "destinationPort": "443"
                }
            ],
            "applicationSets": [
                {
                    "name": "app_set_1",
                    "applications": [
                        "google-drive-web",
                        "pre-defined-ftp"
                    ]
                }
            ],
            "routingGroupSettings": [
                {
                    "groupName": "group_1",
                    "addressSets": [
                        {
                            "name": "group1_addset_1",
                            "addresses": [
                                "172.18.1.0/24"
                            ]
                        }
                    ]
                },
                {
                    "groupName": "group_2",
                    "addressSets": [
                        {
                            "name": "group2_addset_1",
                            "addresses": [
                                "192.168.1.0/24"
                            ]
                        }
                    ]
                },
                {
                    "groupName": "group_3",
                    "addressSets": []
                },
                {
                    "groupName": "group_4",
                    "addressSets": []
                }
            ],
            "rules": [
                {
                    "from": "group_1",
                    "to": "group_2",
                    "entries": [
                        {
                            "name": "rule-01",
                            "match": {
                                "sourceAddressSets": [
                                    "group1_addset_1"
                                ],
                                "destinationAddressSets": [
                                    "group2_addset_1"
                                ],
                                "application": "app_set_1"
                            },
                            "action": "permit"
                        }
                    ]
                }
            ],
            "userIpAddresses": [
                "192.168.0.0/30",
                "192.168.0.4/30",
                "192.168.0.8/30",
                "192.168.0.12/30"
            ]
        }
    ]
}
`,
	idFirewall1,
)

var customApplications = []firewalls.CustomApplication{
	{
		Name:            "google-drive-web",
		Protocol:        "tcp",
		DestinationPort: "443",
	},
}

var applicationSets = []firewalls.ApplicationSet{
	{
		Name: "app_set_1",
		Applications: []string{
			"google-drive-web",
			"pre-defined-ftp",
		},
	},
}

var routingGroupSettings = []firewalls.RoutingGroupSetting{
	{
		GroupName: "group_1",
		AddressSets: []firewalls.AddressSet{
			{
				Name: "group1_addset_1",
				Addresses: []string{
					"172.18.1.0/24",
				},
			},
		},
	},
	{
		GroupName: "group_2",
		AddressSets: []firewalls.AddressSet{
			{
				Name: "group2_addset_1",
				Addresses: []string{
					"192.168.1.0/24",
				},
			},
		},
	},
	{
		GroupName:   "group_3",
		AddressSets: []firewalls.AddressSet{},
	},
	{
		GroupName:   "group_4",
		AddressSets: []firewalls.AddressSet{},
	},
}

var rules = []firewalls.Rule{
	{
		From: "group_1",
		To:   "group_2",
		Entries: []firewalls.Entry{
			{
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
}

var firewall1 = firewalls.Firewall{
	ID:                   idFirewall1,
	TenantID:             "080f290761484afabbec22938adc6a2e",
	Redundant:            false,
	IsActivated:          true,
	OperationStatus:      "Completed",
	CustomApplications:   customApplications,
	ApplicationSets:      applicationSets,
	RoutingGroupSettings: routingGroupSettings,
	Rules:                rules,
	UserIPAddresses: []string{
		"192.168.0.0/30",
		"192.168.0.4/30",
		"192.168.0.8/30",
		"192.168.0.12/30",
	},
}

var expectedFirewallsSlice = []firewalls.Firewall{
	firewall1,
}

var getResponse = fmt.Sprintf(`
{
    "firewall": {
        "id": "%s",
        "tenantId": "080f290761484afabbec22938adc6a2e",
        "redundant": false,
        "isActivated": true,
        "operationStatus": "Completed",
        "customApplications": [
            {
                "name": "google-drive-web",
                "protocol": "tcp",
                "destinationPort": "443"
            }
        ],
        "applicationSets": [
            {
                "name": "app_set_1",
                "applications": [
                    "google-drive-web",
                    "pre-defined-ftp"
                ]
            }
        ],
        "routingGroupSettings": [
            {
                "groupName": "group_1",
                "addressSets": [
                    {
                        "name": "group1_addset_1",
                        "addresses": [
                            "172.18.1.0/24"
                        ]
                    }
                ]
            },
            {
                "groupName": "group_2",
                "addressSets": [
                    {
                        "name": "group2_addset_1",
                        "addresses": [
                            "192.168.1.0/24"
                        ]
                    }
                ]
            },
            {
                "groupName": "group_3",
                "addressSets": []
            },
            {
                "groupName": "group_4",
                "addressSets": []
            }
        ],
        "rules": [
            {
                "from": "group_1",
                "to": "group_2",
                "entries": [
                    {
                        "name": "rule-01",
                        "match": {
                            "sourceAddressSets": [
                                "group1_addset_1"
                            ],
                            "destinationAddressSets": [
                                "group2_addset_1"
                            ],
                            "application": "app_set_1"
                        },
                        "action": "permit"
                    }
                ]
            }
        ],
        "userIpAddresses": [
            "192.168.0.0/30",
            "192.168.0.4/30",
            "192.168.0.8/30",
            "192.168.0.12/30"
        ]        
    }
}
`,
	idFirewall1,
)

const activateRequest = `
{
    "firewall": {
        "userIpAddresses": [
            "192.168.0.0/30",
            "192.168.0.4/30",
            "192.168.0.8/30",
            "192.168.0.12/30"
        ]
    }
}
`

const activateResponse = `
{
    "firewall": {
        "id": "F040123456789",
        "tenantId": "080f290761484afabbec22938adc6a2e",
        "redundant": false,
        "isActivated": true,
        "operationStatus": "Processing",
        "customApplications": [],
        "applicationSets": [],
        "routingGroupSettings": [],
        "rules": [],
        "userIpAddresses": [
            "192.168.0.0/30",
            "192.168.0.4/30",
            "192.168.0.8/30",
            "192.168.0.12/30"
        ],
        "operationId": "4c7b0bfc17e84e9eae7b4d779f30dea1"
    }
}
`

var firewallActivated = firewalls.Firewall{
	ID:                   "F040123456789",
	TenantID:             "080f290761484afabbec22938adc6a2e",
	Redundant:            false,
	IsActivated:          true,
	OperationStatus:      "Processing",
	CustomApplications:   []firewalls.CustomApplication{},
	ApplicationSets:      []firewalls.ApplicationSet{},
	RoutingGroupSettings: []firewalls.RoutingGroupSetting{},
	Rules:                []firewalls.Rule{},
	UserIPAddresses: []string{
		"192.168.0.0/30",
		"192.168.0.4/30",
		"192.168.0.8/30",
		"192.168.0.12/30",
	},
	OperationID: "4c7b0bfc17e84e9eae7b4d779f30dea1",
}

const deactivateRequest = `{}`

var deactivateResponse = fmt.Sprintf(`
{
    "firewall": {
        "id": "%s",
        "tenantId": "080f290761484afabbec22938adc6a2e",
        "redundant": false,
        "isActivated": false,
        "operationStatus": "Processing",
        "customApplications": [],
        "applicationSets": [],
        "routingGroupSettings": [],
        "rules": [],
        "userIpAddresses": [],
        "operationId": "d40999c03a7642c6b86158889c7bebc9"
    }
}`,
	idFirewall1,
)

var firewallDeactivated = firewalls.Firewall{
	ID:                   idFirewall1,
	TenantID:             "080f290761484afabbec22938adc6a2e",
	Redundant:            false,
	IsActivated:          false,
	OperationStatus:      "Processing",
	CustomApplications:   []firewalls.CustomApplication{},
	ApplicationSets:      []firewalls.ApplicationSet{},
	RoutingGroupSettings: []firewalls.RoutingGroupSetting{},
	Rules:                []firewalls.Rule{},
	UserIPAddresses:      []string{},
	OperationID:          "d40999c03a7642c6b86158889c7bebc9",
}

const updateRequest = `
{
    "firewall": {
        "rules": [
            {
                "from": "group_1",
                "to": "group_2",
                "entries": [
                    {
                        "name": "rule-01",
                        "match": {
                            "sourceAddressSets": [
                                "group1_addset_1"
                            ],
                            "destinationAddressSets": [
                                "group2_addset_1"
                            ],
                            "application": "app_set_1"
                        },
                        "action": "permit"
                    }
                ]
            }
        ],
        "customApplications": [
            {
                "name": "google-drive-web",
                "protocol": "tcp",
                "destinationPort": "443"
            }
        ],
        "applicationSets": [
            {
                "name": "app_set_1",
                "applications": [
                    "google-drive-web",
                    "pre-defined-ftp"
                ]
            }
        ],
        "routingGroupSettings": [
            {
                "groupName": "group_1",
                "addressSets": [
                    {
                        "name": "group1_addset_1",
                        "addresses": [
                            "172.18.1.0/24"
                        ]
                    }
                ]
            },
            {
                "groupName": "group_2",
                "addressSets": [
                    {
                        "name": "group2_addset_1",
                        "addresses": [
                            "192.168.1.0/24"
                        ]
                    }
                ]
            },
            {
                "groupName": "group_3",
                "addressSets": []
            },
            {
                "groupName": "group_4",
                "addressSets": []
            }
        ]
    }
}`

var updateResponse = fmt.Sprintf(`
{
    "firewall": {
        "id": "%s",
        "tenantId": "080f290761484afabbec22938adc6a2e",
        "redundant": false,
        "isActivated": true,
        "operationStatus": "Processing",
        "customApplications": [
            {
                "name": "google-drive-web",
                "protocol": "tcp",
                "destinationPort": "443"
            }
        ],
        "applicationSets": [
            {
                "name": "app_set_1",
                "applications": [
                    "google-drive-web",
                    "pre-defined-ftp"
                ]
            }
        ],
        "routingGroupSettings": [
            {
                "groupName": "group_1",
                "addressSets": [
                    {
                        "name": "group1_addset_1",
                        "addresses": [
                            "172.18.1.0/24"
                        ]
                    }
                ]
            },
            {
                "groupName": "group_2",
                "addressSets": [
                    {
                        "name": "group2_addset_1",
                        "addresses": [
                            "192.168.1.0/24"
                        ]
                    }
                ]
            },
            {
                "groupName": "group_3",
                "addressSets": []
            },
            {
                "groupName": "group_4",
                "addressSets": []
            }
        ],
        "rules": [
            {
                "from": "group_1",
                "to": "group_2",
                "entries": [
                    {
                        "name": "rule-01",
                        "match": {
                            "sourceAddressSets": [
                                "group1_addset_1"
                            ],
                            "destinationAddressSets": [
                                "group2_addset_1"
                            ],
                            "application": "app_set_1"
                        },
                        "action": "permit"
                    }
                ]
            }
        ],
        "userIpAddresses": [
            "192.168.0.0/30",
            "192.168.0.4/30",
            "192.168.0.8/30",
            "192.168.0.12/30"
        ],
        "operationId": "2fb5baa5a5834e0b952d3f0d93c3e64a"
    }
}
`,
	idFirewall1,
)

var firewallUpdated = firewalls.Firewall{
	ID:                   idFirewall1,
	TenantID:             "080f290761484afabbec22938adc6a2e",
	Redundant:            false,
	IsActivated:          true,
	OperationStatus:      "Processing",
	CustomApplications:   customApplications,
	ApplicationSets:      applicationSets,
	RoutingGroupSettings: routingGroupSettings,
	Rules:                rules,
	UserIPAddresses: []string{
		"192.168.0.0/30",
		"192.168.0.4/30",
		"192.168.0.8/30",
		"192.168.0.12/30",
	},
	OperationID: "2fb5baa5a5834e0b952d3f0d93c3e64a",
}
