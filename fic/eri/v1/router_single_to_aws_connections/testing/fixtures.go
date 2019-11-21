package testing

import (
	"fmt"

	con "github.com/nttcom/go-fic/fic/eri/v1/router_single_to_aws_connections"
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
            "bandwidth": "100M",
            "source": {
                "routerId": "F010123456789",
                "groupName": "group_1",
                "routeFilter": {
                    "in": "noRoute",
                    "out": "defaultRoute"
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
                    "asn": "65530"
                },
                "qosType": "guarantee",
                "awsAccountId": "123456789012"
            },
            "primaryConnectedNwAddress": "10.0.0.0/30"
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
		RouterID:  "F010123456789",
		GroupName: "group_1",
		RouteFilter: con.RouteFilter{
			In:  "noRoute",
			Out: "defaultRoute",
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
			ASN:          "65530",
		},
		QosType:      "guarantee",
		AWSAccountID: "123456789012",
	},
	PrimaryConnectedNetworkAddress: "10.0.0.0/30",
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
            "routerId": "F010123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "defaultRoute"
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
                "asn": "65530"
            },
            "qosType": "guarantee",
            "awsAccountId": "123456789012"
        },
        "primaryConnectedNwAddress": "10.0.0.0/30"
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
            "routerId": "F010123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "defaultRoute"
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
                "asn": "65530"
            },
            "awsAccountId": "123456789012",
            "qosType": "guarantee"
        },
        "bandwidth": "100M",
        "primaryConnectedNwAddress": "10.0.0.0/30"
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
            "routerId": "F010123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "defaultRoute"
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
                "asn": "65530"
            },
            "qosType": "guarantee",
            "awsAccountId": "123456789012"
        },
        "primaryConnectedNwAddress": "10.0.0.0/30",
        "operationId": "cb1a2e21fae64397a6a59f8474e9a2ec"
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
		RouterID:  "F010123456789",
		GroupName: "group_1",
		RouteFilter: con.RouteFilter{
			In:  "noRoute",
			Out: "defaultRoute",
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
			ASN:          "65530",
		},
		QosType:      "guarantee",
		AWSAccountID: "123456789012",
	},
	PrimaryConnectedNetworkAddress: "10.0.0.0/30",
	OperationID:                    "cb1a2e21fae64397a6a59f8474e9a2ec",
}

const updateRequest = `
{
    "connection": {
        "source": {
            "routeFilter": {
                "in": "noRoute",
                "out": "privateRoute"
            },
            "primary": {
                "asPathPrepend": {
                    "in": 2,
                    "out": 1
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
        "bandwidth": "100M",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "privateRoute"
            },
            "primary": {
                "asPathPrepend": {
                    "in": 2,
                    "out": 1
                }
            }
        },
        "destination": {
            "primary": {
                "interconnect": "Equinix-TY2-1",
                "asn": "65530"
            },
            "qosType": "guarantee",
            "awsAccountId": "123456789012"
        },
        "primaryConnectedNwAddress": "10.0.0.0/30",
        "operationId": "172165e7e3cc483d9cc2248c78614337"
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
			In:  "noRoute",
			Out: "privateRoute",
		},
		Primary: con.Primary{
			ASPathPrepend: con.ASPathPrepend{
				In:  &two,
				Out: &one,
			},
		},
	},
	Destination: con.Destination{
		Primary: con.DestinationHAInfo{
			Interconnect: "Equinix-TY2-1",
			ASN:          "65530",
		},
		QosType:      "guarantee",
		AWSAccountID: "123456789012",
	},
	PrimaryConnectedNetworkAddress: "10.0.0.0/30",
	OperationID:                    "172165e7e3cc483d9cc2248c78614337",
}
