package testing

import (
	"fmt"

	con "github.com/nttcom/go-fic/fic/eri/v1/router_paired_to_gcp_connections"
)

const idConnection1 = "F010123456789"

var listResponse = fmt.Sprintf(`
{
    "connections": [
        {
            "id": "%s",
            "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
            "operationStatus": "Completed",
            "redundant": true,
            "name": "YourConnectionName",
            "bandwidth": "100M",
            "source": {
                "routerId": "F020123456789",
                "groupName": "group_1",
                "routeFilter": {
                    "in": "noRoute",
                    "out": "privateRoute"
                }
            },
            "destination": {
                "primary": {
                    "interconnect": "@Tokyo-CC2-1",
                    "pairingKey": "d27476e6-e8a8-4214-a88f-9d3131db465d/asia-northeast1/2"
                },
                "qosType": "guarantee",
                "secondary": {
                    "interconnect": "@Tokyo-CC2-1",
                    "pairingKey": "17c64c4e-f845-4450-82e9-843095e18526/asia-northeast1/2"
                }
            },
            "primaryConnectedNwAddress": "169.254.51.168/29",
            "secondaryConnectedNwAddress": "169.254.144.248/29"
        }
    ]
}`,
	idConnection1,
)

var connection1 = con.Connection{
	ID:              idConnection1,
	TenantID:        "87e89b8f075a4ee1aa209f6ca6ce242c",
	OperationStatus: "Completed",
	Redundant:       true,
	Name:            "YourConnectionName",
	Bandwidth:       "100M",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		RouteFilter: con.RouteFilter{
			In:  "noRoute",
			Out: "privateRoute",
		},
	},
	Destination: con.Destination{
		Primary: con.DestinationHAInfo{
			Interconnect: "@Tokyo-CC2-1",
			PairingKey:   "d27476e6-e8a8-4214-a88f-9d3131db465d/asia-northeast1/2",
		},
		QosType: "guarantee",
		Secondary: con.DestinationHAInfo{
			Interconnect: "@Tokyo-CC2-1",
			PairingKey:   "17c64c4e-f845-4450-82e9-843095e18526/asia-northeast1/2",
		},
	},
	PrimaryConnectedNetworkAddress:   "169.254.51.168/29",
	SecondaryConnectedNetworkAddress: "169.254.144.248/29",
}

var expectedConnectionsSlice = []con.Connection{
	connection1,
}

var getResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "operationStatus": "Completed",
        "redundant": true,
        "name": "YourConnectionName",
        "bandwidth": "100M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "privateRoute"
            }
        },
        "destination": {
            "primary": {
                "interconnect": "@Tokyo-CC2-1",
                "pairingKey": "d27476e6-e8a8-4214-a88f-9d3131db465d/asia-northeast1/2"
            },
            "qosType": "guarantee",
            "secondary": {
                "interconnect": "@Tokyo-CC2-1",
                "pairingKey": "17c64c4e-f845-4450-82e9-843095e18526/asia-northeast1/2"
            }
        },
        "primaryConnectedNwAddress": "169.254.51.168/29",
        "secondaryConnectedNwAddress": "169.254.144.248/29"
    }
}
`,
	idConnection1,
)

const createRequest = `
{
    "connection": {
        "name": "YourConnectionName",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "privateRoute"
            }
        },
        "destination": {
            "primary": {
                "interconnect": "@Tokyo-CC2-1",
                "pairingKey": "d27476e6-e8a8-4214-a88f-9d3131db465d/asia-northeast1/2"
            },
            "secondary": {
                "interconnect": "@Tokyo-CC2-1",
                "pairingKey": "17c64c4e-f845-4450-82e9-843095e18526/asia-northeast1/2"
            },
            "qosType": "guarantee"
        },
        "bandwidth": "100M"
    }
}`

var createResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "operationStatus": "Processing",
        "redundant": true,
        "name": "YourConnectionName",
        "bandwidth": "100M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "privateRoute"
            }
        },
        "destination": {
            "primary": {
                "interconnect": "@Tokyo-CC2-1",
                "pairingKey": "d27476e6-e8a8-4214-a88f-9d3131db465d/asia-northeast1/2"
            },
            "qosType": "guarantee",
            "secondary": {
                "interconnect": "@Tokyo-CC2-1",
                "pairingKey": "17c64c4e-f845-4450-82e9-843095e18526/asia-northeast1/2"
            }
        },
        "operationId": "c8b8fd37e30b42ddb4d8b1b59adb9120"
    }
}
`,
	idConnection1,
)

