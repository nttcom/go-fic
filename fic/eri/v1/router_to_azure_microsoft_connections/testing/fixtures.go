package testing

import (
	"fmt"

	con "github.com/nttcom/go-fic/fic/eri/v1/router_to_azure_microsoft_connections"
)

const idConnection1 = "F010123456789"

var listResponse = fmt.Sprintf(`
{
    "connections": [
        {
            "id": "%s",
            "name": "YourConnectionName",
            "redundant": true,
            "tenantId": "06a90740d6764465896091b1f0676048",
            "area": "JPEAST",
            "operationStatus": "Processing",
            "bandwidth": "100M",
            "source": {
                "routerId": "F020123456789",
                "groupName": "group_1",
                "routeFilter": {
                    "in": "fullRoute",
                    "out": "natRoute"
                }
            },
            "destination": {
                "interconnect": "Tokyo-1",
                "qosType": "guarantee",
                "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a",
                "advertisedPublicPrefixes": [
                    "100.131.65.2/32"
                ]
            },
            "primaryConnectedNwAddress": "100.131.67.0/30",
            "secondaryConnectedNwAddress": "100.131.67.4/30"
        }
    ]
}
`,
	idConnection1,
)

var connection1 = con.Connection{
	ID:              idConnection1,
	TenantID:        "06a90740d6764465896091b1f0676048",
	OperationStatus: "Processing",
	Redundant:       true,
	Name:            "YourConnectionName",
	Bandwidth:       "100M",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		RouteFilter: con.RouteFilter{
			In:  "fullRoute",
			Out: "natRoute",
		},
	},
	Destination: con.Destination{
		Interconnect: "Tokyo-1",
		QosType:      "guarantee",
		ServiceKey:   "6191af11-82f9-4c15-9894-9a69c8f8628a",
		AdvertisedPublicPrefixes: []string{
			"100.131.65.2/32",
		},
	},
	PrimaryConnectedNetworkAddress:   "100.131.67.0/30",
	SecondaryConnectedNetworkAddress: "100.131.67.4/30",
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
        "tenantId": "06a90740d6764465896091b1f0676048",
        "area": "JPEAST",
        "operationStatus": "Processing",
        "bandwidth": "100M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "fullRoute",
                "out": "natRoute"
            }
        },
        "destination": {
            "interconnect": "Tokyo-1",
            "qosType": "guarantee",
            "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a",
            "advertisedPublicPrefixes": [
                "100.131.65.2/32"
            ]
        },
        "primaryConnectedNwAddress": "100.131.67.0/30",
        "secondaryConnectedNwAddress": "100.131.67.4/30"
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
                "out": "natRoute"
            }
        },
        "destination": {
            "interconnect": "Tokyo-1",
            "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a",
            "qosType": "guarantee",
            "advertisedPublicPrefixes": [
                "100.131.65.2/32"
            ]
        },
        "bandwidth": "100M"
    }
}`

var createResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "name": "YourConnectionName",
        "redundant": true,
        "tenantId": "06a90740d6764465896091b1f0676048",
        "area": "JPEAST",
        "operationStatus": "Processing",
        "bandwidth": "100M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "fullRoute",
                "out": "natRoute"
            }
        },
        "destination": {
            "interconnect": "Tokyo-1",
            "qosType": "guarantee",
            "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a",
            "advertisedPublicPrefixes": [
                "100.131.65.2/32"
            ]
        },
        "primaryConnectedNwAddress": null,
        "secondaryConnectedNwAddress": null,
        "operationId": "adff6493165a40849d504eee0dd3fce4"
    }
}`,
	idConnection1,
)

var connectionCreated = con.Connection{
	ID:              idConnection1,
	TenantID:        "06a90740d6764465896091b1f0676048",
	OperationStatus: "Processing",
	Redundant:       true,
	Name:            "YourConnectionName",
	Bandwidth:       "100M",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		RouteFilter: con.RouteFilter{
			In:  "fullRoute",
			Out: "natRoute",
		},
	},
	Destination: con.Destination{
		Interconnect: "Tokyo-1",
		QosType:      "guarantee",
		ServiceKey:   "6191af11-82f9-4c15-9894-9a69c8f8628a",
		AdvertisedPublicPrefixes: []string{
			"100.131.65.2/32",
		},
	},
	OperationID: "adff6493165a40849d504eee0dd3fce4",
	Area:        "JPEAST",
}

const updateRequest = `
{
    "connection": {
        "source": {
            "routeFilter": {
                "in": "noRoute",
                "out": "noRoute"
            }
        },
        "destination": {
            "advertisedPublicPrefixes": [
                "100.131.65.2/32",
                "100.131.65.3/32"
            ]
        }
    }
}`

var updateResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "name": "YourConnectionName",
        "redundant": true,
        "tenantId": "06a90740d6764465896091b1f0676048",
        "area": "JPEAST",
        "operationStatus": "Processing",
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
            "interconnect": "Tokyo-1",
            "qosType": "guarantee",
            "serviceKey": "6191af11-82f9-4c15-9894-9a69c8f8628a",
            "advertisedPublicPrefixes": [
                "100.131.65.2/32",
                "100.131.65.3/32"
            ]
        },
        "primaryConnectedNwAddress": "100.131.67.12/30",
        "secondaryConnectedNwAddress": "100.131.67.16/30",
        "operationId": "98098f4df3934787b54a3aa4a655c3b8"
    }
}`,
	idConnection1,
)

var connectionUpdated = con.Connection{
	ID:              idConnection1,
	TenantID:        "06a90740d6764465896091b1f0676048",
	OperationStatus: "Processing",
	Redundant:       true,
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
		Interconnect: "Tokyo-1",
		QosType:      "guarantee",
		ServiceKey:   "6191af11-82f9-4c15-9894-9a69c8f8628a",
		AdvertisedPublicPrefixes: []string{
			"100.131.65.2/32",
			"100.131.65.3/32",
		},
	},
	PrimaryConnectedNetworkAddress:   "100.131.67.12/30",
	SecondaryConnectedNetworkAddress: "100.131.67.16/30",
	OperationID:                      "98098f4df3934787b54a3aa4a655c3b8",
	Area:                             "JPEAST",
}
