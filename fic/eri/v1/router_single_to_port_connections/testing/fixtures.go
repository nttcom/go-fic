package testing

import (
	"fmt"

	con "github.com/nttcom/go-fic/fic/eri/v1/router_single_to_port_connections"
)

const idConnection1 = "F010123456789"

var one = interface{}(1)
var two = interface{}(2)
var four = interface{}(4)
var five = interface{}(5)
var null = interface{}(nil)

var listResponse = fmt.Sprintf(`
{
    "connections": [
        {
            "id": "%s",
            "name": "YourConnectionName",
            "redundant": true,
            "tenantId": "06a90740d6764465896091b1f0676048",
            "area": "JPEAST",
            "operationStatus": "Completed",
            "source": {
                "routerId": "F020123456789",
                "groupName": "group_1",
                "routeFilter": {
                    "in": "noRoute",
                    "out": "fullRouteWithDefaultRoute"
                },
                "primary": {
                    "ipAddress": "10.0.0.1/30",
                    "asPathPrepend": {
                        "in": 1,
                        "out": null
                    }
                },
                "secondary": {}
            },
            "destination": {
                "primary": {
                    "portId": "F010123456789",
                    "vlan": 101,
                    "ipAddress": "10.0.0.2/30",
                    "asn": "65000"
                },
                "secondary": {}
            },
            "bandwidth": "10M"
        }
    ]
}
`,
	idConnection1,
)

var connection1 = con.Connection{
	ID:              idConnection1,
	Name:            "YourConnectionName",
	Redundant:       true,
	TenantID:        "06a90740d6764465896091b1f0676048",
	Area:            "JPEAST",
	OperationStatus: "Completed",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		RouteFilter: con.RouteFilter{
			In:  "noRoute",
			Out: "fullRouteWithDefaultRoute",
		},
		Primary: con.SourceHAInfo{
			IPAddress: "10.0.0.1/30",
			ASPathPrepend: con.ASPathPrepend{
				In:  &one,
				Out: &null,
			},
		},
		// Secondary: con.SourceHAInfo{
		// 	IPAddress: "10.30.0.1/30",
		// 	ASPathPrepend: con.ASPathPrepend{
		// 		In:  &five,
		// 		Out: &four,
		// 	},
		// },
	},
	Destination: con.Destination{
		Primary: con.DestinationHAInfo{
			PortID:    "F010123456789",
			VLAN:      101,
			IPAddress: "10.0.0.2/30",
			ASN:       "65000",
		},
		// Secondary: con.DestinationHAInfo{
		// 	PortID:    "F019876543210",
		// 	VLAN:      102,
		// 	IPAddress: "10.0.0.6/30",
		// 	ASN:       "65000",
		// },
	},
	Bandwidth: "10M",
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
        "operationStatus": "Completed",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "fullRouteWithDefaultRoute"
            },
            "primary": {
                "ipAddress": "10.0.0.1/30",
                "asPathPrepend": {
                    "in": 1,
                    "out": null
                }
            },
            "secondary": {
                "ipAddress": "10.30.0.1/30",
                "asPathPrepend": {
                    "in": 5,
                    "out": 4
                }
            }
        },
        "destination": {
            "primary": {
                "portId": "F010123456789",
                "vlan": 101,
                "ipAddress": "10.0.0.2/30",
                "asn": "65000"
            },
            "secondary": {
                "portId": "F019876543210",
                "vlan": 102,
                "ipAddress": "10.0.0.6/30",
                "asn": "65000"
            }
        },
        "bandwidth": "10M"
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
            "primary": {
                "ipAddress": "10.0.0.1/30",
                "asPathPrepend": {
                    "in": 4,
                    "out": 4
                }
            },
            "routeFilter": {
                "in": "fullRoute",
                "out": "fullRouteWithDefaultRoute"
            }
        },
        "destination": {
            "primary": {
                "portId": "F010123456789",
                "vlan": 101,
                "ipAddress": "10.0.0.2/30",
                "asn": "65000"
            }
        },
        "bandwidth": "100M"
    }
}
`

var createResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "name": "YourConnectionName",
        "redundant": true,
        "tenantId": "06a90740d6764465896091b1f0676048",
        "area": "JPEAST",
        "operationStatus": "Processing",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "fullRoute",
                "out": "fullRouteWithDefaultRoute"
            },
            "primary": {
                "ipAddress": "10.0.0.1/30",
                "asPathPrepend": {
                    "in": 4,
                    "out": 4
                }
            },
            "secondary": {}
        },
        "destination": {
            "primary": {
                "portId": "F010123456789",
                "vlan": 101,
                "ipAddress": "10.0.0.2/30",
                "asn": "65000"
            },
            "secondary": {}
        },
        "bandwidth": "100M",
        "operationId": "8d49e2ab41a54598aec02c0f198ab0e3"
    }
}
`,
	idConnection1,
)

