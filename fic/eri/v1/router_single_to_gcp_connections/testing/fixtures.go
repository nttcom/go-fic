package testing

import (
	"fmt"

	con "github.com/nttcom/go-fic/fic/eri/v1/router_single_to_gcp_connections"
)

const idConnection1 = "F010123456789"

var one = interface{}(1)
var null = interface{}(nil)

var listResponse = fmt.Sprintf(`
{
    "connections": [
        {
            "id": "%s",
            "name": "YourConnectionName",
            "redundant": false,
            "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
            "area": "JPEAST",
            "operationStatus": "Completed",
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
                        "in": null,
                        "out": null
                    }
                }
            },
            "destination": {
                "primary": {
                    "interconnect": "Equinix-TY2-2",
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
	Area:            "JPEAST",
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
				In:  &null,
				Out: &null,
			},
		},
	},
	Destination: con.Destination{
		Primary: con.DestinationHAInfo{
			Interconnect: "Equinix-TY2-2",
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
        "name": "YourConnectionName",
        "redundant": false,
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "area": "JPEAST",
        "operationStatus": "Completed",
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
                    "in": null,
                    "out": null
                }
            }
        },
        "destination": {
            "primary": {
                "interconnect": "Equinix-TY2-2",
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
                    "in": null,
                    "out": null
                }
            }
        },
        "destination": {
            "primary": {
                "interconnect": "Equinix-TY2-2",
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
        "name": "YourConnectionName",
        "redundant": false,
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "area": "JPEAST",
        "operationStatus": "Processing",
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
                    "in": null,
                    "out": null
                }
            }
        },
        "destination": {
            "primary": {
                "interconnect": "Equinix-TY2-2",
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
	Area:            "JPEAST",
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
				In:  &null,
				Out: &null,
			},
		},
	},
	Destination: con.Destination{
		Primary: con.DestinationHAInfo{
			Interconnect: "Equinix-TY2-2",
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
                    "out": 1
                }
            }
        },
        "bandwidth": "200M"
    }
}
`

var updateResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "name": "YourConnectionName",
        "redundant": false,
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "area": "JPEAST",
        "operationStatus": "Processing",
        "bandwidth": "200M",
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
                    "out": 1
                }
            }
        },
        "destination": {
            "primary": {
                "interconnect": "Equinix-TY2-2",
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
	Area:            "JPEAST",
	OperationStatus: "Processing",
	Redundant:       false,
	Name:            "YourConnectionName",
	Bandwidth:       "200M",
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
				Out: &one,
			},
		},
	},
	Destination: con.Destination{
		Primary: con.DestinationHAInfo{
			Interconnect: "Equinix-TY2-2",
			PairingKey:   "4f043662-2ed6-45b0-8c3a-77d685716e28/asia-northeast1/2",
		},
		QosType: "guarantee",
	},
	PrimaryConnectedNetworkAddress: "169.254.51.168/29",
	OperationID:                    "6b876b42498e4b8ea2c749d18dbe10de",
}
