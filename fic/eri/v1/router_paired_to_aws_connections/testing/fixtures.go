package testing

import (
	"fmt"

	con "github.com/nttcom/go-fic/fic/eri/v1/router_paired_to_aws_connections"
)

const idConnection1 = "F010123456789"

var listResponse = fmt.Sprintf(`
{
    "connections": [
        {
            "id": "%s",
            "tenantId": "06a90740d6764465896091b1f0676048",
            "operationStatus": "Processing",
            "redundant": true,
            "name": "YourConnectionName",
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
                "primary": {
                    "interconnect": "Equinix-TY2-1",
                    "asn": "65000"
                },
                "qosType": "guarantee",
                "awsAccountId": "123456789012",
                "secondary": {
                    "interconnect": "@Tokyo-CC2-1",
                    "asn": "65001"
                }
            },
            "primaryConnectedNwAddress": "10.0.1.0/30",
            "secondaryConnectedNwAddress": "10.0.1.4/30"
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
			In:  "noRoute",
			Out: "fullRouteWithDefaultRoute",
		},
	},
	Destination: con.Destination{
		Primary: con.DestinationHAInfo{
			Interconnect: "Equinix-TY2-1",
			ASN:          "65000",
		},
		QosType:      "guarantee",
		AWSAccountID: "123456789012",
		Secondary: con.DestinationHAInfo{
			Interconnect: "@Tokyo-CC2-1",
			ASN:          "65001",
		},
	},
	PrimaryConnectedNetworkAddress:   "10.0.1.0/30",
	SecondaryConnectedNetworkAddress: "10.0.1.4/30",
}

var expectedConnectionsSlice = []con.Connection{
	connection1,
}

var getResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "tenantId": "06a90740d6764465896091b1f0676048",
        "operationStatus": "Processing",
        "redundant": true,
        "name": "YourConnectionName",
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
            "primary": {
                "interconnect": "Equinix-TY2-1",
                "asn": "65000"
            },
            "qosType": "guarantee",
            "awsAccountId": "123456789012",
            "secondary": {
                "interconnect": "@Tokyo-CC2-1",
                "asn": "65001"
            }
        },
        "primaryConnectedNwAddress": "10.0.1.0/30",
        "secondaryConnectedNwAddress": "10.0.1.4/30"
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
            }
        },
        "destination": {
            "primary": {
                "interconnect": "Equinix-TY2-1",
                "asn": "65000"
            },
            "secondary": {
                "interconnect": "@Tokyo-CC2-1",
                "asn": "65001"
            },
            "awsAccountId": "123456789012",
            "qosType": "guarantee"
        },
        "bandwidth": "100M",
        "primaryConnectedNwAddress": "10.0.1.0/30",
        "secondaryConnectedNwAddress": "10.0.1.4/30"
    }
}`

var createResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "tenantId": "06a90740d6764465896091b1f0676048",
        "operationStatus": "Processing",
        "redundant": true,
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
            "primary": {
                "interconnect": "Equinix-TY2-1",
                "asn": "65000"
            },
            "qosType": "guarantee",
            "awsAccountId": "123456789012",
            "secondary": {
                "interconnect": "@Tokyo-CC2-1",
                "asn": "65001"
            }
        },
        "primaryConnectedNwAddress": "10.0.1.0/30",
        "secondaryConnectedNwAddress": "10.0.1.4/30",
        "operationId": "e809e28da5e449208e58cd3accdc0454"
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
			Out: "fullRouteWithDefaultRoute",
		},
	},
	Destination: con.Destination{
		Primary: con.DestinationHAInfo{
			Interconnect: "Equinix-TY2-1",
			ASN:          "65000",
		},
		QosType:      "guarantee",
		AWSAccountID: "123456789012",
		Secondary: con.DestinationHAInfo{
			Interconnect: "@Tokyo-CC2-1",
			ASN:          "65001",
		},
	},
	PrimaryConnectedNetworkAddress:   "10.0.1.0/30",
	SecondaryConnectedNetworkAddress: "10.0.1.4/30",
	OperationID:                      "e809e28da5e449208e58cd3accdc0454",
}

const updateRequest = `
{
    "connection": {
        "source": {
            "routeFilter": {
                "in": "noRoute",
                "out": "defaultRoute"
            }
        }
    }
}`

var updateResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "tenantId": "06a90740d6764465896091b1f0676048",
        "operationStatus": "Processing",
        "redundant": true,
        "name": "YourConnectionName",
        "bandwidth": "100M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "defaultRoute"
            }
        },
        "destination": {
            "primary": {
                "interconnect": "Equinix-TY2-1",
                "asn": "65000"
            },
            "qosType": "guarantee",
            "awsAccountId": "123456789012",
            "secondary": {
                "interconnect": "@Tokyo-CC2-1",
                "asn": "65001"
            }
        },
        "primaryConnectedNwAddress": "10.0.1.0/30",
        "secondaryConnectedNwAddress": "10.0.1.4/30",
        "operationId": "e809e28da5e449208e58cd3accdc0454"
    }
}
`,
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
			Out: "defaultRoute",
		},
	},
	Destination: con.Destination{
		Primary: con.DestinationHAInfo{
			Interconnect: "Equinix-TY2-1",
			ASN:          "65000",
		},
		Secondary: con.DestinationHAInfo{
			Interconnect: "@Tokyo-CC2-1",
			ASN:          "65001",
		},
		QosType:      "guarantee",
		AWSAccountID: "123456789012",
	},
	PrimaryConnectedNetworkAddress:   "10.0.1.0/30",
	SecondaryConnectedNetworkAddress: "10.0.1.4/30",
	OperationID:                      "e809e28da5e449208e58cd3accdc0454",
}