var connectionCreated = con.Connection{
	ID:              idConnection1,
	Name:            "YourConnectionName",
	Redundant:       true,
	TenantID:        "06a90740d6764465896091b1f0676048",
	Area:            "JPEAST",
	OperationStatus: "Processing",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		Primary: con.SourceHAInfo{
			IPAddress: "10.0.0.1/30",
			ASPathPrepend: con.ASPathPrepend{
				In:  &four,
				Out: &four,
			},
		},
		// Secondary: con.SourceHAInfo{
		// 	IPAddress: "10.0.0.5/30",
		// 	ASPathPrepend: con.ASPathPrepend{
		// 		In:  &two,
		// 		Out: &one,
		// 	},
		// },
		RouteFilter: con.RouteFilter{
			In:  "fullRoute",
			Out: "fullRouteWithDefaultRoute",
		},
	},
	Destination: con.Destination{
		Primary: con.DestinationHAInfo{
			PortID:    "F010123456789",
			VLAN:      101,
			IPAddress: "10.0.0.2/30",
			ASN:       "65000",
		},
		// Secondary: con.DestinationHAInfo{
		// 	PortID:    "F019876543210",
		// 	VLAN:      102,
		// 	IPAddress: "10.0.0.6/30",
		// 	ASN:       "65000",
		// },
	},
	Bandwidth:   "100M",
	OperationID: "8d49e2ab41a54598aec02c0f198ab0e3",
}

const updateRequest = `
{
    "connection": {
        "source": {
            "primary": {
                "asPathPrepend": {
                    "in": null,
                    "out": 2
                }
            },
            "routeFilter": {
                "in": "noRoute",
                "out": "fullRoute"
            }
        }
    }
}
`

var updateResponse = fmt.Sprintf(`
{
    "connection": {
        "id": "%s",
        "name": "YourConnectionName",
        "redundant": true,
        "tenantId": "06a90740d6764465896091b1f0676048",
        "area": "JPEAST",
        "operationStatus": "Processing",
        "source": {
            "routerId": "F020123456789",
            "groupName": "group_1",
            "routeFilter": {
                "in": "noRoute",
                "out": "fullRoute"
            },
            "primary": {
                "ipAddress": "10.0.0.1/30",
                "asPathPrepend": {
                    "in": null,
                    "out": 2
                }
            },
            "secondary": {}
        },
        "destination": {
            "primary": {
                "portId": "F010123456789",
                "vlan": 101,
                "ipAddress": "10.0.0.2/30",
                "asn": "65000"
            },
            "secondary": {}
        },
        "bandwidth": "100M",
        "operationId": "8d49e2ab41a54598aec02c0f198ab0e3"
    }
}
`,
	idConnection1,
)

var connectionUpdated = con.Connection{
	ID:              idConnection1,
	Name:            "YourConnectionName",
	Redundant:       true,
	TenantID:        "06a90740d6764465896091b1f0676048",
	Area:            "JPEAST",
	OperationStatus: "Processing",
	Source: con.Source{
		RouterID:  "F020123456789",
		GroupName: "group_1",
		Primary: con.SourceHAInfo{
			IPAddress: "10.0.0.1/30",
			ASPathPrepend: con.ASPathPrepend{
				In:  &null,
				Out: &two,
			},
		},
		// Secondary: con.SourceHAInfo{
		// 	IPAddress: "10.0.0.5/30",
		// 	ASPathPrepend: con.ASPathPrepend{
		// 		In:  &null,
		// 		Out: &two,
		// 	},
		// },
		RouteFilter: con.RouteFilter{
			In:  "noRoute",
			Out: "fullRoute",
		},
	},
	Destination: con.Destination{
		Primary: con.DestinationHAInfo{
			PortID:    "F010123456789",
			VLAN:      101,
			IPAddress: "10.0.0.2/30",
			ASN:       "65000",
		},
		// Secondary: con.DestinationHAInfo{
		// 	PortID:    "F019876543210",
		// 	VLAN:      102,
		// 	IPAddress: "10.0.0.6/30",
		// 	ASN:       "65000",
		// },
	},
	Bandwidth:   "100M",
	OperationID: "8d49e2ab41a54598aec02c0f198ab0e3",
}
