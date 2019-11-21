package testing

import (
	"fmt"

	con "github.com/nttcom/go-fic/fic/eri/v1/router_to_ecl_connections"
)

const idConnection1 = "F010123456789"

var listResponse = fmt.Sprintf(`
{
    "connections": [
        {
            "id": "%s",
            "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
            "operationStatus": "Completed",
            "redundant": false,
            "name": "YourConnectionName",
            "bandwidth": "100M",
            "source": {
                "routerId": "F020123456789",
                "groupName": "group_1",
                "routeFilter": {
                    "in": "noRoute",
                    "out": "noRoute"
                }
            },
            "destination": {
                "interconnect": "JP3-1",
                "qosType": "guarantee",
                "eclTenantId": "20c33449388f4071bf629b15fd9237bd"
            },
            "primaryConnectedNwAddress": "10.0.0.0/30",
            "secondaryConnectedNwAddress": "10.10.0.0/30"
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
	Bandwidth:       "100M",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		RouteFilter: con.RouteFilter{
			In:  "noRoute",
			Out: "noRoute",
		},
	},
	Destination: con.Destination{
		Interconnect: "JP3-1",
		QosType:      "guarantee",
		ECLTenantID:  "20c33449388f4071bf629b15fd9237bd",
	},
	PrimaryConnectedNetworkAddress:   "10.0.0.0/30",
	SecondaryConnectedNetworkAddress: "10.10.0.0/30",
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
        "bandwidth": "100M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "noRoute"
            }
        },
        "destination": {
            "interconnect": "JP3-1",
            "qosType": "guarantee",
            "eclTenantId": "20c33449388f4071bf629b15fd9237bd"
        },
        "primaryConnectedNwAddress": "10.0.0.0/30",
        "secondaryConnectedNwAddress": "10.10.0.0/30"
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
                "out": "noRoute"
            }
        },
        "destination": {
            "interconnect": "JP3-1",
            "qosType": "guarantee",
            "eclTenantId": "20c33449388f4071bf629b15fd9237bd",
            "eclApiKey": "oYRQOVmDz3L7vD4oDMX8zaqXa15R6MhL",
            "eclApiSecretKey": "kmoLb2n7rTzXCWvd"
        },
        "bandwidth": "100M",
        "primaryConnectedNwAddress": "10.0.0.0/30",
        "secondaryConnectedNwAddress": "10.10.0.0/30"
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
        "bandwidth": "100M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "noRoute"
            }
        },
        "destination": {
            "interconnect": "JP3-1",
            "qosType": "guarantee",
            "eclTenantId": "20c33449388f4071bf629b15fd9237bd",
            "eclApiKey": "oYRQOVmDz3L7vD4oDMX8zaqXa15R6MhL",
            "eclApiSecretKey": "kmoLb2n7rTzXCWvd"
        },
        "primaryConnectedNwAddress": "10.0.0.0/30",
        "secondaryConnectedNwAddress": "10.10.0.0/30",
        "operationId": "d981d661a4be48bca8b748a84b0325c4"
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
	Bandwidth:       "100M",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		RouteFilter: con.RouteFilter{
			In:  "noRoute",
			Out: "noRoute",
		},
	},
	Destination: con.Destination{
		Interconnect:    "JP3-1",
		QosType:         "guarantee",
		ECLTenantID:     "20c33449388f4071bf629b15fd9237bd",
		ECLAPIKey:       "oYRQOVmDz3L7vD4oDMX8zaqXa15R6MhL",
		ECLAPISecretKey: "kmoLb2n7rTzXCWvd",
	},
	PrimaryConnectedNetworkAddress:   "10.0.0.0/30",
	SecondaryConnectedNetworkAddress: "10.10.0.0/30",
	OperationID:                      "d981d661a4be48bca8b748a84b0325c4",
}

const updateRequest = `
{
    "connection": {
        "source": {
            "routeFilter": {
                "in": "fullRoute",
                "out": "fullRouteWithDefaultRoute"
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
        "redundant": false,
        "name": "YourConnectionName",
        "bandwidth": "100M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "fullRoute",
                "out": "fullRouteWithDefaultRoute"
            }
        },
        "destination": {
            "interconnect": "JP3-1",
            "qosType": "guarantee",
            "eclTenantId": "20c33449388f4071bf629b15fd9237bd"
        },
        "primaryConnectedNwAddress": "10.0.0.0/30",
        "secondaryConnectedNwAddress": "10.10.0.0/30",
        "operationId": "8666715f42e84d6c8007e591f6e135f9"
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
	Bandwidth:       "100M",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		RouteFilter: con.RouteFilter{
			In:  "fullRoute",
			Out: "fullRouteWithDefaultRoute",
		},
	},
	Destination: con.Destination{
		Interconnect: "JP3-1",
		QosType:      "guarantee",
		ECLTenantID:  "20c33449388f4071bf629b15fd9237bd",
	},
	PrimaryConnectedNetworkAddress:   "10.0.0.0/30",
	SecondaryConnectedNetworkAddress: "10.10.0.0/30",
	OperationID:                      "8666715f42e84d6c8007e591f6e135f9",
}
