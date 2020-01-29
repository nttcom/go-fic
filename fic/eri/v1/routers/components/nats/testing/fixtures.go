package testing

import (
	"fmt"

	// "github.com/nttcom/go-fic/fic/eri/v1/ports"

	"github.com/nttcom/go-fic/fic/eri/v1/routers/components/nats"
)

const idRouter = "F000000000001"
const idNAT1 = "F020123456789"

var listResponse = fmt.Sprintf(`
{
    "nats": [
        {
            "id": "%s",
            "tenantId": "080f290761484afabbec22938adc6a2e",
            "redundant": false,
            "isActivated": true,
            "operationStatus": "Processing",
            "sourceNaptRules": [
                {
                    "from": [
                        "group_1"
                    ],
                    "to": "group_2",
                    "entries": [
                        {
                            "then": [
                                "src-set-01",
                                "src-set-02",
                                "src-set-03",
                                "src-set-04"
                            ]
                        }
                    ]
                },
                {
                    "from": [
                        "group_2"
                    ],
                    "to": "group_1",
                    "entries": [
                        {
                            "then": [
                                "src-set-05",
                                "src-set-06",
                                "src-set-07",
                                "src-set-08"
                            ]
                        }
                    ]
                }
            ],
            "destinationNatRules": [
                {
                    "from": "group_1",
                    "to": "group_2",
                    "entries": [
                        {
                            "match": {
                                "destinationAddress": "dst-set-01"
                            },
                            "then": "192.168.0.1/32"
                        },
                        {
                            "match": {
                                "destinationAddress": "dst-set-02"
                            },
                            "then": "192.168.0.2/32"
                        }
                    ]
                },
                {
                    "from": "group_2",
                    "to": "group_1",
                    "entries": [
                        {
                            "match": {
                                "destinationAddress": "dst-set-03"
                            },
                            "then": "192.168.0.3/32"
                        },
                        {
                            "match": {
                                "destinationAddress": "dst-set-04"
                            },
                            "then": "192.168.0.4/32"
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
	idNAT1,
)

var nat1 = nats.NAT{
	ID:              idNAT1,
	Redundant:       false,
	IsActivated:     true,
	TenantID:        "080f290761484afabbec22938adc6a2e",
	OperationStatus: "Processing",
	SourceNAPTRules: []nats.SourceNAPTRule{
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
	},
	DestinationNATRules: []nats.DestinationNATRule{
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
	},
	UserIPAddresses: []string{
		"192.168.0.0/30",
		"192.168.0.4/30",
		"192.168.0.8/30",
		"192.168.0.12/30",
	},
}

var expectedNATsSlice = []nats.NAT{
	nat1,
}

var getResponse = fmt.Sprintf(`
{
    "nat": {
        "id": "%s",
        "tenantId": "080f290761484afabbec22938adc6a2e",
        "redundant": false,
        "isActivated": true,
        "operationStatus": "Processing",
        "sourceNaptRules": [
            {
                "from": [
                    "group_1"
                ],
                "to": "group_2",
                "entries": [
                    {
                        "then": [
                            "src-set-01",
                            "src-set-02",
                            "src-set-03",
                            "src-set-04"
                        ]
                    }
                ]
            },
            {
                "from": [
                    "group_2"
                ],
                "to": "group_1",
                "entries": [
                    {
                        "then": [
                            "src-set-05",
                            "src-set-06",
                            "src-set-07",
                            "src-set-08"
                        ]
                    }
                ]
            }
        ],
        "destinationNatRules": [
            {
                "from": "group_1",
                "to": "group_2",
                "entries": [
                    {
                        "match": {
                            "destinationAddress": "dst-set-01"
                        },
                        "then": "192.168.0.1/32"
                    },
                    {
                        "match": {
                            "destinationAddress": "dst-set-02"
                        },
                        "then": "192.168.0.2/32"
                    }
                ]
            },
            {
                "from": "group_2",
                "to": "group_1",
                "entries": [
                    {
                        "match": {
                            "destinationAddress": "dst-set-03"
                        },
                        "then": "192.168.0.3/32"
                    },
                    {
                        "match": {
                            "destinationAddress": "dst-set-04"
                        },
                        "then": "192.168.0.4/32"
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
	idNAT1,
)

const activateRequest = `
{
    "nat": {
        "userIpAddresses": [
            "192.168.0.0/30",
            "192.168.0.4/30",
            "192.168.0.8/30",
            "192.168.0.12/30"
        ],
        "globalIpAddressSets": [
            {
                "name": "src-set-01",
                "type": "sourceNapt",
                "numOfAddresses": 5
            },
            {
                "name": "dst-set-01",
                "type": "destinationNat",
                "numOfAddresses": 1
            }
        ]
    }
}`

var activateResponse = fmt.Sprintf(`
{
    "nat": {
        "id": "%s",
        "tenantId": "080f290761484afabbec22938adc6a2e",
        "redundant": false,
        "isActivated": true,
        "operationStatus": "Processing",
        "sourceNaptRules": [],
        "destinationNatRules": [],
        "userIpAddresses": [
            "192.168.0.0/30",
            "192.168.0.4/30",
            "192.168.0.8/30",
            "192.168.0.12/30"
        ],
        "globalIpAddressSets": [
            {
                "id": "4033447dc8c548e3afb5432f7deaf0cf",
                "name": "src-set-01",
                "type": "sourceNapt",
                "natComponentId": "F050123456789",
                "operationStatus": "Processing",
                "tenantId": "080f290761484afabbec22938adc6a2e",
                "numOfAddresses": 5,
                "addresses": [
                    "100.131.66.79",
                    "100.131.66.80",
                    "100.131.66.81",
                    "100.131.66.82",
                    "100.131.66.83"
                ]
            },
            {
                "id": "2754efba42f84ae99095b4cdd654896a",
                "name": "dst-set-01",
                "type": "destinationNat",
                "natComponentId": "F050123456789",
                "operationStatus": "Processing",
                "tenantId": "080f290761484afabbec22938adc6a2e",
                "numOfAddresses": 1,
                "addresses": [
                    "100.131.65.2"
                ]
            }
        ],
        "operationId": "a0d04c09a7b6487699467dd5970865dc"
    }
}`,
	idNAT1,
)

var natActivated = nats.NAT{
	ID:                  idNAT1,
	TenantID:            "080f290761484afabbec22938adc6a2e",
	Redundant:           false,
	IsActivated:         true,
	OperationStatus:     "Processing",
	SourceNAPTRules:     []nats.SourceNAPTRule{},
	DestinationNATRules: []nats.DestinationNATRule{},
	UserIPAddresses: []string{
		"192.168.0.0/30",
		"192.168.0.4/30",
		"192.168.0.8/30",
		"192.168.0.12/30",
	},
	GlobalIPAddressSets: []nats.GlobalIpAddressSets{
		{
			ID:                "4033447dc8c548e3afb5432f7deaf0cf",
			Name:              "src-set-01",
			Type:              "sourceNapt",
			NATComponentID:    "F050123456789",
			OperationStatus:   "Processing",
			TenantID:          "080f290761484afabbec22938adc6a2e",
			NumberOfAddresses: 5,
			Addresses: []string{
				"100.131.66.79",
				"100.131.66.80",
				"100.131.66.81",
				"100.131.66.82",
				"100.131.66.83",
			},
		},
		{
			ID:                "2754efba42f84ae99095b4cdd654896a",
			Name:              "dst-set-01",
			Type:              "destinationNat",
			NATComponentID:    "F050123456789",
			OperationStatus:   "Processing",
			TenantID:          "080f290761484afabbec22938adc6a2e",
			NumberOfAddresses: 1,
			Addresses: []string{
				"100.131.65.2",
			},
		},
	},
	OperationID: "a0d04c09a7b6487699467dd5970865dc",
}

const updateRequest = `
{
    "nat": {
        "sourceNaptRules": [
            {
                "from": [
                    "group_1"
                ],
                "to": "group_2",
                "entries": [
                    {
                        "then": [
                            "src-set-01",
                            "src-set-02",
                            "src-set-03",
                            "src-set-04"
                        ]
                    }
                ]
            },
            {
                "from": [
                    "group_2"
                ],
                "to": "group_1",
                "entries": [
                    {
                        "then": [
                            "src-set-05",
                            "src-set-06",
                            "src-set-07",
                            "src-set-08"
                        ]
                    }
                ]
            }
        ],
        "destinationNatRules": [
            {
                "from": "group_1",
                "to": "group_2",
                "entries": [
                    {
                        "match": {
                            "destinationAddress": "dst-set-01"
                        },
                        "then": "192.168.0.1/32"
                    },
                    {
                        "match": {
                            "destinationAddress": "dst-set-02"
                        },
                        "then": "192.168.0.2/32"
                    }
                ]
            },
            {
                "from": "group_2",
                "to": "group_1",
                "entries": [
                    {
                        "match": {
                            "destinationAddress": "dst-set-03"
                        },
                        "then": "192.168.0.3/32"
                    },
                    {
                        "match": {
                            "destinationAddress": "dst-set-04"
                        },
                        "then": "192.168.0.4/32"
                    }
                ]
            }
        ]
    }
}
`

var updateResponse = fmt.Sprintf(`
{
    "nat": {
        "id": "%s",
        "tenantId": "080f290761484afabbec22938adc6a2e",
        "redundant": false,
        "isActivated": true,
        "operationStatus": "Processing",
        "sourceNaptRules": [
            {
                "from": [
                    "group_1"
                ],
                "to": "group_2",
                "entries": [
                    {
                        "then": [
                            "src-set-01",
                            "src-set-02",
                            "src-set-03",
                            "src-set-04"
                        ]
                    }
                ]
            },
            {
                "from": [
                    "group_2"
                ],
                "to": "group_1",
                "entries": [
                    {
                        "then": [
                            "src-set-05",
                            "src-set-06",
                            "src-set-07",
                            "src-set-08"
                        ]
                    }
                ]
            }
        ],
        "destinationNatRules": [
            {
                "from": "group_1",
                "to": "group_2",
                "entries": [
                    {
                        "match": {
                            "destinationAddress": "dst-set-01"
                        },
                        "then": "192.168.0.1/32"
                    },
                    {
                        "match": {
                            "destinationAddress": "dst-set-02"
                        },
                        "then": "192.168.0.2/32"
                    }
                ]
            },
            {
                "from": "group_2",
                "to": "group_1",
                "entries": [
                    {
                        "match": {
                            "destinationAddress": "dst-set-03"
                        },
                        "then": "192.168.0.3/32"
                    },
                    {
                        "match": {
                            "destinationAddress": "dst-set-04"
                        },
                        "then": "192.168.0.4/32"
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
        "operationId": "f6a4b449b40d4660aa3aae7985dfe2a6"
    }
}
`,
	idNAT1,
)

var natUpdated = nats.NAT{
	ID:              idNAT1,
	TenantID:        "080f290761484afabbec22938adc6a2e",
	Redundant:       false,
	IsActivated:     true,
	OperationStatus: "Processing",
	SourceNAPTRules: []nats.SourceNAPTRule{
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
	},
	DestinationNATRules: []nats.DestinationNATRule{
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
	},

	UserIPAddresses: []string{
		"192.168.0.0/30",
		"192.168.0.4/30",
		"192.168.0.8/30",
		"192.168.0.12/30",
	},

	OperationID: "f6a4b449b40d4660aa3aae7985dfe2a6",
}

const deactivateRequest = `{}`

var deactivateResponse = fmt.Sprintf(`
{
    "nat": {
        "id": "%s",
        "tenantId": "080f290761484afabbec22938adc6a2e",
        "redundant": false,
        "isActivated": true,
        "operationStatus": "Processing",
        "sourceNaptRules": [],
        "destinationNatRules": [],
        "userIpAddresses": [
            "192.168.0.0/30",
            "192.168.0.4/30",
            "192.168.0.8/30",
            "192.168.0.12/30"
        ],
        "operationId": "c7c902cdb0784f7895379e6c665e6bb2"
    }
}
`,
	idNAT1,
)

var deactivatedNAT = nats.NAT{
	ID:                  idNAT1,
	TenantID:            "080f290761484afabbec22938adc6a2e",
	Redundant:           false,
	IsActivated:         true,
	OperationStatus:     "Processing",
	SourceNAPTRules:     []nats.SourceNAPTRule{},
	DestinationNATRules: []nats.DestinationNATRule{},
	UserIPAddresses: []string{
		"192.168.0.0/30",
		"192.168.0.4/30",
		"192.168.0.8/30",
		"192.168.0.12/30",
	},
	OperationID: "c7c902cdb0784f7895379e6c665e6bb2",
}