var connectionCreated = con.Connection{
	ID:              idConnection1,
	TenantID:        "87e89b8f075a4ee1aa209f6ca6ce242c",
	OperationStatus: "Processing",
	Redundant:       true,
	Name:            "YourConnectionName",
	Bandwidth:       "100M",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		RouteFilter: con.RouteFilter{
			In:  "noRoute",
			Out: "privateRoute",
		},
	},
	Destination: con.Destination{
		Primary: con.DestinationHAInfo{
			Interconnect: "@Tokyo-CC2-1",
			PairingKey:   "d27476e6-e8a8-4214-a88f-9d3131db465d/asia-northeast1/2",
		},
		QosType: "guarantee",
		Secondary: con.DestinationHAInfo{
			Interconnect: "@Tokyo-CC2-1",
			PairingKey:   "17c64c4e-f845-4450-82e9-843095e18526/asia-northeast1/2",
		},
	},
	OperationID: "c8b8fd37e30b42ddb4d8b1b59adb9120",
}

const updateRequest = `
{
    "connection": {
        "source": {
            "routeFilter": {
                "in": "fullRoute",
                "out": "fullRoute"
            }
        }
    }
}`

var updateResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "operationStatus": "Processing",
        "redundant": true,
        "name": "YourConnectionName",
        "bandwidth": "100M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "fullRoute",
                "out": "fullRoute"
            }
        },
        "destination": {
            "primary": {
                "interconnect": "@Tokyo-CC2-1",
                "pairingKey": "d27476e6-e8a8-4214-a88f-9d3131db465d/asia-northeast1/2"
            },
            "qosType": "guarantee",
            "secondary": {
                "interconnect": "@Tokyo-CC2-1",
                "pairingKey": "17c64c4e-f845-4450-82e9-843095e18526/asia-northeast1/2"
            }
        },
        "primaryConnectedNwAddress": "169.254.51.168/29",
        "secondaryConnectedNwAddress": "169.254.144.248/29",
        "operationId": "70a44564172c48c3babe6f22365885aa"
    }
}
`,
	idConnection1,
)

var connectionUpdated = con.Connection{
	ID:              idConnection1,
	TenantID:        "87e89b8f075a4ee1aa209f6ca6ce242c",
	OperationStatus: "Processing",
	Redundant:       true,
	Name:            "YourConnectionName",
	Bandwidth:       "100M",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		RouteFilter: con.RouteFilter{
			In:  "fullRoute",
			Out: "fullRoute",
		},
	},
	Destination: con.Destination{
		Primary: con.DestinationHAInfo{
			Interconnect: "@Tokyo-CC2-1",
			PairingKey:   "d27476e6-e8a8-4214-a88f-9d3131db465d/asia-northeast1/2",
		},
		QosType: "guarantee",
		Secondary: con.DestinationHAInfo{
			Interconnect: "@Tokyo-CC2-1",
			PairingKey:   "17c64c4e-f845-4450-82e9-843095e18526/asia-northeast1/2",
		},
	},
	PrimaryConnectedNetworkAddress:   "169.254.51.168/29",
	SecondaryConnectedNetworkAddress: "169.254.144.248/29",
	OperationID:                      "70a44564172c48c3babe6f22365885aa",
}
