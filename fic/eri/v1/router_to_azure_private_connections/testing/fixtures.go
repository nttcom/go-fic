package testing

import (
	"fmt"

	con "github.com/nttcom/go-fic/fic/eri/v1/router_to_azure_private_connections"
)

const idConnection1 = "F010123456789"

var listResponse = fmt.Sprintf(`
{
    "connections": [
        {
            "id": "%s",
            "name": "YourConnectionName",
            "redundant": true,
            "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
            "area": "JPEAST",
            "operationStatus": "Processing",
            "bandwidth": "100M",
            "source": {
                "routerId": "F020123456789",
                "groupName": "group_1",
                "routeFilter": {
                    "in": "noRoute",
                    "out": "fullRouteWithDefaultRoute"
                }
            },
            "destination": {
                "interconnect": "Tokyo-1",
                "qosType": "guarantee",
                "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a"
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
	OperationStatus: "Processing",
	Redundant:       true,
	Name:            "YourConnectionName",
	Bandwidth:       "100M",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		RouteFilter: con.RouteFilter{
			In:  "noRoute",
			Out: "fullRouteWithDefaultRoute",
		},
	},
	Destination: con.Destination{
		Interconnect: "Tokyo-1",
		QosType:      "guarantee",
		ServiceKey:   "6191af11-82f9-4c15-9894-9a69c8f8628a",
	},
	PrimaryConnectedNetworkAddress:   "10.0.0.0/30",
	SecondaryConnectedNetworkAddress: "10.10.0.0/30",
	Area:                             "JPEAST",
}

var expectedConnectionsSlice = []con.Connection{
	connection1,
}

var getResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "name": "YourConnectionName",
        "redundant": true,
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "area": "JPEAST",
        "operationStatus": "Processing",
        "bandwidth": "100M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "fullRouteWithDefaultRoute"
            }
        },
        "destination": {
            "interconnect": "Tokyo-1",
            "qosType": "guarantee",
            "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a"
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
                "out": "fullRouteWithDefaultRoute"
            }
        },
        "destination": {
            "interconnect": "Tokyo-1",
            "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a",
            "qosType": "guarantee"
        },
        "bandwidth": "100M",
        "primaryConnectedNwAddress": "10.0.0.0/30",
        "secondaryConnectedNwAddress": "10.10.0.0/30"
    }
}`

var createResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "name": "YourConnectionName",
        "redundant": true,
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "area": "JPEAST",
        "operationStatus": "Processing",
        "bandwidth": "100M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "fullRouteWithDefaultRoute"
            }
        },
        "destination": {
            "interconnect": "Tokyo-1",
            "qosType": "guarantee",
            "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a"
        },
        "primaryConnectedNwAddress": "10.0.0.0/30",
        "secondaryConnectedNwAddress": "10.10.0.0/30",
        "operationId": "52fc9064e62f4c08947eac106e6abf8b"
    }
}`,
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
			Out: "fullRouteWithDefaultRoute",
		},
	},
	Destination: con.Destination{
		Interconnect: "Tokyo-1",
		QosType:      "guarantee",
		ServiceKey:   "6191af11-82f9-4c15-9894-9a69c8f8628a",
	},
	PrimaryConnectedNetworkAddress:   "10.0.0.0/30",
	SecondaryConnectedNetworkAddress: "10.10.0.0/30",
	OperationID:                      "52fc9064e62f4c08947eac106e6abf8b",
	Area:                             "JPEAST",
}

const updateRequest = `
{
    "connection": {
        "source": {
            "routeFilter": {
                "in": "fullRoute",
                "out": "defaultRoute"
            }
        }
    }
}`

var updateResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "name": "YourConnectionName",
        "redundant": true,
        "tenantId": "87e89b8f075a4ee1aa209f6ca6ce242c",
        "area": "JPEAST",
        "operationStatus": "Processing",
        "bandwidth": "100M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "fullRoute",
                "out": "defaultRoute"
            }
        },
        "destination": {
            "interconnect": "Tokyo-1",
            "qosType": "guarantee",
            "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a"
        },
        "primaryConnectedNwAddress": "10.0.0.0/30",
        "secondaryConnectedNwAddress": "10.10.0.0/30",
        "operationId": "eecf15b532d14d0cbf06b10df28fd431"
    }
}`,
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
			Out: "defaultRoute",
		},
	},
	Destination: con.Destination{
		Interconnect: "Tokyo-1",
		QosType:      "guarantee",
		ServiceKey:   "6191af11-82f9-4c15-9894-9a69c8f8628a",
	},
	PrimaryConnectedNetworkAddress:   "10.0.0.0/30",
	SecondaryConnectedNetworkAddress: "10.10.0.0/30",
	OperationID:                      "eecf15b532d14d0cbf06b10df28fd431",
	Area:                             "JPEAST",
}
