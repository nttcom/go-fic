package testing

import (
	"fmt"

	con "github.com/nttcom/go-fic/fic/eri/v1/router_single_to_gcp_connections"
)

const idConnection1 = "F010123456789"

var one = interface{}(1)
var two = interface{}(2)
var null = interface{}(nil)

var listResponse = fmt.Sprintf(`
{
    "connections": [
        {
            "id": "%s",
            "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
            "operationStatus": "Completed",
            "redundant": false,
            "name": "YourConnectionName",
            "bandwidth": "300M",
            "source": {
                "routerId": "F020123456789",
                "groupName": "group_1",
                "routeFilter": {
                    "in": "fullRoute",
                    "out": "fullRouteWithDefaultRoute"
                },
                "primary": {
                    "asPathPrepend": {
                        "in": 1,
                        "out": null
                    }
                }
            },
            "destination": {
                "primary": {
                    "interconnect": "Equinix-TY2-1",
                    "pairingKey": "4f043662-2ed6-45b0-8c3a-77d685716e28/asia-northeast1/2"
                },
                "qosType": "guarantee"
            },
            "primaryConnectedNwAddress": "169.254.51.168/29"
        }
    ]
}
`,
	idConnection1,
)

var connection1 = con.Connection{
	ID:              idConnection1,
	TenantID:        "87e89b8f075a4ee1aa209f6ca6ce242c",
	OperationStatus: "Completed",
	Redundant:       false,
	Name:            "YourConnectionName",
	Bandwidth:       "300M",
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
	PrimaryConnectedNetworkAddress: "169.254.51.168/29",
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
        "redundant": false,
        "name": "YourConnectionName",
        "bandwidth": "300M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "fullRoute",
                "out": "fullRouteWithDefaultRoute"
            },
            "primary": {
                "asPathPrepend": {
                    "in": 1,
                    "out": null
                }
            }
        },
        "destination": {
            "primary": {
                "interconnect": "Equinix-TY2-1",
                "pairingKey": "4f043662-2ed6-45b0-8c3a-77d685716e28/asia-northeast1/2"
            },
            "qosType": "guarantee"
        },
        "primaryConnectedNwAddress": "169.254.51.168/29"
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
                "in": "fullRoute",
                "out": "fullRouteWithDefaultRoute"
            },
            "primary": {
                "asPathPrepend": {
                    "in": 1,
                    "out": null
                }
            }
        },
        "destination": {
            "primary": {
                "interconnect": "Equinix-TY2-1",
                "pairingKey": "4f043662-2ed6-45b0-8c3a-77d685716e28/asia-northeast1/2"
            },
            "qosType": "guarantee"
        },
        "bandwidth": "300M"
    }
}
`

var createResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "operationStatus": "Processing",
        "redundant": false,
        "name": "YourConnectionName",
        "bandwidth": "300M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "fullRoute",
                "out": "fullRouteWithDefaultRoute"
            },
            "primary": {
                "asPathPrepend": {
                    "in": 1,
                    "out": null
                }
            }
        },
        "destination": {
            "primary": {
                "interconnect": "Equinix-TY2-1",
                "pairingKey": "4f043662-2ed6-45b0-8c3a-77d685716e28/asia-northeast1/2"
            },
            "qosType": "guarantee"
        },
        "operationId": "36867f21099f4497b3201753ba58326a"
    }
}
`,
	idConnection1,
)

var connectionCreated = con.Connection{
	ID:              idConnection1,
	TenantID:        "87e89b8f075a4ee1aa209f6ca6ce242c",
	OperationStatus: "Processing",
	Redundant:       false,
	Name:            "YourConnectionName",
	Bandwidth:       "300M",
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
	OperationID: "36867f21099f4497b3201753ba58326a",
}

const updateRequest = `
{
    "connection": {
        "source": {
            "routeFilter": {
                "in": "noRoute",
                "out": "defaultRoute"
            },
            "primary": {
                "asPathPrepend": {
                    "in": 1,
                    "out": 2
                }
            }
        }
    }
}
`

var updateResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "operationStatus": "Processing",
        "redundant": false,
        "name": "YourConnectionName",
        "bandwidth": "300M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "defaultRoute"
            },
            "primary": {
                "asPathPrepend": {
                    "in": 1,
                    "out": 2
                }
            }
        },
        "destination": {
            "primary": {
                "interconnect": "Equinix-TY2-1",
                "pairingKey": "4f043662-2ed6-45b0-8c3a-77d685716e28/asia-northeast1/2"
            },
            "qosType": "guarantee"
        },
        "primaryConnectedNwAddress": "169.254.51.168/29",
        "operationId": "6b876b42498e4b8ea2c749d18dbe10de"
    }
}
`,
	idConnection1,
)

var connectionUpdated = con.Connection{
	ID:              idConnection1,
	TenantID:        "87e89b8f075a4ee1aa209f6ca6ce242c",
	OperationStatus: "Processing",
	Redundant:       false,
	Name:            "YourConnectionName",
	Bandwidth:       "300M",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
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
	Destination: con.Destination{
		Primary: con.DestinationHAInfo{
			Interconnect: "Equinix-TY2-1",
			PairingKey:   "4f043662-2ed6-45b0-8c3a-77d685716e28/asia-northeast1/2",
		},
		QosType: "guarantee",
	},
	PrimaryConnectedNetworkAddress: "169.254.51.168/29",
	OperationID:                    "6b876b42498e4b8ea2c749d18dbe10de",
}
