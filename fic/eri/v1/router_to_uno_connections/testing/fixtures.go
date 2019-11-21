package testing

import (
	"fmt"

	con "github.com/nttcom/go-fic/fic/eri/v1/router_to_uno_connections"
)

const idConnection1 = "F010123456789"

var null = interface{}(nil)

var listResponse = fmt.Sprintf(`
{
    "connections": [
        {
            "id": "%s",
            "tenantId": "06a90740d6764465896091b1f0676048",
            "operationStatus": "Completed",
            "redundant": true,
            "name": "YourConnectionName",
            "bandwidth": "10M",
            "connectedNwAddress": "10.10.0.0/29",
            "source": {
                "routerId": "F020123456789",
                "groupName": "group_1",
                "routeFilter": {
                    "in": "noRoute",
                    "out": "fullRouteWithDefaultRoute"
                }
            },
            "destination": {
                "interconnect": "Interconnect-Tokyo-1",
                "vpnNumber": "V12345678",
                "parentContractNumber": "N123456789",
                "contractNumber": "N987654321",
                "qosType": "guarantee",
                "routeFilter": {
                    "out": "fullRoute"
                }
            }
        }
    ]
}
`,
	idConnection1,
)

var connection1 = con.Connection{
	ID:                      idConnection1,
	TenantID:                "06a90740d6764465896091b1f0676048",
	OperationStatus:         "Completed",
	Redundant:               true,
	Name:                    "YourConnectionName",
	ConnectedNetworkAddress: "10.10.0.0/29",
	Bandwidth:               "10M",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		RouteFilter: con.SourceRouteFilter{
			In:  "noRoute",
			Out: "fullRouteWithDefaultRoute",
		},
	},
	Destination: con.Destination{
		Interconnect:         "Interconnect-Tokyo-1",
		VPNNumber:            "V12345678",
		ParentContractNumber: "N123456789",
		ContractNumber:       "N987654321",
		QosType:              "guarantee",
		RouteFilter: con.DestinationRouteFilter{
			Out: "fullRoute",
		},
	},
}

var expectedConnectionsSlice = []con.Connection{
	connection1,
}

var getResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "tenantId": "06a90740d6764465896091b1f0676048",
        "operationStatus": "Completed",
        "redundant": true,
        "name": "YourConnectionName",
        "bandwidth": "10M",
        "connectedNwAddress": "10.10.0.0/29",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "fullRouteWithDefaultRoute"
            }
        },
        "destination": {
            "interconnect": "Interconnect-Tokyo-1",
            "vpnNumber": "V12345678",
            "parentContractNumber": "N123456789",
            "contractNumber": "N987654321",
            "qosType": "guarantee",
            "routeFilter": {
                "out": "fullRoute"
            }
        }
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
            "interconnect": "Interconnect-Tokyo-1",
            "parentContractNumber": "N123456789",
            "vpnNumber": "V12345678",
            "qosType": "guarantee",
            "routeFilter": {
                "out": "fullRoute"
            }
        },
        "bandwidth": "10M",
        "connectedNwAddress": "10.10.0.0/29"
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
        "bandwidth": "10M",
        "connectedNwAddress": "10.10.0.0/29",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "fullRouteWithDefaultRoute"
            }
        },
        "destination": {
            "interconnect": "Interconnect-Tokyo-1",
            "vpnNumber": "V12345678",
            "parentContractNumber": "N123456789",
            "contractNumber": null,
            "qosType": "guarantee",
            "routeFilter": {
                "out": "fullRoute"
            },
            "cNumber": "C1234567890"
        },
        "operationId": "14d2ea73cf2c477a83ab7ac031f1aee3"
    }
}
`,
	idConnection1,
)

var connectionCreated = con.Connection{
	ID:                      idConnection1,
	TenantID:                "06a90740d6764465896091b1f0676048",
	OperationStatus:         "Processing",
	Redundant:               true,
	Name:                    "YourConnectionName",
	Bandwidth:               "10M",
	ConnectedNetworkAddress: "10.10.0.0/29",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		RouteFilter: con.SourceRouteFilter{
			In:  "noRoute",
			Out: "fullRouteWithDefaultRoute",
		},
	},
	Destination: con.Destination{
		Interconnect:         "Interconnect-Tokyo-1",
		VPNNumber:            "V12345678",
		ParentContractNumber: "N123456789",
		ContractNumber:       null,
		QosType:              "guarantee",
		RouteFilter: con.DestinationRouteFilter{
			Out: "fullRoute",
		},
		CNumber: "C1234567890",
	},
	OperationID: "14d2ea73cf2c477a83ab7ac031f1aee3",
}

const updateRequest = `
{
    "connection": {
        "source": {
            "routeFilter": {
                "in": "fullRoute",
                "out": "fullRoute"
            }
        },
        "destination": {
            "routeFilter": {
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
        "bandwidth": "10M",
        "connectedNwAddress": "10.10.0.0/29",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "fullRoute",
                "out": "fullRoute"
            }
        },
        "destination": {
            "interconnect": "Interconnect-Tokyo-1",
            "vpnNumber": "V12345678",
            "parentContractNumber": "N123456789",
            "contractNumber": "N987654321",
            "qosType": "guarantee",
            "routeFilter": {
                "out": "defaultRoute"
            }
        },
        "operationId": "a95666b85c8d4abcb0419a5ab27b7af0"
    }
}
`,
	idConnection1,
)

var connectionUpdated = con.Connection{
	ID:                      idConnection1,
	TenantID:                "06a90740d6764465896091b1f0676048",
	OperationStatus:         "Processing",
	Redundant:               true,
	Name:                    "YourConnectionName",
	Bandwidth:               "10M",
	ConnectedNetworkAddress: "10.10.0.0/29",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		RouteFilter: con.SourceRouteFilter{
			In:  "fullRoute",
			Out: "fullRoute",
		},
	},
	Destination: con.Destination{
		Interconnect:         "Interconnect-Tokyo-1",
		VPNNumber:            "V12345678",
		ParentContractNumber: "N123456789",
		ContractNumber:       "N987654321",
		QosType:              "guarantee",
		RouteFilter: con.DestinationRouteFilter{
			Out: "defaultRoute",
		},
	},
	OperationID: "a95666b85c8d4abcb0419a5ab27b7af0",
}
